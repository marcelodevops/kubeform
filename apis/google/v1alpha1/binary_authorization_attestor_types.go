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

type BinaryAuthorizationAttestor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              BinaryAuthorizationAttestorSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            BinaryAuthorizationAttestorStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys struct {
	AsciiArmoredPgpPublicKey string `json:"asciiArmoredPgpPublicKey" tf:"ascii_armored_pgp_public_key" protobuf:"bytes,1,opt,name=asciiArmoredPgpPublicKey"`
	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty" protobuf:"bytes,2,opt,name=comment"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty" protobuf:"bytes,3,opt,name=ID"`
}

type BinaryAuthorizationAttestorSpecAttestationAuthorityNote struct {
	// +optional
	DelegationServiceAccountEmail string `json:"delegationServiceAccountEmail,omitempty" tf:"delegation_service_account_email,omitempty" protobuf:"bytes,1,opt,name=delegationServiceAccountEmail"`
	NoteReference                 string `json:"noteReference" tf:"note_reference" protobuf:"bytes,2,opt,name=noteReference"`
	// +optional
	PublicKeys []BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys `json:"publicKeys,omitempty" tf:"public_keys,omitempty" protobuf:"bytes,3,rep,name=publicKeys"`
}

type BinaryAuthorizationAttestorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +kubebuilder:validation:MaxItems=1
	AttestationAuthorityNote []BinaryAuthorizationAttestorSpecAttestationAuthorityNote `json:"attestationAuthorityNote" tf:"attestation_authority_note" protobuf:"bytes,3,rep,name=attestationAuthorityNote"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,4,opt,name=description"`
	Name        string `json:"name" tf:"name" protobuf:"bytes,5,opt,name=name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,6,opt,name=project"`
}

type BinaryAuthorizationAttestorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *BinaryAuthorizationAttestorSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BinaryAuthorizationAttestorList is a list of BinaryAuthorizationAttestors
type BinaryAuthorizationAttestorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of BinaryAuthorizationAttestor CRD objects
	Items []BinaryAuthorizationAttestor `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
