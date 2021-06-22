package logging

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Logger struct{}

func (l *Logger) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha3.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("UpsertTrafficTarget", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}

func (l *Logger) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha3.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("DeleteTrafficTarget", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}

func (l *Logger) UpsertTrafficSplit(
	ctx context.Context,
	r client.Client,
	log logr.Logger,
	tt *splitv1alpha4.TrafficSplit) (ctrl.Result, error) {

	log.Info("UpdateTrafficSplit", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}

func (l *Logger) DeleteTrafficSplit(
	ctx context.Context,
	r client.Client,
	log logr.Logger,
	tt *splitv1alpha4.TrafficSplit) (ctrl.Result, error) {

	log.Info("DeleteTrafficSplit", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}
