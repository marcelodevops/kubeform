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

type DevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DevTestLinuxVirtualMachineSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DevTestLinuxVirtualMachineStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Offer     string `json:"offer" tf:"offer" protobuf:"bytes,1,opt,name=offer"`
	Publisher string `json:"publisher" tf:"publisher" protobuf:"bytes,2,opt,name=publisher"`
	Sku       string `json:"sku" tf:"sku" protobuf:"bytes,3,opt,name=sku"`
	Version   string `json:"version" tf:"version" protobuf:"bytes,4,opt,name=version"`
}

type DevTestLinuxVirtualMachineSpecInboundNATRule struct {
	BackendPort int64 `json:"backendPort" tf:"backend_port" protobuf:"varint,1,opt,name=backendPort"`
	// +optional
	FrontendPort int64  `json:"frontendPort,omitempty" tf:"frontend_port,omitempty" protobuf:"varint,2,opt,name=frontendPort"`
	Protocol     string `json:"protocol" tf:"protocol" protobuf:"bytes,3,opt,name=protocol"`
}

type DevTestLinuxVirtualMachineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	AllowClaim bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty" protobuf:"varint,4,opt,name=allowClaim"`
	// +optional
	DisallowPublicIPAddress bool `json:"disallowPublicIPAddress,omitempty" tf:"disallow_public_ip_address,omitempty" protobuf:"varint,5,opt,name=disallowPublicIPAddress"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty" protobuf:"bytes,6,opt,name=fqdn"`
	// +kubebuilder:validation:MaxItems=1
	GalleryImageReference []DevTestLinuxVirtualMachineSpecGalleryImageReference `json:"galleryImageReference" tf:"gallery_image_reference" protobuf:"bytes,7,rep,name=galleryImageReference"`
	// +optional
	InboundNATRule      []DevTestLinuxVirtualMachineSpecInboundNATRule `json:"inboundNATRule,omitempty" tf:"inbound_nat_rule,omitempty" protobuf:"bytes,8,rep,name=inboundNATRule"`
	LabName             string                                         `json:"labName" tf:"lab_name" protobuf:"bytes,9,opt,name=labName"`
	LabSubnetName       string                                         `json:"labSubnetName" tf:"lab_subnet_name" protobuf:"bytes,10,opt,name=labSubnetName"`
	LabVirtualNetworkID string                                         `json:"labVirtualNetworkID" tf:"lab_virtual_network_id" protobuf:"bytes,11,opt,name=labVirtualNetworkID"`
	Location            string                                         `json:"location" tf:"location" protobuf:"bytes,12,opt,name=location"`
	Name                string                                         `json:"name" tf:"name" protobuf:"bytes,13,opt,name=name"`
	// +optional
	Notes string `json:"notes,omitempty" tf:"notes,omitempty" protobuf:"bytes,14,opt,name=notes"`
	// +optional
	Password          string `json:"-" sensitive:"true" tf:"password,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,15,opt,name=resourceGroupName"`
	Size              string `json:"size" tf:"size"`
	// +optional
	SshKey      string `json:"sshKey,omitempty" tf:"ssh_key,omitempty" protobuf:"bytes,17,opt,name=sshKey"`
	StorageType string `json:"storageType" tf:"storage_type" protobuf:"bytes,18,opt,name=storageType"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,19,rep,name=tags"`
	// +optional
	UniqueIdentifier string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty" protobuf:"bytes,20,opt,name=uniqueIdentifier"`
	Username         string `json:"username" tf:"username" protobuf:"bytes,21,opt,name=username"`
}

type DevTestLinuxVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *DevTestLinuxVirtualMachineSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestLinuxVirtualMachineList is a list of DevTestLinuxVirtualMachines
type DevTestLinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of DevTestLinuxVirtualMachine CRD objects
	Items []DevTestLinuxVirtualMachine `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
