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

type ComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ComputeSubnetworkSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ComputeSubnetworkStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ComputeSubnetworkSpecSecondaryIPRange struct {
	IpCIDRRange string `json:"ipCIDRRange" tf:"ip_cidr_range" protobuf:"bytes,1,opt,name=ipCIDRRange"`
	RangeName   string `json:"rangeName" tf:"range_name" protobuf:"bytes,2,opt,name=rangeName"`
}

type ComputeSubnetworkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty" protobuf:"bytes,3,opt,name=creationTimestamp"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,4,opt,name=description"`
	// +optional
	EnableFlowLogs bool `json:"enableFlowLogs,omitempty" tf:"enable_flow_logs,omitempty" protobuf:"varint,5,opt,name=enableFlowLogs"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty" protobuf:"bytes,6,opt,name=fingerprint"`
	// +optional
	GatewayAddress string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty" protobuf:"bytes,7,opt,name=gatewayAddress"`
	IpCIDRRange    string `json:"ipCIDRRange" tf:"ip_cidr_range" protobuf:"bytes,8,opt,name=ipCIDRRange"`
	Name           string `json:"name" tf:"name" protobuf:"bytes,9,opt,name=name"`
	Network        string `json:"network" tf:"network" protobuf:"bytes,10,opt,name=network"`
	// +optional
	PrivateIPGoogleAccess bool `json:"privateIPGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty" protobuf:"varint,11,opt,name=privateIPGoogleAccess"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,12,opt,name=project"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty" protobuf:"bytes,13,opt,name=region"`
	// +optional
	SecondaryIPRange []ComputeSubnetworkSpecSecondaryIPRange `json:"secondaryIPRange,omitempty" tf:"secondary_ip_range,omitempty" protobuf:"bytes,14,rep,name=secondaryIPRange"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty" protobuf:"bytes,15,opt,name=selfLink"`
}

type ComputeSubnetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ComputeSubnetworkSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkList is a list of ComputeSubnetworks
type ComputeSubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ComputeSubnetwork CRD objects
	Items []ComputeSubnetwork `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
