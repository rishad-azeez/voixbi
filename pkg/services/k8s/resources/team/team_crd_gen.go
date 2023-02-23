// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     CRDTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package team

import (
	_ "embed"
	"fmt"

	"github.com/grafana/grafana/pkg/kinds/team"
	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/grafana/pkg/services/k8s/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var coreReg = corekind.NewBase(nil)

var CRD = crd.Kind{
	GrafanaKind: coreReg.Team(),
	Object:      &Team{},
	ObjectList:  &TeamList{},
}

// The CRD YAML representation of the Team kind.
//
//go:embed team.crd.yml
var CRDYaml []byte

// Team is the Go CRD representation of a single Team object.
// It implements [runtime.Object], and is used in k8s scheme construction.
type Team struct {
	crd.Base[team.Team]
}

// TeamList is the Go CRD representation of a list Team objects.
// It implements [runtime.Object], and is used in k8s scheme construction.
type TeamList struct {
	crd.ListBase[team.Team]
}

// fromUnstructured converts an *unstructured.Unstructured object to a *Team.
func fromUnstructured(obj any) (*Team, error) {
	uObj, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("failed to convert to *unstructured.Unstructured")
	}

	var team crd.Base[team.Team]
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(uObj.UnstructuredContent(), &team)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to Team: %w", err)
	}

	return &Team{team}, nil
}

// toUnstructured converts a Team to an *unstructured.Unstructured.
func toUnstructured(obj *team.Team, metadata metav1.ObjectMeta) (*unstructured.Unstructured, error) {
	teamObj := crd.Base[team.Team]{
		TypeMeta: metav1.TypeMeta{
			Kind:       CRD.GVK().Kind,
			APIVersion: CRD.GVK().Group + "/" + CRD.GVK().Version,
		},
		ObjectMeta: metadata,
		Spec:       *obj,
	}

	out, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&teamObj)
	if err != nil {
		return nil, err
	}

	return &unstructured.Unstructured{
		Object: out,
	}, nil
}
