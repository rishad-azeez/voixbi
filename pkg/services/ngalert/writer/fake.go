package writer

import (
	"context"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type FakeWriterProvider struct {
	GetWriterFunc func(ctx context.Context, orgID int64) (Writer, error)
}

func (p FakeWriterProvider) GetWriter(ctx context.Context, orgID int64) (Writer, error) {
	if p.GetWriterFunc == nil {
		return &FakeWriter{}, nil
	}

	return p.GetWriterFunc(ctx, orgID)
}

type FakeWriter struct {
	WriteFunc func(ctx context.Context, name string, t time.Time, frames data.Frames, extraLabels map[string]string) error
}

func (w FakeWriter) Write(ctx context.Context, name string, t time.Time, frames data.Frames, extraLabels map[string]string) error {
	if w.WriteFunc == nil {
		return nil
	}

	return w.WriteFunc(ctx, name, t, frames, extraLabels)
}
