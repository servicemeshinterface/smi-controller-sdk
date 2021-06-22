package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type V1AlphaMock struct {
	mock.Mock
}

func (v *V1AlphaMock) UpsertTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha3.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha3.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) UpsertTrafficSplit(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *splitv1alpha4.TrafficSplit) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteTrafficSplit(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *splitv1alpha4.TrafficSplit) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}
