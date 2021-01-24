package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha1"
	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// UpsertTrafficTarget defines a callback function for updating or
// inserting a new TrafficTarget
type UpsertTrafficTarget interface {
	UpsertTrafficTarget(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha1.TrafficTarget) (ctrl.Result, error)
}

// DeleteTrafficTarget defines a callback function for deleting
// a new TrafficTarget
type DeleteTrafficTarget interface {
	DeleteTrafficTarget(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha1.TrafficTarget) (ctrl.Result, error)
}

// UpsertTrafficSplit defines a callback function for updating or
// inserting a new TrafficSplit
type UpsertTrafficSplit interface {
	UpsertTrafficSplit(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *splitv1alpha1.TrafficSplit) (ctrl.Result, error)
}

// DeleteTrafficSplit defines a callback function for deleting
// a new TrafficSplit
type DeleteTrafficSplit interface {
	DeleteTrafficSplit(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *splitv1alpha1.TrafficSplit) (ctrl.Result, error)
}

// V1Alpha defines an interface containing callback methods for the v1alpha2 API
// We define the methods as individual interfaces as we want to enable the user to
// implement only the callbacks they need
type V1Alpha interface {
	UpsertTrafficTarget
	DeleteTrafficTarget
	UpsertTrafficSplit
	DeleteTrafficSplit
}

// v1Alpha2Impl is a concrete implementation of the V1Alpha2 interface
type v1AlphaImpl struct {
	userV1alpha interface{}
}

// RegisterV1Alpha2 registers user defined callback functions to the api
func (a *v1AlphaImpl) RegisterV1Alpha(i interface{}) {
	a.userV1alpha = i
}

// UpsertTrafficTarget will call the user defined UpsertTrafficTarget callback
// when defined
func (a *v1AlphaImpl) UpsertTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertTrafficTarget)

	if !ok {
		l.Info("Client code does not implement UpsertTrafficTarget")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertTrafficTarget(ctx, r, l, tt)
}

func (a *v1AlphaImpl) DeleteTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteTrafficTarget)

	if !ok {
		l.Info("Client code does not implement DeleteTrafficTarget")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteTrafficTarget(ctx, r, l, tt)
}

// UpsertTrafficSplit will call the user defined UpsertTrafficSplit callback
// when defined
func (a *v1AlphaImpl) UpsertTrafficSplit(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *splitv1alpha1.TrafficSplit,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertTrafficSplit)

	if !ok {
		l.Info("Client code does not implement UpsertTrafficSplit")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertTrafficSplit(ctx, r, l, tt)
}

func (a *v1AlphaImpl) DeleteTrafficSplit(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *splitv1alpha1.TrafficSplit,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteTrafficSplit)

	if !ok {
		l.Info("Client code does not implement DeleteTrafficSplit")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteTrafficSplit(ctx, r, l, tt)
}
