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

type DataFactoryDatasetPostgresql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DataFactoryDatasetPostgresqlSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DataFactoryDatasetPostgresqlStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DataFactoryDatasetPostgresqlSpecSchemaColumn struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,1,opt,name=description"`
	Name        string `json:"name" tf:"name" protobuf:"bytes,2,opt,name=name"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty" protobuf:"bytes,3,opt,name=type"`
}

type DataFactoryDatasetPostgresqlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty" protobuf:"bytes,3,rep,name=additionalProperties"`
	// +optional
	Annotations     []string `json:"annotations,omitempty" tf:"annotations,omitempty" protobuf:"bytes,4,rep,name=annotations"`
	DataFactoryName string   `json:"dataFactoryName" tf:"data_factory_name" protobuf:"bytes,5,opt,name=dataFactoryName"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty" protobuf:"bytes,6,opt,name=description"`
	// +optional
	Folder            string `json:"folder,omitempty" tf:"folder,omitempty" protobuf:"bytes,7,opt,name=folder"`
	LinkedServiceName string `json:"linkedServiceName" tf:"linked_service_name" protobuf:"bytes,8,opt,name=linkedServiceName"`
	Name              string `json:"name" tf:"name" protobuf:"bytes,9,opt,name=name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty" protobuf:"bytes,10,rep,name=parameters"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,11,opt,name=resourceGroupName"`
	// +optional
	SchemaColumn []DataFactoryDatasetPostgresqlSpecSchemaColumn `json:"schemaColumn,omitempty" tf:"schema_column,omitempty" protobuf:"bytes,12,rep,name=schemaColumn"`
	// +optional
	TableName string `json:"tableName,omitempty" tf:"table_name,omitempty" protobuf:"bytes,13,opt,name=tableName"`
}

type DataFactoryDatasetPostgresqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *DataFactoryDatasetPostgresqlSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryDatasetPostgresqlList is a list of DataFactoryDatasetPostgresqls
type DataFactoryDatasetPostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of DataFactoryDatasetPostgresql CRD objects
	Items []DataFactoryDatasetPostgresql `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
