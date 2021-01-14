package sdk

import (
	"testing"

	"github.com/go-logr/logr"
	accessv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	"github.com/stretchr/testify/mock"
	ctrl "sigs.k8s.io/controller-runtime"
)

func setupSDKTests(t *testing.T) (*V1Alpha2Mock, *api, logr.Logger) {
	v1 := &V1Alpha2Mock{}
	v1.On("UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	a := &api{&v1Alpha2Impl{}}
	a.RegisterV1Alpha2(v1)

	l := ctrl.Log.WithName("controllers").WithName("TrafficTarget")

	return v1, a, l
}

func TestCallsUserDefinedUpsertTrafficTargetWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha2().UpsertTrafficTarget(nil, nil, l, &accessv1alpha2.TrafficTarget{})

	v1.AssertCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertTrafficTargetWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha2(nil)

	a.V1Alpha2().UpsertTrafficTarget(nil, nil, l, &accessv1alpha2.TrafficTarget{})

	v1.AssertNotCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}
