package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type EcsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsServiceSpec   `json:"spec,omitempty"`
	Status            EcsServiceStatus `json:"status,omitempty"`
}

type EcsServiceSpecDeploymentController struct {
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type EcsServiceSpecLoadBalancer struct {
	ContainerName string `json:"containerName" tf:"container_name"`
	ContainerPort int    `json:"containerPort" tf:"container_port"`
	// +optional
	ElbName string `json:"elbName,omitempty" tf:"elb_name,omitempty"`
	// +optional
	TargetGroupArn string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`
}

type EcsServiceSpecNetworkConfiguration struct {
	// +optional
	AssignPublicIP bool `json:"assignPublicIP,omitempty" tf:"assign_public_ip,omitempty"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	Subnets        []string `json:"subnets" tf:"subnets"`
}

type EcsServiceSpecOrderedPlacementStrategy struct {
	// +optional
	Field string `json:"field,omitempty" tf:"field,omitempty"`
	Type  string `json:"type" tf:"type"`
}

type EcsServiceSpecPlacementConstraints struct {
	// +optional
	Expression string `json:"expression,omitempty" tf:"expression,omitempty"`
	Type       string `json:"type" tf:"type"`
}

type EcsServiceSpecServiceRegistries struct {
	// +optional
	ContainerName string `json:"containerName,omitempty" tf:"container_name,omitempty"`
	// +optional
	ContainerPort int `json:"containerPort,omitempty" tf:"container_port,omitempty"`
	// +optional
	Port        int    `json:"port,omitempty" tf:"port,omitempty"`
	RegistryArn string `json:"registryArn" tf:"registry_arn"`
}

type EcsServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cluster string `json:"cluster,omitempty" tf:"cluster,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeploymentController []EcsServiceSpecDeploymentController `json:"deploymentController,omitempty" tf:"deployment_controller,omitempty"`
	// +optional
	DeploymentMaximumPercent int `json:"deploymentMaximumPercent,omitempty" tf:"deployment_maximum_percent,omitempty"`
	// +optional
	DeploymentMinimumHealthyPercent int `json:"deploymentMinimumHealthyPercent,omitempty" tf:"deployment_minimum_healthy_percent,omitempty"`
	// +optional
	DesiredCount int `json:"desiredCount,omitempty" tf:"desired_count,omitempty"`
	// +optional
	EnableEcsManagedTags bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags,omitempty"`
	// +optional
	HealthCheckGracePeriodSeconds int `json:"healthCheckGracePeriodSeconds,omitempty" tf:"health_check_grace_period_seconds,omitempty"`
	// +optional
	IamRole string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`
	// +optional
	LaunchType string `json:"launchType,omitempty" tf:"launch_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoadBalancer []EcsServiceSpecLoadBalancer `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`
	Name         string                       `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkConfiguration []EcsServiceSpecNetworkConfiguration `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	OrderedPlacementStrategy []EcsServiceSpecOrderedPlacementStrategy `json:"orderedPlacementStrategy,omitempty" tf:"ordered_placement_strategy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PlacementConstraints []EcsServiceSpecPlacementConstraints `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`
	// +optional
	PlatformVersion string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`
	// +optional
	PropagateTags string `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`
	// +optional
	SchedulingStrategy string `json:"schedulingStrategy,omitempty" tf:"scheduling_strategy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceRegistries []EcsServiceSpecServiceRegistries `json:"serviceRegistries,omitempty" tf:"service_registries,omitempty"`
	// +optional
	Tags           map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TaskDefinition string            `json:"taskDefinition" tf:"task_definition"`
}

type EcsServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EcsServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcsServiceList is a list of EcsServices
type EcsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcsService CRD objects
	Items []EcsService `json:"items,omitempty"`
}
