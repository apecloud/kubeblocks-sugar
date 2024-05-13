package v1alpha1

import (
	"github.com/spf13/viper"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	dpv1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/constant"
)

type ClusterSpec struct {
	// Specifies the name of the ClusterTopology to be used when creating the Cluster.
	//
	// Override: Should be overridden to describe all the available topologies.
	//
	// +kubebuilder:validation:MaxLength=32
	// +optional
	Topology string `json:"topology,omitempty"`

	// Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster.
	// This field allows for detailed configuration of each Component within the Cluster.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=128
	// +kubebuilder:validation:XValidation:rule="self.all(x, size(self.filter(c, c.name == x.name)) == 1)",message="duplicated component"
	// +kubebuilder:validation:XValidation:rule="self.all(x, size(self.filter(c, has(c.componentDef))) == 0) || self.all(x, size(self.filter(c, has(c.componentDef))) == size(self))",message="two kinds of definition API can not be used simultaneously"
	// +optional
	ComponentSpecs []ClusterComponentSpec `json:"componentSpecs,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name"`

	// Defines a list of additional Services that are exposed by a Cluster.
	// This field allows Services of selected Components,
	// alongside Services defined with ComponentService.
	//
	// Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.
	//
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Services []ClusterService `json:"services,omitempty"`

	// Specifies the scheduling policy for the Cluster.
	//
	// +optional
	SchedulingPolicy *SchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// Specifies runtimeClassName for all Pods managed by this Cluster.
	//
	// +optional
	RuntimeClassName *string `json:"runtimeClassName,omitempty"`

	// Specifies the backup configuration of the Cluster.
	//
	// +optional
	Backup *ClusterBackup `json:"backup,omitempty"`

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
	TerminationPolicy TerminationPolicyType `json:"terminationPolicy"`
}

