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

type SqlDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              SqlDatabaseInstanceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            SqlDatabaseInstanceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type SqlDatabaseInstanceSpecIpAddress struct {
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty" protobuf:"bytes,1,opt,name=ipAddress"`
	// +optional
	TimeToRetire string `json:"timeToRetire,omitempty" tf:"time_to_retire,omitempty" protobuf:"bytes,2,opt,name=timeToRetire"`
}

type SqlDatabaseInstanceSpecReplicaConfiguration struct {
	// +optional
	CaCertificate string `json:"caCertificate,omitempty" tf:"ca_certificate,omitempty" protobuf:"bytes,1,opt,name=caCertificate"`
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty" protobuf:"bytes,2,opt,name=clientCertificate"`
	// +optional
	ClientKey string `json:"clientKey,omitempty" tf:"client_key,omitempty" protobuf:"bytes,3,opt,name=clientKey"`
	// +optional
	ConnectRetryInterval int64 `json:"connectRetryInterval,omitempty" tf:"connect_retry_interval,omitempty" protobuf:"varint,4,opt,name=connectRetryInterval"`
	// +optional
	DumpFilePath string `json:"dumpFilePath,omitempty" tf:"dump_file_path,omitempty" protobuf:"bytes,5,opt,name=dumpFilePath"`
	// +optional
	FailoverTarget bool `json:"failoverTarget,omitempty" tf:"failover_target,omitempty" protobuf:"varint,6,opt,name=failoverTarget"`
	// +optional
	MasterHeartbeatPeriod int64 `json:"masterHeartbeatPeriod,omitempty" tf:"master_heartbeat_period,omitempty" protobuf:"varint,7,opt,name=masterHeartbeatPeriod"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	SslCipher string `json:"sslCipher,omitempty" tf:"ssl_cipher,omitempty" protobuf:"bytes,8,opt,name=sslCipher"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty" protobuf:"bytes,9,opt,name=username"`
	// +optional
	VerifyServerCertificate bool `json:"verifyServerCertificate,omitempty" tf:"verify_server_certificate,omitempty" protobuf:"varint,10,opt,name=verifyServerCertificate"`
}

type SqlDatabaseInstanceSpecServerCaCert struct {
	// +optional
	Cert string `json:"cert,omitempty" tf:"cert,omitempty" protobuf:"bytes,1,opt,name=cert"`
	// +optional
	CommonName string `json:"commonName,omitempty" tf:"common_name,omitempty" protobuf:"bytes,2,opt,name=commonName"`
	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty" protobuf:"bytes,3,opt,name=createTime"`
	// +optional
	ExpirationTime string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty" protobuf:"bytes,4,opt,name=expirationTime"`
	// +optional
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty" protobuf:"bytes,5,opt,name=sha1Fingerprint"`
}

type SqlDatabaseInstanceSpecSettingsBackupConfiguration struct {
	// +optional
	BinaryLogEnabled bool `json:"binaryLogEnabled,omitempty" tf:"binary_log_enabled,omitempty" protobuf:"varint,1,opt,name=binaryLogEnabled"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty" protobuf:"varint,2,opt,name=enabled"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty" protobuf:"bytes,3,opt,name=startTime"`
}

type SqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

type SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks struct {
	// +optional
	ExpirationTime string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty" protobuf:"bytes,1,opt,name=expirationTime"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty" protobuf:"bytes,3,opt,name=value"`
}

type SqlDatabaseInstanceSpecSettingsIpConfiguration struct {
	// +optional
	AuthorizedNetworks []SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks `json:"authorizedNetworks,omitempty" tf:"authorized_networks,omitempty" protobuf:"bytes,1,rep,name=authorizedNetworks"`
	// +optional
	Ipv4Enabled bool `json:"ipv4Enabled,omitempty" tf:"ipv4_enabled,omitempty" protobuf:"varint,2,opt,name=ipv4Enabled"`
	// +optional
	PrivateNetwork string `json:"privateNetwork,omitempty" tf:"private_network,omitempty" protobuf:"bytes,3,opt,name=privateNetwork"`
	// +optional
	RequireSSL bool `json:"requireSSL,omitempty" tf:"require_ssl,omitempty" protobuf:"varint,4,opt,name=requireSSL"`
}

type SqlDatabaseInstanceSpecSettingsLocationPreference struct {
	// +optional
	FollowGaeApplication string `json:"followGaeApplication,omitempty" tf:"follow_gae_application,omitempty" protobuf:"bytes,1,opt,name=followGaeApplication"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty" protobuf:"bytes,2,opt,name=zone"`
}

type SqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	// +optional
	Day int64 `json:"day,omitempty" tf:"day,omitempty" protobuf:"varint,1,opt,name=day"`
	// +optional
	Hour int64 `json:"hour,omitempty" tf:"hour,omitempty" protobuf:"varint,2,opt,name=hour"`
	// +optional
	UpdateTrack string `json:"updateTrack,omitempty" tf:"update_track,omitempty" protobuf:"bytes,3,opt,name=updateTrack"`
}

