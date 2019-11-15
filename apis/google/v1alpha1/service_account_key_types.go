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

type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ServiceAccountKeySpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ServiceAccountKeyStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ServiceAccountKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty" protobuf:"bytes,4,opt,name=keyAlgorithm"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty" protobuf:"bytes,5,opt,name=name"`
	// +optional
	PgpKey string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty" protobuf:"bytes,6,opt,name=pgpKey"`
	// +optional
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key,omitempty"`
	// +optional
	PrivateKeyEncrypted string `json:"privateKeyEncrypted,omitempty" tf:"private_key_encrypted,omitempty" protobuf:"bytes,7,opt,name=privateKeyEncrypted"`
	// +optional
	PrivateKeyFingerprint string `json:"privateKeyFingerprint,omitempty" tf:"private_key_fingerprint,omitempty" protobuf:"bytes,8,opt,name=privateKeyFingerprint"`
	// +optional
	PrivateKeyType string `json:"privateKeyType,omitempty" tf:"private_key_type,omitempty" protobuf:"bytes,9,opt,name=privateKeyType"`
	// +optional
	PublicKey string `json:"publicKey,omitempty" tf:"public_key,omitempty" protobuf:"bytes,10,opt,name=publicKey"`
	// +optional
	PublicKeyType    string `json:"publicKeyType,omitempty" tf:"public_key_type,omitempty" protobuf:"bytes,11,opt,name=publicKeyType"`
	ServiceAccountID string `json:"serviceAccountID" tf:"service_account_id" protobuf:"bytes,12,opt,name=serviceAccountID"`
	// +optional
	ValidAfter string `json:"validAfter,omitempty" tf:"valid_after,omitempty" protobuf:"bytes,13,opt,name=validAfter"`
	// +optional
	ValidBefore string `json:"validBefore,omitempty" tf:"valid_before,omitempty" protobuf:"bytes,14,opt,name=validBefore"`
}

type ServiceAccountKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ServiceAccountKeySpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountKeyList is a list of ServiceAccountKeys
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ServiceAccountKey CRD objects
	Items []ServiceAccountKey `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
