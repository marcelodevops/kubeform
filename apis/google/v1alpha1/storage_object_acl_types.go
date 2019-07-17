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

type StorageObjectAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageObjectAclSpec   `json:"spec,omitempty"`
	Status            StorageObjectAclStatus `json:"status,omitempty"`
}

type StorageObjectAclSpec struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
	// +optional
	PredefinedAcl string `json:"predefined_acl,omitempty"`
}

type StorageObjectAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageObjectAclList is a list of StorageObjectAcls
type StorageObjectAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageObjectAcl CRD objects
	Items []StorageObjectAcl `json:"items,omitempty"`
}