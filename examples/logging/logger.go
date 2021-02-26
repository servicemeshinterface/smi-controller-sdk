package logging

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha1"
	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Logger struct{}

func (l *Logger) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("UpsertTrafficTarget", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}

func (l *Logger) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("DeleteTrafficTarget", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}

func (l *Logger) UpsertTrafficSplit(
	ctx context.Context,
	r client.Client,
	log logr.Logger,
	tt *splitv1alpha1.TrafficSplit) (ctrl.Result, error) {

	log.Info("UpdateTrafficSplit", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}

func (l *Logger) DeleteTrafficSplit(
	ctx context.Context,
	r client.Client,
	log logr.Logger,
	tt *splitv1alpha1.TrafficSplit) (ctrl.Result, error) {

	log.Info("DeleteTrafficSplit", "api", "v1alpha1", "target", tt)

	return ctrl.Result{}, nil
}
