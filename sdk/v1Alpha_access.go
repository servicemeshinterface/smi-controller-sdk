package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha4"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// V1AlphaAccess defines an interface containing callback methods for the v1alpha4 API
// traffic access objects
type v1AlphaAccess interface {
	UpsertTrafficTarget
	DeleteTrafficTarget
	UpsertIdentityBinding
	DeleteIdentityBinding
}

// UpsertTrafficTarget defines a callback function for updating or
// inserting a new TrafficTarget
type UpsertTrafficTarget interface {
	UpsertTrafficTarget(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha4.TrafficTarget) (ctrl.Result, error)
}

// DeleteTrafficTarget defines a callback function for deleting
// a new TrafficTarget
type DeleteTrafficTarget interface {
	DeleteTrafficTarget(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha4.TrafficTarget) (ctrl.Result, error)
}

// UpsertIdentityBinding defines a callback function for updating or
// inserting a new IdentityBinding
type UpsertIdentityBinding interface {
	UpsertIdentityBinding(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha4.IdentityBinding) (ctrl.Result, error)
}

// DeleteIdentityBinding defines a callback function for deleting
// a new IdentityBinding
type DeleteIdentityBinding interface {
	DeleteIdentityBinding(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha4.IdentityBinding) (ctrl.Result, error)
}

// UpsertTrafficTarget will call the user defined UpsertTrafficTarget callback
// when defined
func (a *v1AlphaImpl) UpsertTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha4.TrafficTarget,
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
	tt *accessv1alpha4.TrafficTarget,
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

// UpsertIdentityBinding will call the user defined UpsertIdentityBinding callback
// when defined
func (a *v1AlphaImpl) UpsertIdentityBinding(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	ib *accessv1alpha4.IdentityBinding,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertIdentityBinding)

	if !ok {
		l.Info("Client code does not implement UpsertIdentityBinding")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertIdentityBinding(ctx, r, l, ib)
}

func (a *v1AlphaImpl) DeleteIdentityBinding(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	ib *accessv1alpha4.IdentityBinding,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteIdentityBinding)

	if !ok {
		l.Info("Client code does not implement DeleteIdentityBinding")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteIdentityBinding(ctx, r, l, ib)
}
