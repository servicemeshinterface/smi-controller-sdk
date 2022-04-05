package v1alpha1

import (
	"testing"

	"github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToConvertsFromAlpha4ToAlpha1(t *testing.T) {
	v1Test := &TrafficTarget{}

	err := v1Test.ConvertFrom(v4Access)
	assert.NoError(t, err)

	assert.Equal(t, v4Access.ObjectMeta, v1Test.ObjectMeta)
	assert.Equal(t, v4Access.TypeMeta.Kind, v1Test.TypeMeta.Kind)
	assert.Equal(t, GroupVersion.Identifier(), v1Test.TypeMeta.APIVersion)

	// test detination
	assert.Equal(t, v4Access.Spec.Destination.Kind, v1Test.Destination.Kind)
	assert.Equal(t, v4Access.Spec.Destination.Name, v1Test.Destination.Name)
	assert.Equal(t, v4Access.Spec.Destination.Namespace, v1Test.Destination.Namespace)

	// test sources
	assert.Len(t, v1Test.Sources, len(v4Access.Spec.Sources))
	for i, s := range v4Access.Spec.Sources {
		assert.Equal(t, s.Kind, v1Test.Sources[i].Kind)
		assert.Equal(t, s.Name, v1Test.Sources[i].Name)
		assert.Equal(t, s.Namespace, v1Test.Sources[i].Namespace)
	}

	// test rules
	assert.Len(t, v1Test.Specs, len(v4Access.Spec.Rules))
	for i, s := range v4Access.Spec.Rules {
		assert.Equal(t, s.Kind, v1Test.Specs[i].Kind)
		assert.Equal(t, s.Name, v1Test.Specs[i].Name)

		for n, m := range s.Matches {
			assert.Equal(t, m, v1Test.Specs[i].Matches[n])
		}
	}
}

func TestConvertToConvertsFromAlpha1ToAlpha4(t *testing.T) {
	v4Test := &v1alpha4.TrafficTarget{}

	err := v1Access.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v1Access.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v1Access.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, v1alpha4.GroupVersion.Identifier(), v4Test.TypeMeta.APIVersion)

	// test destination
	assert.Equal(t, v1Access.Destination.Kind, v4Test.Spec.Destination.Kind)
	assert.Equal(t, v1Access.Destination.Name, v4Test.Spec.Destination.Name)
	assert.Equal(t, v1Access.Destination.Namespace, v4Test.Spec.Destination.Namespace)

	// test sources
	assert.Len(t, v4Test.Spec.Sources, len(v1Access.Sources))
	for i, s := range v1Access.Sources {
		assert.Equal(t, s.Kind, v4Test.Spec.Sources[i].Kind)
		assert.Equal(t, s.Name, v4Test.Spec.Sources[i].Name)
		assert.Equal(t, s.Namespace, v4Test.Spec.Sources[i].Namespace)
	}

	// test rules
	assert.Len(t, v4Test.Spec.Rules, len(v1Access.Specs))
	for i, s := range v1Access.Specs {
		assert.Equal(t, s.Kind, v4Test.Spec.Rules[i].Kind)
		assert.Equal(t, s.Name, v4Test.Spec.Rules[i].Name)

		for n, m := range s.Matches {
			assert.Equal(t, m, v4Test.Spec.Rules[i].Matches[n])
		}
	}
}

var v4Access = &v1alpha4.TrafficTarget{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: "v1alpha4",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4Access",
		Namespace: "default",
	},
	Spec: v1alpha4.TrafficTargetSpec{
		Destination: v1alpha4.IdentityBindingSubject{
			Kind:      "ServiceAccount",
			Name:      "myservice",
			Namespace: "default",
		},
		Sources: []v1alpha4.IdentityBindingSubject{
			{
				Kind:      "ServiceAccount",
				Name:      "mydestination1",
				Namespace: "default",
			},
			{
				Kind:      "ServiceAccount",
				Name:      "mydestination2",
				Namespace: "default",
			},
		},
		Rules: []v1alpha4.TrafficTargetRule{
			{
				Kind:    "HTTPRouteGroup",
				Name:    "myname1",
				Matches: []string{"abc", "124"},
			},
			{
				Kind:    "HTTPRouteGroup",
				Name:    "myname2",
				Matches: []string{"124", "abc"},
			},
		},
	},
}

var v1Access = &TrafficTarget{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: "v1alpha1",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v1Access",
		Namespace: "default",
	},
	Destination: IdentityBindingSubject{
		Kind:      "ServiceAccount",
		Name:      "myservice",
		Namespace: "default",
	},
	Sources: []IdentityBindingSubject{
		{
			Kind:      "ServiceAccount",
			Name:      "mydestination1",
			Namespace: "default",
		},
		{
			Kind:      "ServiceAccount",
			Name:      "mydestination2",
			Namespace: "default",
		},
	},
	Specs: []TrafficTargetSpec{
		{
			Kind:    "HTTPRouteGroup",
			Name:    "myname1",
			Matches: []string{"abc", "124"},
		},
		{
			Kind:    "HTTPRouteGroup",
			Name:    "myname2",
			Matches: []string{"124", "abc"},
		},
	},
}