type ClusterStatus struct {
	// The most recent generation number of the Cluster object that has been observed by the controller.
	//
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// The current phase of the Cluster includes:
	// `Creating`, `Running`, `Updating`, `Stopping`, `Stopped`, `Deleting`, `Failed`, `Abnormal`.
	//
	// +optional
	Phase ClusterPhase `json:"phase,omitempty"`

	// Provides additional information about the current phase.
	//
	// +optional
	Message string `json:"message,omitempty"`

	// Records the current status information of all Components within the Cluster.
	//
	// +optional
	Components map[string]ClusterComponentStatus `json:"components,omitempty"`

	// Represents the generation number of the referenced ClusterDefinition.
	//
	// +optional
	ClusterDefGeneration int64 `json:"clusterDefGeneration,omitempty"`

	// Represents a list of detailed status of the Cluster object.
	// Each condition in the list provides real-time information about certain aspect of the Cluster object.
	//
	// This field is crucial for administrators and developers to monitor and respond to changes within the Cluster.
	// It provides a history of state transitions and a snapshot of the current state that can be used for
	// automated logic or direct inspection.
	//
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// ClusterComponentSpec defines the specification of a Component within a Cluster.
type ClusterComponentSpec struct {
	// Specifies the Component's name.
	// It's part of the Service DNS name and must comply with the IANA service naming rule.
	//
	// +kubebuilder:validation:MaxLength=22
	// +kubebuilder:validation:Pattern:=`^[a-z]([a-z0-9\-]*[a-z0-9])?$`
	Name string `json:"name"`

	// Specifies the name of the ComponentTopology to be used when creating the Component.
	// The ComponentTopology must present in the chosen ClusterTopology.
	//
	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Pattern:=`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`
	// +optional
	ComponentTopology string `json:"componentDef,omitempty"`

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
	ServiceRefs []ServiceRef `json:"serviceRefs,omitempty"`

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
	VolumeClaimTemplates []ClusterComponentVolumeClaimTemplate `json:"volumeClaimTemplates,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name"`

	// Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.
	//
	// +optional
	Services []ClusterComponentService `json:"services,omitempty"`

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
	Issuer *Issuer `json:"issuer,omitempty"`

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
	UserResourceRefs *UserResourceRefs `json:"userResourceRefs,omitempty"`

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
	Instances []InstanceTemplate `json:"instances,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name"`

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

// Issuer defines the TLS certificates issuer for the Cluster.
type Issuer struct {
	// The issuer for TLS certificates.
	// It only allows two enum values: `KubeBlocks` and `UserProvided`.
	//
	// - `KubeBlocks` indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used.
	// - `UserProvided` means that the user is responsible for providing their own CA, Cert, and Key.
	//   In this case, the user-provided CA certificate, server certificate, and private key will be used
	//   for TLS communication.
	//
	// +kubebuilder:validation:Enum={KubeBlocks, UserProvided}
	// +kubebuilder:default=KubeBlocks
	// +kubebuilder:validation:Required
	Name IssuerName `json:"name"`

	// SecretRef is the reference to the secret that contains user-provided certificates.
	// It is required when the issuer is set to `UserProvided`.
	//
	// +optional
	SecretRef *TLSSecretRef `json:"secretRef,omitempty"`
}

// TLSSecretRef defines Secret contains Tls certs
type TLSSecretRef struct {
	// Name of the Secret that contains user-provided certificates.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Key of CA cert in Secret
	// +kubebuilder:validation:Required
	CA string `json:"ca"`

	// Key of Cert in Secret
	// +kubebuilder:validation:Required
	Cert string `json:"cert"`

	// Key of TLS private key in Secret
	// +kubebuilder:validation:Required
	Key string `json:"key"`
}

// IssuerName defines the name of the TLS certificates issuer.
// +enum
// +kubebuilder:validation:Enum={KubeBlocks,UserProvided}
type IssuerName string

const (
	// IssuerKubeBlocks represents certificates that are signed by the KubeBlocks Operator.
	IssuerKubeBlocks IssuerName = "KubeBlocks"

	// IssuerUserProvided indicates that the user has provided their own CA-signed certificates.
	IssuerUserProvided IssuerName = "UserProvided"
)

type ServiceRef struct {
	// Specifies the identifier of the service reference declaration.
	// It corresponds to the serviceRefDeclaration name defined in either:
	//
	// - `componentDefinition.spec.serviceRefDeclarations[*].name`
	// - `clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name` (deprecated)
	//
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.
	// If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current
	// Cluster by default.
	//
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// References a service provided by another KubeBlocks Cluster.
	// It specifies the ClusterService and the account credentials needed for access.
	//
	// +optional
	ClusterServiceSelector *ServiceRefClusterSelector `json:"clusterServiceSelector,omitempty"`

	// Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.
	//
	// When referencing a service provided by external sources, a ServiceDescriptor object is required to establish
	// the service binding.
	// The `serviceDescriptor.spec.serviceKind` and `serviceDescriptor.spec.serviceVersion` should match the serviceKind
	// and serviceVersion declared in the definition.
	//
	// If both `cluster` and `serviceDescriptor` are specified, the `cluster` takes precedence.
	//
	// +optional
	ServiceDescriptor string `json:"serviceDescriptor,omitempty"`
}

type ServiceRefClusterSelector struct {
	// The name of the Cluster being referenced.
	//
	// +kubebuilder:validation:Required
	Cluster string `json:"cluster"`

	// Identifies a ClusterService from the list of Services defined in `cluster.spec.services` of the referenced Cluster.
	//
	// +optional
	Service *ServiceRefServiceSelector `json:"service,omitempty"`

	// Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster.
	// The SystemAccount should be defined in `componentDefinition.spec.systemAccounts`
	// of the Component providing the service in the referenced Cluster.
	//
	// +optional
	Credential *ServiceRefCredentialSelector `json:"credential,omitempty"`
}

type ServiceRefServiceSelector struct {
	// The name of the Component where the Service resides in.
	//
	// It is required when referencing a Component's Service.
	//
	// +optional
	Component string `json:"component,omitempty"`

	// The name of the Service to be referenced.
	//
	// Leave it empty to reference the default Service. Set it to "headless" to reference the default headless Service.
	//
	// If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,
	// and the resolved value will be presented in the following format: service1.name,service2.name...
	//
	// +kubebuilder:validation:Required
	Service string `json:"service"`

	// The port name of the Service to be referenced.
	//
	// If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.
	//
	// If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,
	// and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...
	//
	// +optional
	Port string `json:"port,omitempty"`
}

type ServiceRefCredentialSelector struct {
	// The name of the Component where the credential resides in.
	//
	// +kubebuilder:validation:Required
	Component string `json:"component"`

	// The name of the credential (SystemAccount) to reference.
	//
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type ClusterComponentVolumeClaimTemplate struct {
	// Refers to the name of a volumeMount defined in either:
	//
	// - `componentDefinition.spec.runtime.containers[*].volumeMounts`
	// - `clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts` (deprecated)
	//
	// The value of `name` must match the `name` field of a volumeMount specified in the corresponding `volumeMounts` array.
	//
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume
	// with the mount name specified in the `name` field.
	//
	// When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification
	// defined in the `spec` field. The PVC will be associated with the volume mount specified by the `name` field.
	//
	// +optional
	Spec PersistentVolumeClaimSpec `json:"spec,omitempty"`
}

func (r *ClusterComponentVolumeClaimTemplate) toVolumeClaimTemplate() corev1.PersistentVolumeClaimTemplate {
	return corev1.PersistentVolumeClaimTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name: r.Name,
		},
		Spec: r.Spec.ToV1PersistentVolumeClaimSpec(),
	}
}

type PersistentVolumeClaimSpec struct {
	// Contains the desired access modes the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.
	//
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,1,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`

	// Represents the minimum resources the volume should have.
	// If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that
	// are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.
	//
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources"`

	// The name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.
	//
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty" protobuf:"bytes,5,opt,name=storageClassName"`

	// Defines what type of volume is required by the claim, either Block or Filesystem.
	//
	// +optional
	VolumeMode *corev1.PersistentVolumeMode `json:"volumeMode,omitempty" protobuf:"bytes,6,opt,name=volumeMode,casttype=PersistentVolumeMode"`
}

// ToV1PersistentVolumeClaimSpec converts to corev1.PersistentVolumeClaimSpec.
func (r *PersistentVolumeClaimSpec) ToV1PersistentVolumeClaimSpec() corev1.PersistentVolumeClaimSpec {
	return corev1.PersistentVolumeClaimSpec{
		AccessModes:      r.AccessModes,
		Resources:        r.Resources,
		StorageClassName: r.getStorageClassName(viper.GetString(constant.CfgKeyDefaultStorageClass)),
		VolumeMode:       r.VolumeMode,
	}
}

// getStorageClassName returns PersistentVolumeClaimSpec.StorageClassName if a value is assigned; otherwise,
// it returns preferSC argument.
func (r *PersistentVolumeClaimSpec) getStorageClassName(preferSC string) *string {
	if r.StorageClassName != nil && *r.StorageClassName != "" {
		return r.StorageClassName
	}
	if preferSC != "" {
		return &preferSC
	}
	return nil
}

type Affinity struct {
	// Specifies the anti-affinity level of Pods within a Component.
	// It determines how pods should be spread across nodes to improve availability and performance.
	// It can have the following values: `Preferred` and `Required`.
	// The default value is `Preferred`.
	//
	// +kubebuilder:default=Preferred
	// +optional
	PodAntiAffinity PodAntiAffinity `json:"podAntiAffinity,omitempty"`

	// Represents the key of node labels used to define the topology domain for Pod anti-affinity
	// and Pod spread constraints.
	//
	// In K8s, a topology domain is a set of nodes that have the same value for a specific label key.
	// Nodes with labels containing any of the specified TopologyKeys and identical values are considered
	// to be in the same topology domain.
	//
	// Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of
	// topology in the ClusterDefinition.
	//
	// When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the
	// Pod on nodes with different values for the specified TopologyKeys.
	// This ensures that Pods are spread across different topology domains, promoting high availability and
	// reducing the impact of node failures.
	//
	// Some well-known label keys, such as `kubernetes.io/hostname` and `topology.kubernetes.io/zone`,
	// are often used as TopologyKey.
	// These keys represent the hostname and zone of a node, respectively.
	// By including these keys in the TopologyKeys list, Pods will be spread across nodes with
	// different hostnames or zones.
	//
	// In addition to the well-known keys, users can also specify custom label keys as TopologyKeys.
	// This allows for more flexible and custom topology definitions based on the specific needs
	// of the application or environment.
	//
	// The TopologyKeys field is a slice of strings, where each string represents a label key.
	// The order of the keys in the slice does not matter.
	//
	// +listType=set
	// +optional
	TopologyKeys []string `json:"topologyKeys,omitempty"`

	// Indicates the node labels that must be present on nodes for pods to be scheduled on them.
	// It is a map where the keys are the label keys and the values are the corresponding label values.
	// Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.
	//
	// For example, if NodeLabels is set to {"nodeType": "ssd", "environment": "production"},
	// pods will only be scheduled on nodes that have both the "nodeType" label with value "ssd"
	// and the "environment" label with value "production".
	//
	// This field allows users to control Pod placement based on specific node labels.
	// It can be used to ensure that Pods are scheduled on nodes with certain characteristics,
	// such as specific hardware (e.g., SSD), environment (e.g., production, staging),
	// or any other custom labels assigned to nodes.
	//
	// +optional
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`

	// Determines the level of resource isolation between Pods.
	// It can have the following values: `SharedNode` and `DedicatedNode`.
	//
	// - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s.
	// - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node.
	//   In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node.
	//   Which provides a higher level of isolation and resource guarantee for Pods.
	//
	//  The default value is `SharedNode`.
	//
	// +kubebuilder:default=SharedNode
	// +optional
	Tenancy TenancyType `json:"tenancy,omitempty"`
}

// PodAntiAffinity defines the pod anti-affinity strategy.
//
// This strategy determines how pods are scheduled in relation to other pods, with the aim of either spreading pods
// across nodes (Preferred) or ensuring that certain pods do not share a node (Required).
//
// +enum
// +kubebuilder:validation:Enum={Preferred,Required}
type PodAntiAffinity string

const (
	// Preferred indicates that the scheduler will try to enforce the anti-affinity rules, but it will not guarantee it.
	Preferred PodAntiAffinity = "Preferred"

	// Required indicates that the scheduler must enforce the anti-affinity rules and will not schedule the pods unless
	// the rules are met.
	Required PodAntiAffinity = "Required"
)

// TenancyType defines the type of tenancy for cluster tenant resources.
//
// +enum
// +kubebuilder:validation:Enum={SharedNode,DedicatedNode}
type TenancyType string

const (
	// SharedNode means multiple pods may share the same node.
	SharedNode TenancyType = "SharedNode"

	// DedicatedNode means each pod runs on their own dedicated node.
	DedicatedNode TenancyType = "DedicatedNode"
)

type ClusterComponentService struct {
	// References the ComponentService name defined in the `componentDefinition.spec.services[*].name`.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=25
	Name string `json:"name"`

	// Determines how the Service is exposed. Valid options are `ClusterIP`, `NodePort`, and `LoadBalancer`.
	//
	// - `ClusterIP` allocates a Cluster-internal IP address for load-balancing to endpoints.
	//    Endpoints are determined by the selector or if that is not specified,
	//    they are determined by manual construction of an Endpoints object or EndpointSlice objects.
	// - `NodePort` builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP.
	// - `LoadBalancer` builds on NodePort and creates an external load-balancer (if supported in the current cloud)
	//    which routes to the same endpoints as the ClusterIP.
	//
	// Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for ClusterComponentService.
	//
	// For more info, see:
	// https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.
	//
	// +kubebuilder:default=ClusterIP
	// +kubebuilder:validation:Enum={ClusterIP,NodePort,LoadBalancer}
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	ServiceType corev1.ServiceType `json:"serviceType,omitempty"`

	// If ServiceType is LoadBalancer, cloud provider related parameters can be put here.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.
	//
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Indicates whether to generate individual Services for each Pod.
	// If set to true, a separate Service will be created for each Pod in the Cluster.
	//
	// +optional
	PodService *bool `json:"podService,omitempty"`
}

// ClusterService defines a service that is exposed externally, allowing entities outside the cluster to access it.
// For example, external applications, or other Clusters.
// And another Cluster managed by the same KubeBlocks operator can resolve the address exposed by a ClusterService
// using the `serviceRef` field.
//
// When a Component needs to access another Cluster's ClusterService using the `serviceRef` field,
// it must also define the service type and version information in the `componentDefinition.spec.serviceRefDeclarations`
// section.
type ClusterService struct {
	Service `json:",inline"`

	// Extends the ServiceSpec.Selector by allowing the specification of a component, to be used as a selector for the service.
	// Note that this and the `shardingSelector` are mutually exclusive and cannot be set simultaneously.
	//
	// +optional
	ComponentSelector string `json:"componentSelector,omitempty"`
}

type Service struct {
	// Name defines the name of the service.
	// otherwise, it indicates the name of the service.
	// Others can refer to this service by its name. (e.g., connection credential)
	// Cannot be updated.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=25
	Name string `json:"name"`

	// ServiceName defines the name of the underlying service object.
	// If not specified, the default service name with different patterns will be used:
	//
	// - CLUSTER_NAME: for cluster-level services
	// - CLUSTER_NAME-COMPONENT_NAME: for component-level services
	//
	// Only one default service name is allowed.
	// Cannot be updated.
	//
	// +optional
	ServiceName string `json:"serviceName,omitempty"`

	// If ServiceType is LoadBalancer, cloud provider related parameters can be put here
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.
	//
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Spec defines the behavior of a service.
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	//
	// +optional
	Spec corev1.ServiceSpec `json:"spec,omitempty"`

	// Extends the above `serviceSpec.selector` by allowing you to specify defined role as selector for the service.
	// When `roleSelector` is set, it adds a label selector "kubeblocks.io/role: {roleSelector}"
	// to the `serviceSpec.selector`.
	// Example usage:
	//
	//	  roleSelector: "leader"
	//
	// In this example, setting `roleSelector` to "leader" will add a label selector
	// "kubeblocks.io/role: leader" to the `serviceSpec.selector`.
	// This means that the service will select and route traffic to Pods with the label
	// "kubeblocks.io/role" set to "leader".
	//
	// Note that if `podService` sets to true, RoleSelector will be ignored.
	// The `podService` flag takes precedence over `roleSelector` and generates a service for each Pod.
	//
	// +optional
	RoleSelector string `json:"roleSelector,omitempty"`
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

type ClusterBackup struct {
	// Specifies whether automated backup is enabled for the Cluster.
	//
	// +kubebuilder:default=false
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	// Determines the duration to retain backups. Backups older than this period are automatically removed.
	//
	// For example, RetentionPeriod of `30d` will keep only the backups of last 30 days.
	// Sample duration format:
	//
	// - years: 	2y
	// - months: 	6mo
	// - days: 		30d
	// - hours: 	12h
	// - minutes: 	30m
	//
	// You can also combine the above durations. For example: 30d12h30m.
	// Default value is 7d.
	//
	// +kubebuilder:default="7d"
	// +optional
	RetentionPeriod dpv1alpha1.RetentionPeriod `json:"retentionPeriod,omitempty"`

	// Specifies the backup method to use, as defined in backupPolicy.
	//
	// +kubebuilder:validation:Required
	Method string `json:"method"`

	// The cron expression for the schedule. The timezone is in UTC. See https://en.wikipedia.org/wiki/Cron.
	//
	// +optional
	CronExpression string `json:"cronExpression,omitempty"`

	// Specifies the maximum time in minutes that the system will wait to start a missed backup job.
	// If the scheduled backup time is missed for any reason, the backup job must start within this deadline.
	// Values must be between 0 (immediate execution) and 1440 (one day).
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1440
	// +optional
	StartingDeadlineMinutes *int64 `json:"startingDeadlineMinutes,omitempty"`

	// Specifies the name of the backupRepo. If not set, the default backupRepo will be used.
	//
	// +optional
	RepoName string `json:"repoName,omitempty"`

	// Specifies whether to enable point-in-time recovery.
	//
	// +kubebuilder:default=false
	// +optional
	PITREnabled *bool `json:"pitrEnabled,omitempty"`
}

// ResourceMeta encapsulates metadata and configuration for referencing ConfigMaps and Secrets as volumes.
type ResourceMeta struct {
	// Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern:=`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`
	Name string `json:"name"`

	// MountPoint is the filesystem path where the volume will be mounted.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:Pattern:=`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`
	MountPoint string `json:"mountPoint"`

	// SubPath specifies a path within the volume from which to mount.
	//
	// +optional
	SubPath string `json:"subPath,omitempty"`

	// AsVolumeFrom lists the names of containers in which the volume should be mounted.
	//
	// +listType=set
	// +optional
	AsVolumeFrom []string `json:"asVolumeFrom,omitempty"`
}

// SecretRef defines a reference to a Secret.
type SecretRef struct {
	ResourceMeta `json:",inline"`

	// Secret specifies the Secret to be mounted as a volume.
	//
	// +kubebuilder:validation:Required
	Secret corev1.SecretVolumeSource `json:"secret"`
}

// ConfigMapRef defines a reference to a ConfigMap.
type ConfigMapRef struct {
	ResourceMeta `json:",inline"`

	// ConfigMap specifies the ConfigMap to be mounted as a volume.
	//
	// +kubebuilder:validation:Required
	ConfigMap corev1.ConfigMapVolumeSource `json:"configMap"`
}

// UserResourceRefs defines references to user-defined Secrets and ConfigMaps.
type UserResourceRefs struct {
	// SecretRefs defines the user-defined Secrets.
	//
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=name
	// +optional
	SecretRefs []SecretRef `json:"secretRefs,omitempty"`

	// ConfigMapRefs defines the user-defined ConfigMaps.
	//
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=name
	// +optional
	ConfigMapRefs []ConfigMapRef `json:"configMapRefs,omitempty"`
}

// InstanceTemplate allows customization of individual replica configurations in a Component.
type InstanceTemplate struct {
	// Name specifies the unique name of the instance Pod created using this InstanceTemplate.
	// This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal
	// using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.
	// The specified name overrides any default naming conventions or patterns.
	//
	// +kubebuilder:validation:MaxLength=54
	// +kubebuilder:validation:Pattern:=`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Specifies the number of instances (Pods) to create from this InstanceTemplate.
	// This field allows setting how many replicated instances of the Component,
	// with the specific overrides in the InstanceTemplate, are created.
	// The default value is 1. A value of 0 disables instance creation.
	//
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=0
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// Specifies a map of key-value pairs to be merged into the Pod's existing annotations.
	// Existing keys will have their values overwritten, while new keys will be added to the annotations.
	//
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Specifies a map of key-value pairs that will be merged into the Pod's existing labels.
	// Values for existing keys will be overwritten, and new keys will be added.
	//
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Specifies an override for the first container's image in the Pod.
	//
	// +optional
	Image *string `json:"image,omitempty"`

	// Specifies the name of the node where the Pod should be scheduled.
	// If set, the Pod will be directly assigned to the specified node, bypassing the Kubernetes scheduler.
	// This is useful for controlling Pod placement on specific nodes.
	//
	// Important considerations:
	// - `nodeName` bypasses default scheduling constraints (e.g., resource requirements, node selectors, affinity rules).
	// - It is the user's responsibility to ensure the node is suitable for the Pod.
	// - If the node is unavailable, the Pod will remain in "Pending" state until the node is available or the Pod is deleted.
	//
	// +optional
	NodeName *string `json:"nodeName,omitempty"`

	// Defines NodeSelector to override.
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Tolerations specifies a list of tolerations to be applied to the Pod, allowing it to tolerate node taints.
	// This field can be used to add new tolerations or override existing ones.
	//
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Specifies an override for the resource requirements of the first container in the Pod.
	// This field allows for customizing resource allocation (CPU, memory, etc.) for the container.
	//
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Defines Env to override.
	// Add new or override existing envs.
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Defines Volumes to override.
	// Add new or override existing volumes.
	// +optional
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Defines VolumeMounts to override.
	// Add new or override existing volume mounts of the first container in the Pod.
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Defines VolumeClaimTemplates to override.
	// Add new or override existing volume claim templates.
	// +optional
	VolumeClaimTemplates []ClusterComponentVolumeClaimTemplate `json:"volumeClaimTemplates,omitempty"`
}

// TerminationPolicyType defines termination policy types.
//
// +enum
// +kubebuilder:validation:Enum={DoNotTerminate,Halt,Delete,WipeOut}
type TerminationPolicyType string

const (
	// DoNotTerminate will block delete operation.
	DoNotTerminate TerminationPolicyType = "DoNotTerminate"

	// Halt will delete workload resources such as statefulset, deployment workloads but keep PVCs.
	Halt TerminationPolicyType = "Halt"

	// Delete is based on Halt and deletes PVCs.
	Delete TerminationPolicyType = "Delete"

	// WipeOut is based on Delete and wipe out all volume snapshots and snapshot data from backup storage location.
	WipeOut TerminationPolicyType = "WipeOut"
)

// ClusterPhase defines the phase of the Cluster within the .status.phase field.
//
// +enum
// +kubebuilder:validation:Enum={Creating,Running,Updating,Stopping,Stopped,Deleting,Failed,Abnormal}
type ClusterPhase string

const (
	// CreatingClusterPhase represents all components are in `Creating` phase.
	CreatingClusterPhase ClusterPhase = "Creating"

	// RunningClusterPhase represents all components are in `Running` phase, indicates that the cluster is functioning properly.
	RunningClusterPhase ClusterPhase = "Running"

	// UpdatingClusterPhase represents all components are in `Creating`, `Running` or `Updating` phase, and at least one
	// component is in `Creating` or `Updating` phase, indicates that the cluster is undergoing an update.
	UpdatingClusterPhase ClusterPhase = "Updating"

	// StoppingClusterPhase represents at least one component is in `Stopping` phase, indicates that the cluster is in
	// the process of stopping.
	StoppingClusterPhase ClusterPhase = "Stopping"

	// StoppedClusterPhase represents all components are in `Stopped` phase, indicates that the cluster has stopped and
	// is not providing any functionality.
	StoppedClusterPhase ClusterPhase = "Stopped"

	// DeletingClusterPhase indicates the cluster is being deleted.
	DeletingClusterPhase ClusterPhase = "Deleting"

	// FailedClusterPhase represents all components are in `Failed` phase, indicates that the cluster is unavailable.
	FailedClusterPhase ClusterPhase = "Failed"

	// AbnormalClusterPhase represents some components are in `Failed` or `Abnormal` phase, indicates that the cluster
	// is in a fragile state and troubleshooting is required.
	AbnormalClusterPhase ClusterPhase = "Abnormal"
)

// ClusterComponentPhase defines the phase of a cluster component as represented in cluster.status.components.phase field.
//
// +enum
// +kubebuilder:validation:Enum={Creating,Running,Updating,Stopping,Stopped,Deleting,Failed,Abnormal}
type ClusterComponentPhase string

const (
	// CreatingClusterCompPhase indicates the component is being created.
	CreatingClusterCompPhase ClusterComponentPhase = "Creating"

	// RunningClusterCompPhase indicates the component has more than zero replicas, and all pods are up-to-date and
	// in a 'Running' state.
	RunningClusterCompPhase ClusterComponentPhase = "Running"

	// UpdatingClusterCompPhase indicates the component has more than zero replicas, and there are no failed pods,
	// it is currently being updated.
	UpdatingClusterCompPhase ClusterComponentPhase = "Updating"

	// StoppingClusterCompPhase indicates the component has zero replicas, and there are pods that are terminating.
	StoppingClusterCompPhase ClusterComponentPhase = "Stopping"

	// StoppedClusterCompPhase indicates the component has zero replicas, and all pods have been deleted.
	StoppedClusterCompPhase ClusterComponentPhase = "Stopped"

	// DeletingClusterCompPhase indicates the component is currently being deleted.
	DeletingClusterCompPhase ClusterComponentPhase = "Deleting"

	// FailedClusterCompPhase indicates the component has more than zero replicas, but there are some failed pods.
	// The component is not functioning.
	FailedClusterCompPhase ClusterComponentPhase = "Failed"

	// AbnormalClusterCompPhase indicates the component has more than zero replicas, but there are some failed pods.
	// The component is functioning, but it is in a fragile state.
	AbnormalClusterCompPhase ClusterComponentPhase = "Abnormal"
)

type ComponentMessageMap map[string]string

// ClusterComponentStatus records Component status.
type ClusterComponentStatus struct {
	// Specifies the current state of the Component.
	Phase ClusterComponentPhase `json:"phase,omitempty"`

	// Records detailed information about the Component in its current phase.
	// The keys are either podName, deployName, or statefulSetName, formatted as 'ObjectKind/Name'.
	//
	// +optional
	Message ComponentMessageMap `json:"message,omitempty"`

	// Checks if all Pods of the Component are ready.
	//
	// +optional
	PodsReady *bool `json:"podsReady,omitempty"`

	// Indicates the time when all Component Pods became ready.
	// This is the readiness time of the last Component Pod.
	//
	// +optional
	PodsReadyTime *metav1.Time `json:"podsReadyTime,omitempty"`

	// Represents the status of the members.
	//
	// +optional
	MembersStatus []MemberStatus `json:"membersStatus,omitempty"`
}

type MemberStatus struct {
	// Represents the name of the pod.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:default=Unknown
	PodName string `json:"podName"`

	// Defines the role of the replica in the cluster.
	//
	// +optional
	ReplicaRole *ReplicaRole `json:"role,omitempty"`
}

type ReplicaRole struct {

	// Defines the role name of the replica.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:default=leader
	Name string `json:"name"`

	// Specifies the service capabilities of this member.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:default=ReadWrite
	// +kubebuilder:validation:Enum={None, Readonly, ReadWrite}
	AccessMode AccessMode `json:"accessMode"`

	// Indicates if this member has voting rights.
	//
	// +kubebuilder:default=true
	// +optional
	CanVote bool `json:"canVote"`

	// Determines if this member is the leader.
	//
	// +kubebuilder:default=false
	// +optional
	IsLeader bool `json:"isLeader"`
}

// AccessMode defines SVC access mode enums.
// +enum
type AccessMode string

const (
	ReadWriteMode AccessMode = "ReadWrite"
	ReadonlyMode  AccessMode = "Readonly"
	NoneMode      AccessMode = "None"
)
