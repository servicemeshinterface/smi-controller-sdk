/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/nicholasjackson/smi-controller-sdk/sdk"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
)

// TrafficTargetReconciler reconciles a TrafficTarget object
type TrafficSplitReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets/status,verbs=get;update;patch

func (r *TrafficSplitReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("trafficsplit", req.NamespacedName)

	ts := &splitv1alpha1.TrafficSplit{}
	if err := r.Get(ctx, req.NamespacedName, ts); err != nil {
		r.Log.Info("unable to fetch TrafficSplit, most likely deleted")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	tsFinalizerName := "trafficsplit.finalizers.smi-controller"

	// examine DeletionTimestamp to determine if object is under deletion
	if ts.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !containsString(ts.ObjectMeta.Finalizers, tsFinalizerName) {
			ts.ObjectMeta.Finalizers = append(ts.ObjectMeta.Finalizers, tsFinalizerName)
			if err := r.Update(context.Background(), ts); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if containsString(ts.ObjectMeta.Finalizers, tsFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			sdk.API().V1Alpha().DeleteTrafficSplit(ctx, r.Client, r.Log, ts)

			// remove our finalizer from the list and update it.
			ts.ObjectMeta.Finalizers = removeString(ts.ObjectMeta.Finalizers, tsFinalizerName)
			if err := r.Update(context.Background(), ts); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	return sdk.API().V1Alpha().UpsertTrafficSplit(ctx, r.Client, r.Log, ts)
}

func (r *TrafficSplitReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&splitv1alpha1.TrafficSplit{}).
		Complete(r)
}
