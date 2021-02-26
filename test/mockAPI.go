package main

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha1"
	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
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
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	args := ms.Called(ctx, c, log, tt)

	return args.Get(0).(ctrl.Result), args.Error(1)
}

func (ms *MockAPI) UpsertTrafficSplit(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	ts *splitv1alpha1.TrafficSplit,
) (ctrl.Result, error) {

	args := ms.Called(ctx, c, log, ts)

	log.Info("Upsert TrafficSplit called")

	return args.Get(0).(ctrl.Result), args.Error(1)
}
