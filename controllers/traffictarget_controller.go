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
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	accessv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha1"
)

// TrafficTargetReconciler reconciles a TrafficTarget object
type TrafficTargetReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets/status,verbs=get;update;patch

func (r *TrafficTargetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("traffictarget", req.NamespacedName)

	tt := &accessv1alpha1.TrafficTarget{}
	if err := r.Get(ctx, req.NamespacedName, tt); err != nil {
		r.Log.Info("unable to fetch TrafficTarget, most likely deleted")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	ttFinalizerName := "traffictarget.finalizers.smi-controller"

	// examine DeletionTimestamp to determine if object is under deletion
	if tt.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !containsString(tt.ObjectMeta.Finalizers, ttFinalizerName) {
			tt.ObjectMeta.Finalizers = append(tt.ObjectMeta.Finalizers, ttFinalizerName)
			if err := r.Update(context.Background(), tt); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if containsString(tt.ObjectMeta.Finalizers, ttFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			sdk.API().V1Alpha().DeleteTrafficTarget(ctx, r.Client, r.Log, tt)

			// remove our finalizer from the list and update it.
			tt.ObjectMeta.Finalizers = removeString(tt.ObjectMeta.Finalizers, ttFinalizerName)
			if err := r.Update(context.Background(), tt); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	return sdk.API().V1Alpha().UpsertTrafficTarget(ctx, r.Client, r.Log, tt)
}

func (r *TrafficTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&accessv1alpha1.TrafficTarget{}).
		Complete(r)
}

// Helper functions to check and remove string from a slice of strings.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