type SqlDatabaseInstanceSpecSettings struct {
	// +optional
	ActivationPolicy string `json:"activationPolicy,omitempty" tf:"activation_policy,omitempty" protobuf:"bytes,1,opt,name=activationPolicy"`
	// +optional
	AuthorizedGaeApplications []string `json:"authorizedGaeApplications,omitempty" tf:"authorized_gae_applications,omitempty" protobuf:"bytes,2,rep,name=authorizedGaeApplications"`
	// +optional
	AvailabilityType string `json:"availabilityType,omitempty" tf:"availability_type,omitempty" protobuf:"bytes,3,opt,name=availabilityType"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BackupConfiguration []SqlDatabaseInstanceSpecSettingsBackupConfiguration `json:"backupConfiguration,omitempty" tf:"backup_configuration,omitempty" protobuf:"bytes,4,rep,name=backupConfiguration"`
	// +optional
	CrashSafeReplication bool `json:"crashSafeReplication,omitempty" tf:"crash_safe_replication,omitempty" protobuf:"varint,5,opt,name=crashSafeReplication"`
	// +optional
	DatabaseFlags []SqlDatabaseInstanceSpecSettingsDatabaseFlags `json:"databaseFlags,omitempty" tf:"database_flags,omitempty" protobuf:"bytes,6,rep,name=databaseFlags"`
	// +optional
	DiskAutoresize bool `json:"diskAutoresize,omitempty" tf:"disk_autoresize,omitempty" protobuf:"varint,7,opt,name=diskAutoresize"`
	// +optional
	DiskSize int64 `json:"diskSize,omitempty" tf:"disk_size,omitempty" protobuf:"varint,8,opt,name=diskSize"`
	// +optional
	DiskType string `json:"diskType,omitempty" tf:"disk_type,omitempty" protobuf:"bytes,9,opt,name=diskType"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpConfiguration []SqlDatabaseInstanceSpecSettingsIpConfiguration `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty" protobuf:"bytes,10,rep,name=ipConfiguration"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LocationPreference []SqlDatabaseInstanceSpecSettingsLocationPreference `json:"locationPreference,omitempty" tf:"location_preference,omitempty" protobuf:"bytes,11,rep,name=locationPreference"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenanceWindow []SqlDatabaseInstanceSpecSettingsMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty" protobuf:"bytes,12,rep,name=maintenanceWindow"`
	// +optional
	PricingPlan string `json:"pricingPlan,omitempty" tf:"pricing_plan,omitempty" protobuf:"bytes,13,opt,name=pricingPlan"`
	// +optional
	ReplicationType string `json:"replicationType,omitempty" tf:"replication_type,omitempty" protobuf:"bytes,14,opt,name=replicationType"`
	Tier            string `json:"tier" tf:"tier" protobuf:"bytes,15,opt,name=tier"`
	// +optional
	UserLabels map[string]string `json:"userLabels,omitempty" tf:"user_labels,omitempty" protobuf:"bytes,16,rep,name=userLabels"`
	// +optional
	Version int64 `json:"version,omitempty" tf:"version,omitempty" protobuf:"varint,17,opt,name=version"`
}

type SqlDatabaseInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-" protobuf:"bytes,1,opt,name=providerRef"`

	ID string `json:"id,omitempty" tf:"id,omitempty" protobuf:"bytes,2,opt,name=id"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-" protobuf:"bytes,3,opt,name=secretRef"`

	// +optional
	ConnectionName string `json:"connectionName,omitempty" tf:"connection_name,omitempty" protobuf:"bytes,4,opt,name=connectionName"`
	// +optional
	DatabaseVersion string `json:"databaseVersion,omitempty" tf:"database_version,omitempty" protobuf:"bytes,5,opt,name=databaseVersion"`
	// +optional
	FirstIPAddress string `json:"firstIPAddress,omitempty" tf:"first_ip_address,omitempty" protobuf:"bytes,6,opt,name=firstIPAddress"`
	// +optional
	IpAddress []SqlDatabaseInstanceSpecIpAddress `json:"ipAddress,omitempty" tf:"ip_address,omitempty" protobuf:"bytes,7,rep,name=ipAddress"`
	// +optional
	MasterInstanceName string `json:"masterInstanceName,omitempty" tf:"master_instance_name,omitempty" protobuf:"bytes,8,opt,name=masterInstanceName"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty" protobuf:"bytes,9,opt,name=name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty" protobuf:"bytes,10,opt,name=project"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty" protobuf:"bytes,11,opt,name=region"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReplicaConfiguration []SqlDatabaseInstanceSpecReplicaConfiguration `json:"replicaConfiguration,omitempty" tf:"replica_configuration,omitempty" protobuf:"bytes,12,rep,name=replicaConfiguration"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty" protobuf:"bytes,13,opt,name=selfLink"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServerCaCert []SqlDatabaseInstanceSpecServerCaCert `json:"serverCaCert,omitempty" tf:"server_ca_cert,omitempty" protobuf:"bytes,14,rep,name=serverCaCert"`
	// +optional
	ServiceAccountEmailAddress string `json:"serviceAccountEmailAddress,omitempty" tf:"service_account_email_address,omitempty" protobuf:"bytes,15,opt,name=serviceAccountEmailAddress"`
	// +kubebuilder:validation:MaxItems=1
	Settings []SqlDatabaseInstanceSpecSettings `json:"settings" tf:"settings" protobuf:"bytes,16,rep,name=settings"`
}

type SqlDatabaseInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// +optional
	Output *SqlDatabaseInstanceSpec `json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
	// +optional
	State *base.State `json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	// +optional
	Phase base.Phase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=kubeform.dev/kubeform/apis/base/v1alpha1.Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseInstanceList is a list of SqlDatabaseInstances
type SqlDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of SqlDatabaseInstance CRD objects
	Items []SqlDatabaseInstance `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
