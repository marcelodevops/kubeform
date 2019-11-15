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

type RedisInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              RedisInstanceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            RedisInstanceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type RedisInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	AlternativeLocationID string `json:"alternativeLocationID,omitempty" tf:"alternative_location_id,omitempty" protobuf:"bytes,3,opt,name=alternativeLocationID"`
	// +optional
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty" protobuf:"bytes,4,opt,name=authorizedNetwork"`
	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty" protobuf:"bytes,5,opt,name=createTime"`
	// +optional
	CurrentLocationID string `json:"currentLocationID,omitempty" tf:"current_location_id,omitempty" protobuf:"bytes,6,opt,name=currentLocationID"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty" protobuf:"bytes,7,opt,name=displayName"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty" protobuf:"bytes,8,opt,name=host"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty" protobuf:"bytes,9,rep,name=labels"`
	// +optional
	LocationID   string `json:"locationID,omitempty" tf:"location_id,omitempty" protobuf:"bytes,10,opt,name=locationID"`
	MemorySizeGb int64  `json:"memorySizeGb" tf:"memory_size_gb" protobuf:"varint,11,opt,name=memorySizeGb"`
	Name         string `json:"name" tf:"name" protobuf:"bytes,12,opt,name=name"`
	// +optional
	Port int64 `json:"port,omitempty" tf:"port,omitempty" protobuf:"varint,13,opt,name=port"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,14,opt,name=project"`
	// +optional
	RedisConfigs map[string]string `json:"redisConfigs,omitempty" tf:"redis_configs,omitempty" protobuf:"bytes,15,rep,name=redisConfigs"`
	// +optional
	RedisVersion string `json:"redisVersion,omitempty" tf:"redis_version,omitempty" protobuf:"bytes,16,opt,name=redisVersion"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty" protobuf:"bytes,17,opt,name=region"`
	// +optional
	ReservedIPRange string `json:"reservedIPRange,omitempty" tf:"reserved_ip_range,omitempty" protobuf:"bytes,18,opt,name=reservedIPRange"`
	// +optional
	Tier string `json:"tier,omitempty" tf:"tier,omitempty" protobuf:"bytes,19,opt,name=tier"`
}

type RedisInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *RedisInstanceSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisInstanceList is a list of RedisInstances
type RedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of RedisInstance CRD objects
	Items []RedisInstance `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
