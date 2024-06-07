package writer

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/grafana/dataplane/sdata/numeric"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/auth/identity"
	"github.com/grafana/grafana/pkg/services/datasources"
	"github.com/grafana/grafana/pkg/services/org"
	"github.com/grafana/grafana/pkg/services/pluginsintegration/adapters"
	"github.com/m3db/prometheus_remote_client_golang/promremote"

	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

const (
	// Fixed error messages
	MimirDuplicateTimestampError = "err-mimir-sample-duplicate-timestamp"

	// Best effort error messages
	PrometheusDuplicateTimestampError = "duplicate sample for timestamp"
)

var DuplicateTimestampErrors = [...]string{
	MimirDuplicateTimestampError,
	PrometheusDuplicateTimestampError,
}

// Metric represents a Prometheus time series metric.
type Metric struct {
	T time.Time
	V float64
}

// Point is a logical representation of a single point in time for a Prometheus time series.
type Point struct {
	Name   string
	Labels map[string]string
	Metric Metric
}

func PointsFromFrames(name string, t time.Time, frames data.Frames, extraLabels map[string]string) ([]Point, error) {
	cr, err := numeric.CollectionReaderFromFrames(frames)
	if err != nil {
		return nil, err
	}

	col, err := cr.GetCollection(false)
	if err != nil {
		return nil, err
	}

	points := make([]Point, 0, len(col.Refs))
	for _, ref := range col.Refs {
		f := math.NaN()
		if fp, empty, _ := ref.NullableFloat64Value(); !empty && fp != nil {
			f = *fp
		}

		metric := Metric{
			T: t,
			V: f,
		}

		labels := ref.GetLabels().Copy()
		if labels == nil {
			labels = data.Labels{}
		}
		delete(labels, "__name__")
		for k, v := range extraLabels {
			labels[k] = v
		}

		points = append(points, Point{
			Name:   name,
			Labels: labels,
			Metric: metric,
		})
	}

	return points, nil
}

type PrometheusWriterConfig struct {
	DatasourceUID string
	WritePath     string
	TenantID      string
	Timeout       time.Duration
}

type datasourceCache interface {
	GetDatasourceByUID(ctx context.Context, uid string, user identity.Requester, skipCache bool) (*datasources.DataSource, error)
}

type httpClientProvider interface {
	GetTransport(opts ...httpclient.Options) (http.RoundTripper, error)
}

type datasourceDecrypter interface {
	DecryptedValues(ctx context.Context, ds *datasources.DataSource) (map[string]string, error)
}

type PrometheusWriterFactory struct {
	cfg                 PrometheusWriterConfig
	dsCache             datasourceCache
	datasourceDecrypter datasourceDecrypter
	httpClientProvider  httpClientProvider
	logger              log.Logger
}

func NewPrometheusWriterFactory(
	cfg PrometheusWriterConfig,
	dsCache datasourceCache,
	datasourceDecrypter datasourceDecrypter,
	httpClientProvider httpClientProvider,
	l log.Logger,
) *PrometheusWriterFactory {
	return &PrometheusWriterFactory{
		cfg:                 cfg,
		dsCache:             dsCache,
		datasourceDecrypter: datasourceDecrypter,
		httpClientProvider:  httpClientProvider,
		logger:              l,
	}
}

func (p PrometheusWriterFactory) GetWriter(ctx context.Context, orgID int64) (Writer, error) {
	bgUser := accesscontrol.BackgroundUser("alerting_recording_rules", orgID, org.RoleAdmin, []accesscontrol.Permission{
		{Action: datasources.ActionRead, Scope: datasources.ScopeAll},
	})
	ds, err := p.dsCache.GetDatasourceByUID(ctx, p.cfg.DatasourceUID, bgUser, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get datasource: %w", err)
	}

	decryptFn := func(ds *datasources.DataSource) (map[string]string, error) {
		return p.datasourceDecrypter.DecryptedValues(ctx, ds)
	}
	instanceSettings, err := adapters.ModelToInstanceSettings(ds, decryptFn)
	if err != nil {
		return nil, err
	}

	httpClientOptions, err := instanceSettings.HTTPClientOptions(ctx)
	if err != nil {
		return nil, err
	}

	rt, err := p.httpClientProvider.GetTransport(httpclient.Options{
		Timeouts:  httpClientOptions.Timeouts,
		TLS:       httpClientOptions.TLS,
		BasicAuth: httpClientOptions.BasicAuth,
	})
	if err != nil {
		return nil, err
	}

	instanceURL, err := url.Parse(instanceSettings.URL)
	if err != nil {
		return nil, err
	}
	instanceURL.Path = p.cfg.WritePath

	clientCfg := promremote.NewConfig(
		promremote.UserAgent("grafana-recording-rule"),
		promremote.WriteURLOption(instanceURL.String()),
		promremote.HTTPClientTimeoutOption(p.cfg.Timeout),
		promremote.HTTPClientOption(&http.Client{Transport: rt}),
	)

	client, err := promremote.NewClient(clientCfg)
	if err != nil {
		return nil, err
	}

	return &PrometheusWriter{
		OrgID:    orgID,
		TenantID: p.cfg.TenantID,
		client:   client,
		logger:   p.logger,
	}, nil
}

type PrometheusWriter struct {
	OrgID    int64
	TenantID string
	client   promremote.Client
	logger   log.Logger
}

// Write writes the given frames to the Prometheus remote write endpoint.
func (w PrometheusWriter) Write(ctx context.Context, name string, t time.Time, frames data.Frames, extraLabels map[string]string) error {
	l := w.logger.FromContext(ctx)

	points, err := PointsFromFrames(name, t, frames, extraLabels)
	if err != nil {
		return err
	}

	series := make([]promremote.TimeSeries, 0, len(points))
	for _, p := range points {
		series = append(series, promremote.TimeSeries{
			Labels: promremoteLabelsFromPoint(p),
			Datapoint: promremote.Datapoint{
				Timestamp: p.Metric.T,
				Value:     p.Metric.V,
			},
		})
	}

	headers := map[string]string{}
	if w.TenantID != "" {
		headers["X-Scope-OrgID"] = fmt.Sprintf("%v", w.OrgID)
	}

	l.Debug("Writing time series", "series", name)
	_, writeErr := w.client.WriteTimeSeries(ctx, series, promremote.WriteOptions{
		Headers: headers,
	})
	if err := checkWriteError(writeErr); err != nil {
		return fmt.Errorf("failed to write time series: %w", err)
	}

	return nil
}

func promremoteLabelsFromPoint(point Point) []promremote.Label {
	labels := make([]promremote.Label, 0, len(point.Labels))
	labels = append(labels, promremote.Label{
		Name:  "__name__",
		Value: point.Name,
	})
	for k, v := range point.Labels {
		labels = append(labels, promremote.Label{
			Name:  k,
			Value: v,
		})
	}
	return labels
}

func checkWriteError(writeErr promremote.WriteError) error {
	if writeErr == nil {
		return nil
	}

	// special case for 400 status code
	if writeErr.StatusCode() == 400 {
		msg := writeErr.Error()
		for _, e := range DuplicateTimestampErrors {
			if strings.Contains(msg, e) {
				return nil
			}
		}
	}

	return writeErr
}
