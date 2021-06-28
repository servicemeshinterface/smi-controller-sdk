package helpers

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MockAPI struct {
	mock.Mock
}

func (ms *MockAPI) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha3.TrafficTarget,
) (ctrl.Result, error) {

	args := ms.Called(ctx, c, log, tt)

	log.Info("Upsert TrafficTarget called")

	return args.Get(0).(ctrl.Result), args.Error(1)
}

func (ms *MockAPI) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha3.TrafficTarget,
) (ctrl.Result, error) {

	args := ms.Called(ctx, c, log, tt)

	log.Info("Upsert TrafficTarget called")

	return args.Get(0).(ctrl.Result), args.Error(1)
}

func (ms *MockAPI) UpsertTrafficSplit(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	ts *splitv1alpha4.TrafficSplit,
) (ctrl.Result, error) {

	args := ms.Called(ctx, c, log, ts)

	log.Info("Upsert TrafficSplit called")

	return args.Get(0).(ctrl.Result), args.Error(1)
}

func (ms *MockAPI) DeleteTrafficSplit(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	ts *splitv1alpha4.TrafficSplit,
) (ctrl.Result, error) {

	args := ms.Called(ctx, c, log, ts)

	log.Info("Upsert TrafficSplit called")

	return args.Get(0).(ctrl.Result), args.Error(1)
}
