package v1alpha1

import (
	"testing"

	"github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToConvertsFromAlpha4ToAlpha2(t *testing.T) {
	v1Test := &TCPRoute{}

	err := v1Test.ConvertFrom(v4TCPRoute)
	assert.NoError(t, err)

	assert.Equal(t, v4TCPRoute.ObjectMeta, v1Test.ObjectMeta)
	assert.Equal(t, v4TCPRoute.TypeMeta.Kind, v1Test.TypeMeta.Kind)
	assert.Equal(t, GroupVersion.Identifier(), v1Test.TypeMeta.APIVersion)
}

func TestConvertToConvertsFromAlpha2ToAlpha4(t *testing.T) {
	v4Test := &v1alpha4.TCPRoute{}

	err := v1TCPRoute.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v1TCPRoute.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v1TCPRoute.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, v1alpha4.GroupVersion.Identifier(), v4Test.TypeMeta.APIVersion)

	// should have a blank matches as this does not exist in the v2 spec
	assert.Equal(t, v1alpha4.TCPRouteSpec{}, v4Test.Spec)
}

var v4TCPRoute = &v1alpha4.TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TCPRoute",
		APIVersion: "v1alpha4",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4Specs",
		Namespace: "default",
	},
	Spec: v1alpha4.TCPRouteSpec{
		Matches: v1alpha4.TCPMatch{
			Name:  "testing",
			Ports: []int{9090, 8080},
		},
	},
}

var v1TCPRoute = &TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TCPRoute",
		APIVersion: "v1alpha1",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v1Specs",
		Namespace: "default",
	},
}
