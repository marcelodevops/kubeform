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

type ApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ApiManagementSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ApiManagementStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ApiManagementSpecAdditionalLocation struct {
	// +optional
	GatewayRegionalURL string `json:"gatewayRegionalURL,omitempty" tf:"gateway_regional_url,omitempty" protobuf:"bytes,1,opt,name=gatewayRegionalURL"`
	Location           string `json:"location" tf:"location" protobuf:"bytes,2,opt,name=location"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses,omitempty" protobuf:"bytes,3,rep,name=publicIPAddresses"`
}

type ApiManagementSpecCertificate struct {
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password"`
	EncodedCertificate  string `json:"-" sensitive:"true" tf:"encoded_certificate"`
	StoreName           string `json:"storeName" tf:"store_name" protobuf:"bytes,1,opt,name=storeName"`
}

type ApiManagementSpecHostnameConfigurationManagement struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name" protobuf:"bytes,1,opt,name=hostName"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty" protobuf:"bytes,2,opt,name=keyVaultID"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty" protobuf:"varint,3,opt,name=negotiateClientCertificate"`
}

type ApiManagementSpecHostnameConfigurationPortal struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name" protobuf:"bytes,1,opt,name=hostName"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty" protobuf:"bytes,2,opt,name=keyVaultID"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty" protobuf:"varint,3,opt,name=negotiateClientCertificate"`
}

type ApiManagementSpecHostnameConfigurationProxy struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	// +optional
	DefaultSSLBinding bool   `json:"defaultSSLBinding,omitempty" tf:"default_ssl_binding,omitempty" protobuf:"varint,1,opt,name=defaultSSLBinding"`
	HostName          string `json:"hostName" tf:"host_name" protobuf:"bytes,2,opt,name=hostName"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty" protobuf:"bytes,3,opt,name=keyVaultID"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty" protobuf:"varint,4,opt,name=negotiateClientCertificate"`
}

type ApiManagementSpecHostnameConfigurationScm struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name" protobuf:"bytes,1,opt,name=hostName"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty" protobuf:"bytes,2,opt,name=keyVaultID"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty" protobuf:"varint,3,opt,name=negotiateClientCertificate"`
}

type ApiManagementSpecHostnameConfiguration struct {
	// +optional
	Management []ApiManagementSpecHostnameConfigurationManagement `json:"management,omitempty" tf:"management,omitempty" protobuf:"bytes,1,rep,name=management"`
	// +optional
	Portal []ApiManagementSpecHostnameConfigurationPortal `json:"portal,omitempty" tf:"portal,omitempty" protobuf:"bytes,2,rep,name=portal"`
	// +optional
	Proxy []ApiManagementSpecHostnameConfigurationProxy `json:"proxy,omitempty" tf:"proxy,omitempty" protobuf:"bytes,3,rep,name=proxy"`
	// +optional
	Scm []ApiManagementSpecHostnameConfigurationScm `json:"scm,omitempty" tf:"scm,omitempty" protobuf:"bytes,4,rep,name=scm"`
}

type ApiManagementSpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty" protobuf:"bytes,1,opt,name=principalID"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty" protobuf:"bytes,2,opt,name=tenantID"`
	Type     string `json:"type" tf:"type" protobuf:"bytes,3,opt,name=type"`
}

type ApiManagementSpecPolicy struct {
	// +optional
	XmlContent string `json:"xmlContent,omitempty" tf:"xml_content,omitempty" protobuf:"bytes,1,opt,name=xmlContent"`
	// +optional
	XmlLink string `json:"xmlLink,omitempty" tf:"xml_link,omitempty" protobuf:"bytes,2,opt,name=xmlLink"`
}

type ApiManagementSpecSecurity struct {
	// +optional
	DisableBackendSSL30 bool `json:"disableBackendSSL30,omitempty" tf:"disable_backend_ssl30,omitempty" protobuf:"varint,1,opt,name=disableBackendSSL30"`
	// +optional
	DisableBackendTLS10 bool `json:"disableBackendTLS10,omitempty" tf:"disable_backend_tls10,omitempty" protobuf:"varint,2,opt,name=disableBackendTLS10"`
	// +optional
	DisableBackendTLS11 bool `json:"disableBackendTLS11,omitempty" tf:"disable_backend_tls11,omitempty" protobuf:"varint,3,opt,name=disableBackendTLS11"`
	// +optional
	DisableFrontendSSL30 bool `json:"disableFrontendSSL30,omitempty" tf:"disable_frontend_ssl30,omitempty" protobuf:"varint,4,opt,name=disableFrontendSSL30"`
	// +optional
	DisableFrontendTLS10 bool `json:"disableFrontendTLS10,omitempty" tf:"disable_frontend_tls10,omitempty" protobuf:"varint,5,opt,name=disableFrontendTLS10"`
	// +optional
	DisableFrontendTLS11 bool `json:"disableFrontendTLS11,omitempty" tf:"disable_frontend_tls11,omitempty" protobuf:"varint,6,opt,name=disableFrontendTLS11"`
	// +optional
	// Deprecated
	DisableTripleDESChipers bool `json:"disableTripleDESChipers,omitempty" tf:"disable_triple_des_chipers,omitempty" protobuf:"varint,7,opt,name=disableTripleDESChipers"`
	// +optional
	DisableTripleDESCiphers bool `json:"disableTripleDESCiphers,omitempty" tf:"disable_triple_des_ciphers,omitempty" protobuf:"varint,8,opt,name=disableTripleDESCiphers"`
}

type ApiManagementSpecSignIn struct {
	Enabled bool `json:"enabled" tf:"enabled" protobuf:"varint,1,opt,name=enabled"`
}

type ApiManagementSpecSignUpTermsOfService struct {
	ConsentRequired bool `json:"consentRequired" tf:"consent_required" protobuf:"varint,1,opt,name=consentRequired"`
	Enabled         bool `json:"enabled" tf:"enabled" protobuf:"varint,2,opt,name=enabled"`
	// +optional
	Text string `json:"text,omitempty" tf:"text,omitempty" protobuf:"bytes,3,opt,name=text"`
}

type ApiManagementSpecSignUp struct {
	Enabled bool `json:"enabled" tf:"enabled" protobuf:"varint,1,opt,name=enabled"`
	// +kubebuilder:validation:MaxItems=1
	TermsOfService []ApiManagementSpecSignUpTermsOfService `json:"termsOfService" tf:"terms_of_service" protobuf:"bytes,2,rep,name=termsOfService"`
}

type ApiManagementSpecSku struct {
	Capacity int64  `json:"capacity" tf:"capacity" protobuf:"varint,1,opt,name=capacity"`
	Name     string `json:"name" tf:"name" protobuf:"bytes,2,opt,name=name"`
}

type ApiManagementSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdditionalLocation []ApiManagementSpecAdditionalLocation `json:"additionalLocation,omitempty" tf:"additional_location,omitempty" protobuf:"bytes,4,rep,name=additionalLocation"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Certificate []ApiManagementSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty" protobuf:"bytes,5,rep,name=certificate"`
	// +optional
	GatewayRegionalURL string `json:"gatewayRegionalURL,omitempty" tf:"gateway_regional_url,omitempty" protobuf:"bytes,6,opt,name=gatewayRegionalURL"`
	// +optional
	GatewayURL string `json:"gatewayURL,omitempty" tf:"gateway_url,omitempty" protobuf:"bytes,7,opt,name=gatewayURL"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HostnameConfiguration []ApiManagementSpecHostnameConfiguration `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration,omitempty" protobuf:"bytes,8,rep,name=hostnameConfiguration"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []ApiManagementSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty" protobuf:"bytes,9,rep,name=identity"`
	Location string                      `json:"location" tf:"location" protobuf:"bytes,10,opt,name=location"`
	// +optional
	ManagementAPIURL string `json:"managementAPIURL,omitempty" tf:"management_api_url,omitempty" protobuf:"bytes,11,opt,name=managementAPIURL"`
	Name             string `json:"name" tf:"name" protobuf:"bytes,12,opt,name=name"`
	// +optional
	NotificationSenderEmail string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email,omitempty" protobuf:"bytes,13,opt,name=notificationSenderEmail"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Policy []ApiManagementSpecPolicy `json:"policy,omitempty" tf:"policy,omitempty" protobuf:"bytes,14,rep,name=policy"`
	// +optional
	PortalURL string `json:"portalURL,omitempty" tf:"portal_url,omitempty" protobuf:"bytes,15,opt,name=portalURL"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses,omitempty" protobuf:"bytes,16,rep,name=publicIPAddresses"`
	PublisherEmail    string   `json:"publisherEmail" tf:"publisher_email" protobuf:"bytes,17,opt,name=publisherEmail"`
	PublisherName     string   `json:"publisherName" tf:"publisher_name" protobuf:"bytes,18,opt,name=publisherName"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name" protobuf:"bytes,19,opt,name=resourceGroupName"`
	// +optional
	ScmURL string `json:"scmURL,omitempty" tf:"scm_url,omitempty" protobuf:"bytes,20,opt,name=scmURL"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Security []ApiManagementSpecSecurity `json:"security,omitempty" tf:"security,omitempty" protobuf:"bytes,21,rep,name=security"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SignIn []ApiManagementSpecSignIn `json:"signIn,omitempty" tf:"sign_in,omitempty" protobuf:"bytes,22,rep,name=signIn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SignUp []ApiManagementSpecSignUp `json:"signUp,omitempty" tf:"sign_up,omitempty" protobuf:"bytes,23,rep,name=signUp"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ApiManagementSpecSku `json:"sku" tf:"sku" protobuf:"bytes,24,rep,name=sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty" protobuf:"bytes,25,rep,name=tags"`
}

type ApiManagementStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *ApiManagementSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementList is a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of ApiManagement CRD objects
	Items []ApiManagement `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
