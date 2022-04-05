package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha4"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
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
	tt *accessv1alpha4.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha4.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) UpsertIdentityBinding(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha4.IdentityBinding) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteIdentityBinding(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha4.IdentityBinding) (ctrl.Result, error) {

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

func (v *V1AlphaMock) UpsertHTTPRouteGroup(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *specsv1alpha4.HTTPRouteGroup) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteHTTPRouteGroup(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *specsv1alpha4.HTTPRouteGroup) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) UpsertTCPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *specsv1alpha4.TCPRoute) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteTCPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *specsv1alpha4.TCPRoute) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) UpsertUDPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *specsv1alpha4.UDPRoute) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteUDPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *specsv1alpha4.UDPRoute) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}
