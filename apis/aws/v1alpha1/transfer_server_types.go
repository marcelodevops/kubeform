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

type TransferServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferServerSpec   `json:"spec,omitempty"`
	Status            TransferServerStatus `json:"status,omitempty"`
}

type TransferServerSpecEndpointDetails struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`
}

type TransferServerSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EndpointDetails *[]TransferServerSpec `json:"endpoint_details,omitempty"`
	// +optional
	EndpointType string `json:"endpoint_type,omitempty"`
	// +optional
	ForceDestroy bool `json:"force_destroy,omitempty"`
	// +optional
	IdentityProviderType string `json:"identity_provider_type,omitempty"`
	// +optional
	InvocationRole string `json:"invocation_role,omitempty"`
	// +optional
	LoggingRole string `json:"logging_role,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Url string `json:"url,omitempty"`
}

type TransferServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TransferServerList is a list of TransferServers
type TransferServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TransferServer CRD objects
	Items []TransferServer `json:"items,omitempty"`
}