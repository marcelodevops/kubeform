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

type LoadBalancerListenerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerListenerPolicySpec   `json:"spec,omitempty"`
	Status            LoadBalancerListenerPolicyStatus `json:"status,omitempty"`
}

type LoadBalancerListenerPolicySpec struct {
	LoadBalancerName string `json:"load_balancer_name"`
	LoadBalancerPort int    `json:"load_balancer_port"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PolicyNames []string `json:"policy_names,omitempty"`
}

type LoadBalancerListenerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoadBalancerListenerPolicyList is a list of LoadBalancerListenerPolicys
type LoadBalancerListenerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancerListenerPolicy CRD objects
	Items []LoadBalancerListenerPolicy `json:"items,omitempty"`
}