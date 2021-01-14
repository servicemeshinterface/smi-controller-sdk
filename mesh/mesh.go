package mesh

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// API defines an object containing functions that the
// specific service meshes must implement to use this controller
type API struct {
	v1alpha2 V1Alpha2
}

func (a *API) RegisterV1Alpha2(i V1Alpha2) {
	a.v1alpha2 = i
}

func (a *API) V1Alpha2() V1Alpha2 {
	return a.v1alpha2
}

type V1Alpha2 interface {
	UpsertTrafficTarget(ctx context.Context, r client.Client, l logr.Logger, tt *accessv1alpha2.TrafficTarget) (ctrl.Result, error)
	DeleteTrafficTarget(ctx context.Context, r client.Client, l logr.Logger, tt *accessv1alpha2.TrafficTarget) (ctrl.Result, error)
}
