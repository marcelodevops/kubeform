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

type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ComputeInstanceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ComputeInstanceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ComputeInstanceSpecAttachedDisk struct {
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty" protobuf:"bytes,1,opt,name=deviceName"`
	// +optional
	DiskEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	DiskEncryptionKeySha256 string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty" protobuf:"bytes,2,opt,name=diskEncryptionKeySha256"`
	// +optional
	Mode   string `json:"mode,omitempty" tf:"mode,omitempty" protobuf:"bytes,3,opt,name=mode"`
	Source string `json:"source" tf:"source" protobuf:"bytes,4,opt,name=source"`
}

type ComputeInstanceSpecBootDiskInitializeParams struct {
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty" protobuf:"bytes,1,opt,name=image"`
	// +optional
	Size int64 `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty" protobuf:"bytes,3,opt,name=type"`
}

type ComputeInstanceSpecBootDisk struct {
	// +optional
	AutoDelete bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty" protobuf:"varint,1,opt,name=autoDelete"`
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty" protobuf:"bytes,2,opt,name=deviceName"`
	// +optional
	DiskEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	DiskEncryptionKeySha256 string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty" protobuf:"bytes,3,opt,name=diskEncryptionKeySha256"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InitializeParams []ComputeInstanceSpecBootDiskInitializeParams `json:"initializeParams,omitempty" tf:"initialize_params,omitempty" protobuf:"bytes,4,rep,name=initializeParams"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty" protobuf:"bytes,5,opt,name=source"`
}

type ComputeInstanceSpecGuestAccelerator struct {
	Count int64  `json:"count" tf:"count" protobuf:"varint,1,opt,name=count"`
	Type  string `json:"type" tf:"type" protobuf:"bytes,2,opt,name=type"`
}

type ComputeInstanceSpecNetworkInterfaceAccessConfig struct {
	// +optional
	// Deprecated
	AssignedNATIP string `json:"assignedNATIP,omitempty" tf:"assigned_nat_ip,omitempty" protobuf:"bytes,1,opt,name=assignedNATIP"`
	// +optional
	NatIP string `json:"natIP,omitempty" tf:"nat_ip,omitempty" protobuf:"bytes,2,opt,name=natIP"`
	// +optional
	NetworkTier string `json:"networkTier,omitempty" tf:"network_tier,omitempty" protobuf:"bytes,3,opt,name=networkTier"`
	// +optional
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty" protobuf:"bytes,4,opt,name=publicPtrDomainName"`
}

type ComputeInstanceSpecNetworkInterfaceAliasIPRange struct {
	IpCIDRRange string `json:"ipCIDRRange" tf:"ip_cidr_range" protobuf:"bytes,1,opt,name=ipCIDRRange"`
	// +optional
	SubnetworkRangeName string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty" protobuf:"bytes,2,opt,name=subnetworkRangeName"`
}

type ComputeInstanceSpecNetworkInterface struct {
	// +optional
	AccessConfig []ComputeInstanceSpecNetworkInterfaceAccessConfig `json:"accessConfig,omitempty" tf:"access_config,omitempty" protobuf:"bytes,1,rep,name=accessConfig"`
	// +optional
	// Deprecated
	Address string `json:"address,omitempty" tf:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
	// +optional
	AliasIPRange []ComputeInstanceSpecNetworkInterfaceAliasIPRange `json:"aliasIPRange,omitempty" tf:"alias_ip_range,omitempty" protobuf:"bytes,3,rep,name=aliasIPRange"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty" protobuf:"bytes,4,opt,name=name"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty" protobuf:"bytes,5,opt,name=network"`
	// +optional
	NetworkIP string `json:"networkIP,omitempty" tf:"network_ip,omitempty" protobuf:"bytes,6,opt,name=networkIP"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty" protobuf:"bytes,7,opt,name=subnetwork"`
	// +optional
	SubnetworkProject string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty" protobuf:"bytes,8,opt,name=subnetworkProject"`
}

type ComputeInstanceSpecScheduling struct {
	// +optional
	AutomaticRestart bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty" protobuf:"varint,1,opt,name=automaticRestart"`
	// +optional
	OnHostMaintenance string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty" protobuf:"bytes,2,opt,name=onHostMaintenance"`
	// +optional
	Preemptible bool `json:"preemptible,omitempty" tf:"preemptible,omitempty" protobuf:"varint,3,opt,name=preemptible"`
}

