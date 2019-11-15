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

type DataFactory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DataFactorySpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DataFactoryStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DataFactorySpecGithubConfiguration struct {
	AccountName    string `json:"accountName" tf:"account_name" protobuf:"bytes,1,opt,name=accountName"`
	BranchName     string `json:"branchName" tf:"branch_name" protobuf:"bytes,2,opt,name=branchName"`
	GitURL         string `json:"gitURL" tf:"git_url" protobuf:"bytes,3,opt,name=gitURL"`
	RepositoryName string `json:"repositoryName" tf:"repository_name" protobuf:"bytes,4,opt,name=repositoryName"`
	RootFolder     string `json:"rootFolder" tf:"root_folder" protobuf:"bytes,5,opt,name=rootFolder"`
}

type DataFactorySpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty" protobuf:"bytes,1,opt,name=principalID"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty" protobuf:"bytes,2,opt,name=tenantID"`
	Type     string `json:"type" tf:"type" protobuf:"bytes,3,opt,name=type"`
}

type DataFactorySpecVstsConfiguration struct {
	AccountName    string `json:"accountName" tf:"account_name" protobuf:"bytes,1,opt,name=accountName"`
	BranchName     string `json:"branchName" tf:"branch_name" protobuf:"bytes,2,opt,name=branchName"`
	ProjectName    string `json:"projectName" tf:"project_name" protobuf:"bytes,3,opt,name=projectName"`
	RepositoryName string `json:"repositoryName" tf:"repository_name" protobuf:"bytes,4,opt,name=repositoryName"`
	RootFolder     string `json:"rootFolder" tf:"root_folder" protobuf:"bytes,5,opt,name=rootFolder"`
	TenantID       string `json:"tenantID" tf:"tenant_id" protobuf:"bytes,6,opt,name=tenantID"`
}

type DataFactorySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	GithubConfiguration []DataFactorySpecGithubConfiguration `json:"githubConfiguration,omitempty" tf:"github_configuration,omitempty" protobuf:"bytes,3,rep,name=githubConfiguration"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity          []DataFactorySpecIdentity `json:"identity,omitempty" tf:"identity,omitempty" protobuf:"bytes,4,rep,name=identity"`
	Location          string                    `json:"location" tf:"location" protobuf:"bytes,5,opt,name=location"`
	Name              string                    `json:"name" tf:"name" protobuf:"bytes,6,opt,name=name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,7,opt,name=resourceGroupName"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,8,rep,name=tags"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VstsConfiguration []DataFactorySpecVstsConfiguration `json:"vstsConfiguration,omitempty" tf:"vsts_configuration,omitempty" protobuf:"bytes,9,rep,name=vstsConfiguration"`
}

type DataFactoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *DataFactorySpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryList is a list of DataFactorys
type DataFactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of DataFactory CRD objects
	Items []DataFactory `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
