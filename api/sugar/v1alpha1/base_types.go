package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
)

type BaseSpec struct {
	// Defines a list of additional Services that are exposed by a Cluster.
	// This field allows Services of selected Components,
	// alongside Services defined with ComponentService.
	//
	// Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.
	//
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Services []appsv1alpha1.ClusterService `json:"services,omitempty"`

	// Specifies the scheduling policy for the Cluster.
	//
	// +optional
	SchedulingPolicy *SchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// Specifies the backup configuration of the Cluster.
	//
	// +optional
	Backup *appsv1alpha1.ClusterBackup `json:"backup,omitempty"`

	// Specifies the behavior when a Cluster is deleted.
	// It defines how resources, data, and backups associated with a Cluster are managed during termination.
	// Choose a policy based on the desired level of resource cleanup and data preservation:
	//
	// - `DoNotTerminate`: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.
	// - `Halt`: Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs),
	//   allowing for data preservation while stopping other operations.
	// - `Delete`: Extends the `Halt` policy by also removing PVCs, leading to a thorough cleanup while
	//   removing all persistent data.
	// - `WipeOut`: An aggressive policy that deletes all Cluster resources, including volume snapshots and
	//   backups in external storage.
	//   This results in complete data removal and should be used cautiously, primarily in non-production environments
	//   to avoid irreversible data loss.
	//
	// Warning: Choosing an inappropriate termination policy can result in data loss.
	// The `WipeOut` policy is particularly risky in production environments due to its irreversible nature.
	//
	// +kubebuilder:validation:Required
	TerminationPolicy appsv1alpha1.TerminationPolicyType `json:"terminationPolicy"`
}

type BaseStatus struct {
	appsv1alpha1.ClusterStatus `json:",inline"`
}

// ClusterComponentSpec defines the specification of a Component within a Cluster.
type ClusterComponentSpec struct {
	// Specifies the Component's name.
	// It's part of the Service DNS name and must comply with the IANA service naming rule.
	//
	// +kubebuilder:validation:MaxLength=22
	// +kubebuilder:validation:Pattern:=`^[a-z]([a-z0-9\-]*[a-z0-9])?$`
	Name string `json:"name"`

	// ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
	// The version should follow the syntax and semantics of the "Semantic Versioning" specification (http://semver.org/).
	// If no version is specified, the latest available version will be used.
	//
	// +kubebuilder:validation:MaxLength=32
	// +optional
	ServiceVersion string `json:"serviceVersion,omitempty"`

	// Defines a list of ServiceRef for a Component, enabling access to both external services and
	// Services provided by other Clusters.
	//
	// Types of services:
	//
	// - External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
	//   Require a ServiceDescriptor for connection details.
	// - Services provided by a Cluster: Managed by the same KubeBlocks operator;
	//   identified using Cluster, Component and Service names.
	//
	// ServiceRefs with identical `serviceRef.name` in the same Cluster are considered the same.
	//
	// Example:
	// ```yaml
	// serviceRefs:
	//   - name: "redis-sentinel"
	//     serviceDescriptor:
	//       name: "external-redis-sentinel"
	//   - name: "postgres-cluster"
	//     clusterServiceSelector:
	//       cluster: "my-postgres-cluster"
	//       service:
	//         component: "postgresql"
	// ```
	// The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.
	//
	// +optional
	ServiceRefs []appsv1alpha1.ServiceRef `json:"serviceRefs,omitempty"`

	// Specifies which types of logs should be collected for the Component.
	// The log types are defined in the `componentDefinition.spec.logConfigs` field with the LogConfig entries.
	//
	// The elements in the `enabledLogs` array correspond to the names of the LogConfig entries.
	// For example, if the `componentDefinition.spec.logConfigs` defines LogConfig entries with
	// names "slow_query_log" and "error_log",
	// you can enable the collection of these logs by including their names in the `enabledLogs` array:
	// ```yaml
	// enabledLogs:
	// - slow_query_log
	// - error_log
	// ```
	//
	// +listType=set
	// +optional
	EnabledLogs []string `json:"enabledLogs,omitempty"`

	// Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas"`

	// Specifies the scheduling policy for the Component.
	//
	// +optional
	SchedulingPolicy *SchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// Specifies the resources required by the Component.
	// It allows defining the CPU, memory requirements and limits for the Component's containers.
	//
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component.
	// Each template specifies the desired characteristics of a persistent volume, such as storage class,
	// size, and access modes.
	// These templates are used to dynamically provision persistent volumes for the Component.
	//
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	// +optional
	VolumeClaimTemplates []appsv1alpha1.ClusterComponentVolumeClaimTemplate `json:"volumeClaimTemplates,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name"`

	// Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.
	//
	// +optional
	Services []appsv1alpha1.ClusterComponentService `json:"services,omitempty"`

	// A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)
	// for secure communication.
	// When set to true, the Component will be configured to use TLS encryption for its network connections.
	// This ensures that the data transmitted between the Component and its clients or other Components is encrypted
	// and protected from unauthorized access.
	// If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys,
	// to properly set up the secure communication channel.
	//
	// +optional
	TLS bool `json:"tls,omitempty"`

	// Specifies the configuration for the TLS certificates issuer.
	// It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
	// The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.
	// Required when TLS is enabled.
	//
	// +optional
	Issuer *appsv1alpha1.Issuer `json:"issuer,omitempty"`

	// Specifies the name of the ServiceAccount required by the running Component.
	// This ServiceAccount is used to grant necessary permissions for the Component's Pods to interact
	// with other Kubernetes resources, such as modifying Pod labels or sending events.
	//
	// Defaults:
	// If not specified, KubeBlocks automatically assigns a default ServiceAccount named "kb-{cluster.name}",
	// bound to a default role installed together with KubeBlocks.
	//
	// Future Changes:
	// Future versions might change the default ServiceAccount creation strategy to one per Component,
	// potentially revising the naming to "kb-{cluster.name}-{component.name}".
	//
	// Users can override the automatic ServiceAccount assignment by explicitly setting the name of
	// an existed ServiceAccount in this field.
	//
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes
	// in the Cluster's Pods.
	// This is useful in scenarios where users need to provide additional resources to the Cluster, such as:
	//
	// - Mounting custom scripts or configuration files during Cluster startup.
	// - Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster.
	//
	// +optional
	UserResourceRefs *appsv1alpha1.UserResourceRefs `json:"userResourceRefs,omitempty"`

	// Allows for the customization of configuration values for each instance within a Component.
	// An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
	// While instances typically share a common configuration as defined in the ClusterComponentSpec,
	// they can require unique settings in various scenarios:
	//
	// For example:
	// - A database Component might require different resource allocations for primary and secondary instances,
	//   with primaries needing more resources.
	// - During a rolling upgrade, a Component may first update the image for one or a few instances,
	//   and then update the remaining instances after verifying that the updated instances are functioning correctly.
	//
	// InstanceTemplate allows for specifying these unique configurations per instance.
	// Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
	// starting with an ordinal of 0.
	// It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.
	//
	// The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component.
	// Any remaining replicas will be generated using the default template and will follow the default naming rules.
	//
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=name
	Instances []appsv1alpha1.InstanceTemplate `json:"instances,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name"`

	// Specifies the names of instances to be transitioned to offline status.
	//
	// Marking an instance as offline results in the following:
	//
	// 1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
	//    future reuse or data recovery, but it is no longer actively used.
	// 2. The ordinal number assigned to this instance is preserved, ensuring it remains unique
	//    and avoiding conflicts with new instances.
	//
	// Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
	// ordinal consistency within the Cluster.
	// Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
	// The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.
	//
	// +optional
	OfflineInstances []string `json:"offlineInstances,omitempty"`

	// Defines the sidecar containers that will be attached to the Component's main container.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=32
	// +listType=set
	// +optional
	Sidecars []string `json:"sidecars,omitempty"`

	// Determines whether metrics exporter information is annotated on the Component's headless Service.
	//
	// If set to true, the following annotations will be patched into the Service:
	//
	// - "monitor.kubeblocks.io/path"
	// - "monitor.kubeblocks.io/port"
	// - "monitor.kubeblocks.io/scheme"
	//
	// These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.
	//
	// +optional
	MonitorEnabled *bool `json:"monitorEnabled,omitempty"`
}

