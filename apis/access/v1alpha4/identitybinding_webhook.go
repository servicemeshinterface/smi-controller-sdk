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

package v1alpha4

import (
	"fmt"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var identitybindinglog = logf.Log.WithName("identitybinding-resource")

func (r *IdentityBinding) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/validate-access-smi-spec-io-v1alpha4-identitybinding,mutating=false,failurePolicy=fail,sideEffects=None,groups=access.smi-spec.io,resources=identitybindings,verbs=create;update,versions=v1alpha4,name=videntitybinding.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &IdentityBinding{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *IdentityBinding) ValidateCreate() error {
	identitybindinglog.Info("validate create", "name", r.Name)

	return r.validateIdentityBinding()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *IdentityBinding) ValidateUpdate(old runtime.Object) error {
	identitybindinglog.Info("validate update", "name", r.Name)

	return r.validateIdentityBinding()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
// We don't need validation on delete so just return nil
func (r *IdentityBinding) ValidateDelete() error {
	return nil
}

func (r *IdentityBinding) validateIdentityBinding() error {
	var allErrs field.ErrorList
	if err := r.validatePodLabelSelector(); err != nil {
		allErrs = append(allErrs, err)
	}

	if errs := r.valdiateSPIFFEIdentities(); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if err := r.validateServiceAccount(); err != nil {
		allErrs = append(allErrs, err)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "access.smi-spec.io", Kind: "IdentityBinding"},
		r.Name, allErrs)
}

func (r *IdentityBinding) validatePodLabelSelector() *field.Error {
	selector := r.Spec.Schemes.PodLabelSelector

	if selector.MatchLabels != nil && r.Spec.Schemes.ServiceAccount != "" {
		return field.Invalid(field.NewPath("spec").Child("schemes").Child("podLabelSelector"), selector, "podLabelSelector and serviceAccount are mutually exclusive")
	}

	return nil
}

func (r *IdentityBinding) valdiateSPIFFEIdentities() []*field.Error {
	var allErrs []*field.Error
	ids := r.Spec.Schemes.SPIFFEIdentities

	for idx, id := range ids {
		_, err := spiffeid.FromString(id)
		if err != nil {
			allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("schemes").Child("spiffeIdentities").Index(idx), id, fmt.Sprintf("invalid spffieID: %s", err)))
		}
	}

	if len(allErrs) == 0 {
		return nil
	}

	return allErrs
}

func (r *IdentityBinding) validateServiceAccount() *field.Error {
	sa := r.Spec.Schemes.ServiceAccount

	if sa != "" && r.Spec.Schemes.PodLabelSelector.MatchLabels != nil {
		return field.Invalid(field.NewPath("spec").Child("schemes").Child("serviceAccount"), sa, "serviceAccount and podLabelSelector are mutually exclusive")
	}

	return nil
}
