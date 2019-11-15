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

type ComputeRegionDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ComputeRegionDiskSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ComputeRegionDiskStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ComputeRegionDiskSpecDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty" protobuf:"bytes,1,opt,name=rawKey"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty" protobuf:"bytes,2,opt,name=sha256"`
}

type ComputeRegionDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty" protobuf:"bytes,1,opt,name=rawKey"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty" protobuf:"bytes,2,opt,name=sha256"`
}

type ComputeRegionDiskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty" protobuf:"bytes,3,opt,name=creationTimestamp"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,4,opt,name=description"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ComputeRegionDiskSpecDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty" protobuf:"bytes,5,rep,name=diskEncryptionKey"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty" protobuf:"bytes,6,opt,name=labelFingerprint"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty" protobuf:"bytes,7,rep,name=labels"`
	// +optional
	LastAttachTimestamp string `json:"lastAttachTimestamp,omitempty" tf:"last_attach_timestamp,omitempty" protobuf:"bytes,8,opt,name=lastAttachTimestamp"`
	// +optional
	LastDetachTimestamp string `json:"lastDetachTimestamp,omitempty" tf:"last_detach_timestamp,omitempty" protobuf:"bytes,9,opt,name=lastDetachTimestamp"`
	Name                string `json:"name" tf:"name" protobuf:"bytes,10,opt,name=name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,11,opt,name=project"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty" protobuf:"bytes,12,opt,name=region"`
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=2
	ReplicaZones []string `json:"replicaZones" tf:"replica_zones" protobuf:"bytes,13,rep,name=replicaZones"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty" protobuf:"bytes,14,opt,name=selfLink"`
	// +optional
	Size int64 `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Snapshot string `json:"snapshot,omitempty" tf:"snapshot,omitempty" protobuf:"bytes,16,opt,name=snapshot"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey []ComputeRegionDiskSpecSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty" protobuf:"bytes,17,rep,name=sourceSnapshotEncryptionKey"`
	// +optional
	SourceSnapshotID string `json:"sourceSnapshotID,omitempty" tf:"source_snapshot_id,omitempty" protobuf:"bytes,18,opt,name=sourceSnapshotID"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty" protobuf:"bytes,19,opt,name=type"`
	// +optional
	Users []string `json:"users,omitempty" tf:"users,omitempty" protobuf:"bytes,20,rep,name=users"`
}

type ComputeRegionDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ComputeRegionDiskSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRegionDiskList is a list of ComputeRegionDisks
type ComputeRegionDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ComputeRegionDisk CRD objects
	Items []ComputeRegionDisk `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
