package access

import (
	"context"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha4"
	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func testCreateIdentityBinding(t *testing.T) {
	ib := &v1alpha4.IdentityBinding{}
	copier.Copy(ib, &v4IdentityBinding)

	ctx := context.Background()
	err := k8sClient.Create(ctx, ib)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "UpsertIdentityBinding", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func testDeleteIdentityBinding(t *testing.T) {
	ib := &v1alpha4.IdentityBinding{}
	copier.Copy(ib, &v4IdentityBinding)

	ctx := context.Background()
	err := k8sClient.Delete(ctx, ib)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "DeleteIdentityBinding", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

var v4IdentityBinding = &v1alpha4.IdentityBinding{
	TypeMeta: v1.TypeMeta{
		Kind:       "IdentityBinding",
		APIVersion: v1alpha4.GroupVersion.Identifier(),
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4access",
		Namespace: "default",
	},
	Spec: v1alpha4.IdentityBindingSpec{
		Schemes: v1alpha4.IdentityBindingSchemes{
			ServiceAccount: "default",
		},
	},
}
