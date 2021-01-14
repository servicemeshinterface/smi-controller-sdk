package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	"github.com/stretchr/testify/mock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type V1Alpha2Mock struct {
	mock.Mock
}

func (v *V1Alpha2Mock) UpsertTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha2.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}

func (v *V1Alpha2Mock) DeleteTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha2.TrafficTarget) (ctrl.Result, error) {

	v.Called(ctx, r, l, tt)

	return ctrl.Result{}, nil
}
