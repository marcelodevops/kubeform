package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type OrganizationIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationIamBindingSpec   `json:"spec,omitempty"`
	Status            OrganizationIamBindingStatus `json:"status,omitempty"`
}

type OrganizationIamBindingSpec struct {
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members"`
	OrgId   string   `json:"org_id"`
	Role    string   `json:"role"`
}

type OrganizationIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationIamBindingList is a list of OrganizationIamBindings
type OrganizationIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationIamBinding CRD objects
	Items []OrganizationIamBinding `json:"items,omitempty"`
}