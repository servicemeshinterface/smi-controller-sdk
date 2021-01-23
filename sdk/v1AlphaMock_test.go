package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha1"
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
	tt *accessv1alpha1.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1AlphaMock) DeleteTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha1.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}
