package sdk

import (
	"testing"

	"github.com/go-logr/logr"
	accessv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha1"
	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	"github.com/stretchr/testify/mock"
	ctrl "sigs.k8s.io/controller-runtime"
)

func setupSDKTests(t *testing.T) (*V1AlphaMock, *api, logr.Logger) {
	v1 := &V1AlphaMock{}
	v1.On("UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("UpsertTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	a := &api{&v1AlphaImpl{}}
	a.RegisterV1Alpha(v1)

	l := ctrl.Log.WithName("controllers").WithName("TrafficTarget")

	return v1, a, l
}

func TestCallsUserDefinedUpsertTrafficTargetWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertTrafficTarget(nil, nil, l, &accessv1alpha1.TrafficTarget{})

	v1.AssertCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertTrafficTargetWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertTrafficTarget(nil, nil, l, &accessv1alpha1.TrafficTarget{})

	v1.AssertNotCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestCallsUserDefinedUpsertTrafficSplitWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertTrafficSplit(nil, nil, l, &splitv1alpha1.TrafficSplit{})

	v1.AssertCalled(t, "UpsertTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertTrafficSplitWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertTrafficSplit(nil, nil, l, &splitv1alpha1.TrafficSplit{})

	v1.AssertNotCalled(t, "UpsertTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}
