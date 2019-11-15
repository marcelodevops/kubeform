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

type EventgridEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              EventgridEventSubscriptionSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            EventgridEventSubscriptionStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type EventgridEventSubscriptionSpecEventhubEndpoint struct {
	EventhubID string `json:"eventhubID" tf:"eventhub_id" protobuf:"bytes,1,opt,name=eventhubID"`
}

type EventgridEventSubscriptionSpecHybridConnectionEndpoint struct {
	HybridConnectionID string `json:"hybridConnectionID" tf:"hybrid_connection_id" protobuf:"bytes,1,opt,name=hybridConnectionID"`
}

type EventgridEventSubscriptionSpecRetryPolicy struct {
	EventTimeToLive     int64 `json:"eventTimeToLive" tf:"event_time_to_live" protobuf:"varint,1,opt,name=eventTimeToLive"`
	MaxDeliveryAttempts int64 `json:"maxDeliveryAttempts" tf:"max_delivery_attempts" protobuf:"varint,2,opt,name=maxDeliveryAttempts"`
}

type EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination struct {
	StorageAccountID         string `json:"storageAccountID" tf:"storage_account_id" protobuf:"bytes,1,opt,name=storageAccountID"`
	StorageBlobContainerName string `json:"storageBlobContainerName" tf:"storage_blob_container_name" protobuf:"bytes,2,opt,name=storageBlobContainerName"`
}

type EventgridEventSubscriptionSpecStorageQueueEndpoint struct {
	QueueName        string `json:"queueName" tf:"queue_name" protobuf:"bytes,1,opt,name=queueName"`
	StorageAccountID string `json:"storageAccountID" tf:"storage_account_id" protobuf:"bytes,2,opt,name=storageAccountID"`
}

type EventgridEventSubscriptionSpecSubjectFilter struct {
	// +optional
	CaseSensitive bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty" protobuf:"varint,1,opt,name=caseSensitive"`
	// +optional
	SubjectBeginsWith string `json:"subjectBeginsWith,omitempty" tf:"subject_begins_with,omitempty" protobuf:"bytes,2,opt,name=subjectBeginsWith"`
	// +optional
	SubjectEndsWith string `json:"subjectEndsWith,omitempty" tf:"subject_ends_with,omitempty" protobuf:"bytes,3,opt,name=subjectEndsWith"`
}

type EventgridEventSubscriptionSpecWebhookEndpoint struct {
	Url string `json:"url" tf:"url" protobuf:"bytes,1,opt,name=url"`
}

type EventgridEventSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	EventDeliverySchema string `json:"eventDeliverySchema,omitempty" tf:"event_delivery_schema,omitempty" protobuf:"bytes,3,opt,name=eventDeliverySchema"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EventhubEndpoint []EventgridEventSubscriptionSpecEventhubEndpoint `json:"eventhubEndpoint,omitempty" tf:"eventhub_endpoint,omitempty" protobuf:"bytes,4,rep,name=eventhubEndpoint"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HybridConnectionEndpoint []EventgridEventSubscriptionSpecHybridConnectionEndpoint `json:"hybridConnectionEndpoint,omitempty" tf:"hybrid_connection_endpoint,omitempty" protobuf:"bytes,5,rep,name=hybridConnectionEndpoint"`
	// +optional
	IncludedEventTypes []string `json:"includedEventTypes,omitempty" tf:"included_event_types,omitempty" protobuf:"bytes,6,rep,name=includedEventTypes"`
	// +optional
	Labels []string `json:"labels,omitempty" tf:"labels,omitempty" protobuf:"bytes,7,rep,name=labels"`
	Name   string   `json:"name" tf:"name" protobuf:"bytes,8,opt,name=name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetryPolicy []EventgridEventSubscriptionSpecRetryPolicy `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty" protobuf:"bytes,9,rep,name=retryPolicy"`
	Scope       string                                      `json:"scope" tf:"scope" protobuf:"bytes,10,opt,name=scope"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageBlobDeadLetterDestination []EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination `json:"storageBlobDeadLetterDestination,omitempty" tf:"storage_blob_dead_letter_destination,omitempty" protobuf:"bytes,11,rep,name=storageBlobDeadLetterDestination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageQueueEndpoint []EventgridEventSubscriptionSpecStorageQueueEndpoint `json:"storageQueueEndpoint,omitempty" tf:"storage_queue_endpoint,omitempty" protobuf:"bytes,12,rep,name=storageQueueEndpoint"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SubjectFilter []EventgridEventSubscriptionSpecSubjectFilter `json:"subjectFilter,omitempty" tf:"subject_filter,omitempty" protobuf:"bytes,13,rep,name=subjectFilter"`
	// +optional
	TopicName string `json:"topicName,omitempty" tf:"topic_name,omitempty" protobuf:"bytes,14,opt,name=topicName"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WebhookEndpoint []EventgridEventSubscriptionSpecWebhookEndpoint `json:"webhookEndpoint,omitempty" tf:"webhook_endpoint,omitempty" protobuf:"bytes,15,rep,name=webhookEndpoint"`
}

type EventgridEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *EventgridEventSubscriptionSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventgridEventSubscriptionList is a list of EventgridEventSubscriptions
type EventgridEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of EventgridEventSubscription CRD objects
	Items []EventgridEventSubscription `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
