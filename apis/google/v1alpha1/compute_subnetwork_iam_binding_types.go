package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeSubnetworkIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSubnetworkIamBindingSpec   `json:"spec,omitempty"`
	Status            ComputeSubnetworkIamBindingStatus `json:"status,omitempty"`
}

type ComputeSubnetworkIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag    string   `json:"etag,omitempty" tf:"etag,omitempty"`
	Members []string `json:"members" tf:"members"`
	// +optional
	// Deprecated
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// Deprecated
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	Role   string `json:"role" tf:"role"`
	// Deprecated
	Subnetwork string `json:"subnetwork" tf:"subnetwork"`
}

type ComputeSubnetworkIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeSubnetworkIamBindingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkIamBindingList is a list of ComputeSubnetworkIamBindings
type ComputeSubnetworkIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSubnetworkIamBinding CRD objects
	Items []ComputeSubnetworkIamBinding `json:"items,omitempty"`
}
