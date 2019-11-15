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

type ComputeImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ComputeImageSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ComputeImageStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ComputeImageSpecRawDisk struct {
	// +optional
	ContainerType string `json:"containerType,omitempty" tf:"container_type,omitempty" protobuf:"bytes,1,opt,name=containerType"`
	// +optional
	Sha1   string `json:"sha1,omitempty" tf:"sha1,omitempty" protobuf:"bytes,2,opt,name=sha1"`
	Source string `json:"source" tf:"source" protobuf:"bytes,3,opt,name=source"`
}

type ComputeImageSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	// Deprecated
	CreateTimeout int64 `json:"createTimeout,omitempty" tf:"create_timeout,omitempty" protobuf:"varint,3,opt,name=createTimeout"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,4,opt,name=description"`
	// +optional
	Family string `json:"family,omitempty" tf:"family,omitempty" protobuf:"bytes,5,opt,name=family"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty" protobuf:"bytes,6,opt,name=labelFingerprint"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty" protobuf:"bytes,7,rep,name=labels"`
	// +optional
	Licenses []string `json:"licenses,omitempty" tf:"licenses,omitempty" protobuf:"bytes,8,rep,name=licenses"`
	Name     string   `json:"name" tf:"name" protobuf:"bytes,9,opt,name=name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,10,opt,name=project"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RawDisk []ComputeImageSpecRawDisk `json:"rawDisk,omitempty" tf:"raw_disk,omitempty" protobuf:"bytes,11,rep,name=rawDisk"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty" protobuf:"bytes,12,opt,name=selfLink"`
	// +optional
	SourceDisk string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty" protobuf:"bytes,13,opt,name=sourceDisk"`
}

type ComputeImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ComputeImageSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeImageList is a list of ComputeImages
type ComputeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ComputeImage CRD objects
	Items []ComputeImage `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
