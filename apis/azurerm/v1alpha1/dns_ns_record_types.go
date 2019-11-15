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

type DnsNsRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DnsNsRecordSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DnsNsRecordStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DnsNsRecordSpecRecord struct {
	Nsdname string `json:"nsdname" tf:"nsdname" protobuf:"bytes,1,opt,name=nsdname"`
}

type DnsNsRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	Name string `json:"name" tf:"name" protobuf:"bytes,3,opt,name=name"`
	// +optional
	// Deprecated
	Record []DnsNsRecordSpecRecord `json:"record,omitempty" tf:"record,omitempty" protobuf:"bytes,4,rep,name=record"`
	// +optional
	Records           []string `json:"records,omitempty" tf:"records,omitempty" protobuf:"bytes,5,rep,name=records"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,6,opt,name=resourceGroupName"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,7,rep,name=tags"`
	Ttl      int64             `json:"ttl" tf:"ttl" protobuf:"varint,8,opt,name=ttl"`
	ZoneName string            `json:"zoneName" tf:"zone_name" protobuf:"bytes,9,opt,name=zoneName"`
}

type DnsNsRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *DnsNsRecordSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsNsRecordList is a list of DnsNsRecords
type DnsNsRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of DnsNsRecord CRD objects
	Items []DnsNsRecord `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
