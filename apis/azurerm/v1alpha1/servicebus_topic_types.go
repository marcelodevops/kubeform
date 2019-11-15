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

type ServicebusTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ServicebusTopicSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ServicebusTopicStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ServicebusTopicSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	AutoDeleteOnIdle string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty" protobuf:"bytes,3,opt,name=autoDeleteOnIdle"`
	// +optional
	DefaultMessageTtl string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty" protobuf:"bytes,4,opt,name=defaultMessageTtl"`
	// +optional
	DuplicateDetectionHistoryTimeWindow string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty" protobuf:"bytes,5,opt,name=duplicateDetectionHistoryTimeWindow"`
	// +optional
	EnableBatchedOperations bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty" protobuf:"varint,6,opt,name=enableBatchedOperations"`
	// +optional
	EnableExpress bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty" protobuf:"varint,7,opt,name=enableExpress"`
	// +optional
	// Deprecated
	EnableFilteringMessagesBeforePublishing bool `json:"enableFilteringMessagesBeforePublishing,omitempty" tf:"enable_filtering_messages_before_publishing,omitempty" protobuf:"varint,8,opt,name=enableFilteringMessagesBeforePublishing"`
	// +optional
	EnablePartitioning bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty" protobuf:"varint,9,opt,name=enablePartitioning"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty" protobuf:"bytes,10,opt,name=location"`
	// +optional
	MaxSizeInMegabytes int64  `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty" protobuf:"varint,11,opt,name=maxSizeInMegabytes"`
	Name               string `json:"name" tf:"name" protobuf:"bytes,12,opt,name=name"`
	NamespaceName      string `json:"namespaceName" tf:"namespace_name" protobuf:"bytes,13,opt,name=namespaceName"`
	// +optional
	RequiresDuplicateDetection bool   `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty" protobuf:"varint,14,opt,name=requiresDuplicateDetection"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,15,opt,name=resourceGroupName"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty" protobuf:"bytes,16,opt,name=status"`
	// +optional
	SupportOrdering bool `json:"supportOrdering,omitempty" tf:"support_ordering,omitempty" protobuf:"varint,17,opt,name=supportOrdering"`
}

type ServicebusTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ServicebusTopicSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusTopicList is a list of ServicebusTopics
type ServicebusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ServicebusTopic CRD objects
	Items []ServicebusTopic `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
