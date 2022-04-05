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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type PodLabelSelectorScheme struct {
	// +kubebuilder:validation:MinProperties=1
	// +nullable
	MatchLabels map[string]string `json:"matchLabels"`
}

// IdentityBinding is composed of a set of schemes that describe
// a service's identity.
// +kubebuilder:validation:MinProperties=1
type IdentityBindingSchemes struct {
	// +kubebuilder:validation:Optional
	// +nullable
	PodLabelSelector PodLabelSelectorScheme `json:"podLabelSelector,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	// +nullable
	SPIFFEIdentities []string `json:"spiffeIdentities,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	ServiceAccount string `json:"serviceAccount,omitempty"`
}

// IdentityBindingSpec defines the desired state of IdentityBinding
type IdentityBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Schemes IdentityBindingSchemes `json:"schemes"`
}

// IdentityBindingStatus defines the observed state of IdentityBinding
type IdentityBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=ib

// An `IdentityBinding` declares the set of identities belonging to a particular workload
// for the purposes of policy (i.e. TrafficTarget).
type IdentityBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityBindingSpec   `json:"spec,omitempty"`
	Status IdentityBindingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IdentityBindingList contains a list of IdentityBinding
type IdentityBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IdentityBinding{}, &IdentityBindingList{})
}
