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

type EventgridTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              EventgridTopicSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            EventgridTopicStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type EventgridTopicSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty" protobuf:"bytes,4,opt,name=endpoint"`
	Location string `json:"location" tf:"location" protobuf:"bytes,5,opt,name=location"`
	Name     string `json:"name" tf:"name" protobuf:"bytes,6,opt,name=name"`
	// +optional
	PrimaryAccessKey  string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,7,opt,name=resourceGroupName"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,8,rep,name=tags"`
}

type EventgridTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *EventgridTopicSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventgridTopicList is a list of EventgridTopics
type EventgridTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of EventgridTopic CRD objects
	Items []EventgridTopic `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
