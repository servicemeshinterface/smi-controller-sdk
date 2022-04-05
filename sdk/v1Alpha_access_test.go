package sdk

import (
	"testing"

	accessv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha4"
	"github.com/stretchr/testify/mock"
)

func TestCallsUserDefinedUpsertTrafficTargetWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertTrafficTarget(nil, nil, l, &accessv1alpha4.TrafficTarget{})

	v1.AssertCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertTrafficTargetWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertTrafficTarget(nil, nil, l, &accessv1alpha4.TrafficTarget{})

	v1.AssertNotCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestCallsUserDefinedUpsertIdentityBindingWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertIdentityBinding(nil, nil, l, &accessv1alpha4.IdentityBinding{})

	v1.AssertCalled(t, "UpsertIdentityBinding", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertIdentityBindingWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertIdentityBinding(nil, nil, l, &accessv1alpha4.IdentityBinding{})

	v1.AssertNotCalled(t, "UpsertIdentityBinding", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}