type ComputeInstanceSpecScratchDisk struct {
	// +optional
	Interface string `json:"interface,omitempty" tf:"interface,omitempty" protobuf:"bytes,1,opt,name=interface"`
}

type ComputeInstanceSpecServiceAccount struct {
	// +optional
	Email  string   `json:"email,omitempty" tf:"email,omitempty" protobuf:"bytes,1,opt,name=email"`
	Scopes []string `json:"scopes" tf:"scopes" protobuf:"bytes,2,rep,name=scopes"`
}

type ComputeInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	AllowStoppingForUpdate bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty" protobuf:"varint,4,opt,name=allowStoppingForUpdate"`
	// +optional
	AttachedDisk []ComputeInstanceSpecAttachedDisk `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty" protobuf:"bytes,5,rep,name=attachedDisk"`
	// +kubebuilder:validation:MaxItems=1
	BootDisk []ComputeInstanceSpecBootDisk `json:"bootDisk" tf:"boot_disk" protobuf:"bytes,6,rep,name=bootDisk"`
	// +optional
	CanIPForward bool `json:"canIPForward,omitempty" tf:"can_ip_forward,omitempty" protobuf:"varint,7,opt,name=canIPForward"`
	// +optional
	CpuPlatform string `json:"cpuPlatform,omitempty" tf:"cpu_platform,omitempty" protobuf:"bytes,8,opt,name=cpuPlatform"`
	// +optional
	// Deprecated
	CreateTimeout int64 `json:"createTimeout,omitempty" tf:"create_timeout,omitempty" protobuf:"varint,9,opt,name=createTimeout"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty" protobuf:"varint,10,opt,name=deletionProtection"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,11,opt,name=description"`
	// +optional
	GuestAccelerator []ComputeInstanceSpecGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty" protobuf:"bytes,12,rep,name=guestAccelerator"`
	// +optional
	InstanceID string `json:"instanceID,omitempty" tf:"instance_id,omitempty" protobuf:"bytes,13,opt,name=instanceID"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty" protobuf:"bytes,14,opt,name=labelFingerprint"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty" tf:"labels,omitempty" protobuf:"bytes,15,rep,name=labels"`
	MachineType string            `json:"machineType" tf:"machine_type" protobuf:"bytes,16,opt,name=machineType"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty" protobuf:"bytes,17,rep,name=metadata"`
	// +optional
	MetadataFingerprint string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty" protobuf:"bytes,18,opt,name=metadataFingerprint"`
	// +optional
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty" protobuf:"bytes,19,opt,name=metadataStartupScript"`
	// +optional
	MinCPUPlatform   string                                `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty" protobuf:"bytes,20,opt,name=minCPUPlatform"`
	Name             string                                `json:"name" tf:"name" protobuf:"bytes,21,opt,name=name"`
	NetworkInterface []ComputeInstanceSpecNetworkInterface `json:"networkInterface" tf:"network_interface" protobuf:"bytes,22,rep,name=networkInterface"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,23,opt,name=project"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scheduling []ComputeInstanceSpecScheduling `json:"scheduling,omitempty" tf:"scheduling,omitempty" protobuf:"bytes,24,rep,name=scheduling"`
	// +optional
	ScratchDisk []ComputeInstanceSpecScratchDisk `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty" protobuf:"bytes,25,rep,name=scratchDisk"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty" protobuf:"bytes,26,opt,name=selfLink"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceAccount []ComputeInstanceSpecServiceAccount `json:"serviceAccount,omitempty" tf:"service_account,omitempty" protobuf:"bytes,27,rep,name=serviceAccount"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,28,rep,name=tags"`
	// +optional
	TagsFingerprint string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty" protobuf:"bytes,29,opt,name=tagsFingerprint"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty" protobuf:"bytes,30,opt,name=zone"`
}

type ComputeInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ComputeInstanceSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceList is a list of ComputeInstances
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ComputeInstance CRD objects
	Items []ComputeInstance `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
