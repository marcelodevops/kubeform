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

type SignalrService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              SignalrServiceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            SignalrServiceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type SignalrServiceSpecSku struct {
	Capacity int64  `json:"capacity" tf:"capacity" protobuf:"varint,1,opt,name=capacity"`
	Name     string `json:"name" tf:"name" protobuf:"bytes,2,opt,name=name"`
}

type SignalrServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty" protobuf:"bytes,4,opt,name=hostname"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty" protobuf:"bytes,5,opt,name=ipAddress"`
	Location  string `json:"location" tf:"location" protobuf:"bytes,6,opt,name=location"`
	Name      string `json:"name" tf:"name" protobuf:"bytes,7,opt,name=name"`
	// +optional
	PrimaryAccessKey string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PublicPort        int64  `json:"publicPort,omitempty" tf:"public_port,omitempty" protobuf:"varint,8,opt,name=publicPort"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,9,opt,name=resourceGroupName"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	ServerPort int64 `json:"serverPort,omitempty" tf:"server_port,omitempty" protobuf:"varint,10,opt,name=serverPort"`
	// +kubebuilder:validation:MaxItems=1
	Sku []SignalrServiceSpecSku `json:"sku" tf:"sku" protobuf:"bytes,11,rep,name=sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,12,rep,name=tags"`
}

type SignalrServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *SignalrServiceSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SignalrServiceList is a list of SignalrServices
type SignalrServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of SignalrService CRD objects
	Items []SignalrService `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
