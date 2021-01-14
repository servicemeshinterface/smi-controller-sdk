package mesh

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type loggerV2 struct {
}

func NewLogger() *API {
	a := &API{}
	a.RegisterV1Alpha2(newLoggerV1Alpha2())

	return a
}

func newLoggerV1Alpha2() *loggerV2 {
	return &loggerV2{}
}

func (l *loggerV2) UpsertTrafficTarget(ctx context.Context, c client.Client, log logr.Logger, tt *accessv1alpha2.TrafficTarget) (ctrl.Result, error) {
	log.Info("UpsertTrafficTarget", "api", "v1alpha2", "target", tt)

	return ctrl.Result{}, nil
}

func (l *loggerV2) DeleteTrafficTarget(ctx context.Context, c client.Client, log logr.Logger, tt *accessv1alpha2.TrafficTarget) (ctrl.Result, error) {
	log.Info("DeleteTrafficTarget", "api", "v1alpha2", "target", tt)

	return ctrl.Result{}, nil
}
