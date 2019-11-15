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

type FolderOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              FolderOrganizationPolicySpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            FolderOrganizationPolicyStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type FolderOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced" tf:"enforced" protobuf:"varint,1,opt,name=enforced"`
}

type FolderOrganizationPolicySpecListPolicyAllow struct {
	// +optional
	All bool `json:"all,omitempty" tf:"all,omitempty" protobuf:"varint,1,opt,name=all"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values,omitempty" protobuf:"bytes,2,rep,name=values"`
}

type FolderOrganizationPolicySpecListPolicyDeny struct {
	// +optional
	All bool `json:"all,omitempty" tf:"all,omitempty" protobuf:"varint,1,opt,name=all"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values,omitempty" protobuf:"bytes,2,rep,name=values"`
}

type FolderOrganizationPolicySpecListPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Allow []FolderOrganizationPolicySpecListPolicyAllow `json:"allow,omitempty" tf:"allow,omitempty" protobuf:"bytes,1,rep,name=allow"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deny []FolderOrganizationPolicySpecListPolicyDeny `json:"deny,omitempty" tf:"deny,omitempty" protobuf:"bytes,2,rep,name=deny"`
	// +optional
	SuggestedValue string `json:"suggestedValue,omitempty" tf:"suggested_value,omitempty" protobuf:"bytes,3,opt,name=suggestedValue"`
}

type FolderOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default" tf:"default" protobuf:"varint,1,opt,name=default"`
}

type FolderOrganizationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	BooleanPolicy []FolderOrganizationPolicySpecBooleanPolicy `json:"booleanPolicy,omitempty" tf:"boolean_policy,omitempty" protobuf:"bytes,3,rep,name=booleanPolicy"`
	Constraint    string                                      `json:"constraint" tf:"constraint" protobuf:"bytes,4,opt,name=constraint"`
	// +optional
	Etag   string `json:"etag,omitempty" tf:"etag,omitempty" protobuf:"bytes,5,opt,name=etag"`
	Folder string `json:"folder" tf:"folder" protobuf:"bytes,6,opt,name=folder"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ListPolicy []FolderOrganizationPolicySpecListPolicy `json:"listPolicy,omitempty" tf:"list_policy,omitempty" protobuf:"bytes,7,rep,name=listPolicy"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RestorePolicy []FolderOrganizationPolicySpecRestorePolicy `json:"restorePolicy,omitempty" tf:"restore_policy,omitempty" protobuf:"bytes,8,rep,name=restorePolicy"`
	// +optional
	UpdateTime string `json:"updateTime,omitempty" tf:"update_time,omitempty" protobuf:"bytes,9,opt,name=updateTime"`
	// +optional
	Version int64 `json:"version,omitempty" tf:"version,omitempty" protobuf:"varint,10,opt,name=version"`
}

type FolderOrganizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *FolderOrganizationPolicySpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FolderOrganizationPolicyList is a list of FolderOrganizationPolicys
type FolderOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of FolderOrganizationPolicy CRD objects
	Items []FolderOrganizationPolicy `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