type SchedulingPolicy struct {
	// If specified, the Pod will be dispatched by specified scheduler.
	// If not specified, the Pod will be dispatched by default scheduler.
	//
	// +optional
	SchedulerName string `json:"schedulerName,omitempty"`

	// NodeSelector is a selector which must be true for the Pod to fit on a node.
	// Selector which must match a node's labels for the Pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	//
	// +optional
	// +mapType=atomic
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,
	// the scheduler simply schedules this Pod onto that node, assuming that it fits resource
	// requirements.
	//
	// +optional
	NodeName string `json:"nodeName,omitempty"`

	// Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.
	//
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// Allows Pods to be scheduled onto nodes with matching taints.
	// Each toleration in the array allows the Pod to tolerate node taints based on
	// specified `key`, `value`, `effect`, and `operator`.
	//
	// - The `key`, `value`, and `effect` identify the taint that the toleration matches.
	// - The `operator` determines how the toleration matches the taint.
	//
	// Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.
	//
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// TopologySpreadConstraints describes how a group of Pods ought to spread across topology
	// domains. Scheduler will schedule Pods in a way which abides by the constraints.
	// All topologySpreadConstraints are ANDed.
	//
	// +optional
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}

func (in *ClusterComponentSpec) TranslateTo() *appsv1alpha1.ClusterComponentSpec {
	return &appsv1alpha1.ClusterComponentSpec{
		Name:           in.Name,
		ServiceVersion: in.ServiceVersion,
		ServiceRefs:    in.ServiceRefs,
		EnabledLogs:    in.EnabledLogs,
		Replicas:       in.Replicas,
		// SchedulingPolicy: (&in.SchedulingPolicy).TranslateTo(),
		Resources:            in.Resources,
		VolumeClaimTemplates: in.VolumeClaimTemplates,
		Services:             in.Services,
		TLS:                  in.TLS,
		Issuer:               in.Issuer,
		ServiceAccountName:   in.ServiceAccountName,
		UserResourceRefs:     in.UserResourceRefs,
		Instances:            in.Instances,
		OfflineInstances:     in.OfflineInstances,
		Sidecars:             in.Sidecars,
		MonitorEnabled:       in.MonitorEnabled,
	}
}

func (in *BaseSpec) TranslateTo() *appsv1alpha1.ClusterSpec {
	return &appsv1alpha1.ClusterSpec{
		TerminationPolicy: in.TerminationPolicy,
		Services:          in.Services,
		// SchedulingPolicy: (&in.SchedulingPolicy).TranslateTo(),
		Backup: in.Backup,
	}
}
