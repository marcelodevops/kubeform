/*
Copyright The Kubeform Authors.

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeSSLPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ComputeSSLPolicySpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ComputeSSLPolicyStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ComputeSSLPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty" protobuf:"bytes,3,opt,name=creationTimestamp"`
	// +optional
	CustomFeatures []string `json:"customFeatures,omitempty" tf:"custom_features,omitempty" protobuf:"bytes,4,rep,name=customFeatures"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,5,opt,name=description"`
	// +optional
	EnabledFeatures []string `json:"enabledFeatures,omitempty" tf:"enabled_features,omitempty" protobuf:"bytes,6,rep,name=enabledFeatures"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty" protobuf:"bytes,7,opt,name=fingerprint"`
	// +optional
	MinTLSVersion string `json:"minTLSVersion,omitempty" tf:"min_tls_version,omitempty" protobuf:"bytes,8,opt,name=minTLSVersion"`
	Name          string `json:"name" tf:"name" protobuf:"bytes,9,opt,name=name"`
	// +optional
	Profile string `json:"profile,omitempty" tf:"profile,omitempty" protobuf:"bytes,10,opt,name=profile"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,11,opt,name=project"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty" protobuf:"bytes,12,opt,name=selfLink"`
}

type ComputeSSLPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ComputeSSLPolicySpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSSLPolicyList is a list of ComputeSSLPolicys
type ComputeSSLPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ComputeSSLPolicy CRD objects
	Items []ComputeSSLPolicy `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
