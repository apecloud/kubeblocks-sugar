# API Reference

## Packages
- [sugar.kubeblocks.io/v1alpha1](#sugarkubeblocksiov1alpha1)


## sugar.kubeblocks.io/v1alpha1

Package v1alpha1 contains API Schema definitions for the sugar v1alpha1 API group

### Resource Types
- [ApeCloudMySQL](#apecloudmysql)
- [ApeCloudMySQLList](#apecloudmysqllist)



#### ApeCloudMySQL



ApeCloudMySQL is the Schema for the apecloudmysqls API



_Appears in:_
- [ApeCloudMySQLList](#apecloudmysqllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `sugar.kubeblocks.io/v1alpha1` | | |
| `kind` _string_ | `ApeCloudMySQL` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ApeCloudMySQLSpec](#apecloudmysqlspec)_ |  |  |  |


#### ApeCloudMySQLList



ApeCloudMySQLList contains a list of ApeCloudMySQL





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `sugar.kubeblocks.io/v1alpha1` | | |
| `kind` _string_ | `ApeCloudMySQLList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ApeCloudMySQL](#apecloudmysql) array_ |  |  |  |


#### ApeCloudMySQLSpec



ApeCloudMySQLSpec defines the desired state of ApeCloudMySQL



_Appears in:_
- [ApeCloudMySQL](#apecloudmysql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `services` _ClusterService array_ | Defines a list of additional Services that are exposed by a Cluster.<br />This field allows Services of selected Components,<br />alongside Services defined with ComponentService.<br /><br />Services defined here can be referenced by other clusters using the ServiceRefClusterSelector. |  |  |
| `schedulingPolicy` _[SchedulingPolicy](#schedulingpolicy)_ | Specifies the scheduling policy for the Cluster. |  |  |
| `backup` _[ClusterBackup](#clusterbackup)_ | Specifies the backup configuration of the Cluster. |  |  |
| `terminationPolicy` _[TerminationPolicyType](#terminationpolicytype)_ | Specifies the behavior when a Cluster is deleted.<br />It defines how resources, data, and backups associated with a Cluster are managed during termination.<br />Choose a policy based on the desired level of resource cleanup and data preservation:<br /><br />- `DoNotTerminate`: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.<br />- `Halt`: Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs),<br />  allowing for data preservation while stopping other operations.<br />- `Delete`: Extends the `Halt` policy by also removing PVCs, leading to a thorough cleanup while<br />  removing all persistent data.<br />- `WipeOut`: An aggressive policy that deletes all Cluster resources, including volume snapshots and<br />  backups in external storage.<br />  This results in complete data removal and should be used cautiously, primarily in non-production environments<br />  to avoid irreversible data loss.<br /><br />Warning: Choosing an inappropriate termination policy can result in data loss.<br />The `WipeOut` policy is particularly risky in production environments due to its irreversible nature. |  | Required: {} <br /> |
| `mysqlSpec` _[BaseComponentSpec](#basecomponentspec)_ | Specified the MySQL Component Spec. |  |  |
| `proxySpec` _[BaseComponentSpec](#basecomponentspec)_ | Specified the Proxy Component Spec. |  |  |






#### BaseComponentSpec



BaseComponentSpec defines the specification of a Component within a Cluster.



_Appears in:_
- [ApeCloudMySQLSpec](#apecloudmysqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `serviceVersion` _string_ | ServiceVersion specifies the version of the Service expected to be provisioned by this Component.<br />The version should follow the syntax and semantics of the "Semantic Versioning" specification (http://semver.org/).<br />If no version is specified, the latest available version will be used. |  | MaxLength: 32 <br /> |
| `serviceRefs` _ServiceRef array_ | Defines a list of ServiceRef for a Component, enabling access to both external services and<br />Services provided by other Clusters.<br /><br />Types of services:<br /><br />- External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;<br />  Require a ServiceDescriptor for connection details.<br />- Services provided by a Cluster: Managed by the same KubeBlocks operator;<br />  identified using Cluster, Component and Service names.<br /><br />ServiceRefs with identical `serviceRef.name` in the same Cluster are considered the same.<br /><br />Example:<br />```yaml<br />serviceRefs:<br />  - name: "redis-sentinel"<br />    serviceDescriptor:<br />      name: "external-redis-sentinel"<br />  - name: "postgres-cluster"<br />    clusterServiceSelector:<br />      cluster: "my-postgres-cluster"<br />      service:<br />        component: "postgresql"<br />```<br />The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster. |  |  |
| `enabledLogs` _string array_ | Specifies which types of logs should be collected for the Component.<br />The log types are defined in the `componentDefinition.spec.logConfigs` field with the LogConfig entries.<br /><br />The elements in the `enabledLogs` array correspond to the names of the LogConfig entries.<br />For example, if the `componentDefinition.spec.logConfigs` defines LogConfig entries with<br />names "slow_query_log" and "error_log",<br />you can enable the collection of these logs by including their names in the `enabledLogs` array:<br />```yaml<br />enabledLogs:<br />- slow_query_log<br />- error_log<br />``` |  |  |
| `replicas` _integer_ | Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing. | 1 | Minimum: 0 <br />Required: {} <br /> |
| `schedulingPolicy` _[SchedulingPolicy](#schedulingpolicy)_ | Specifies the scheduling policy for the Component. |  |  |
| `resources` _[ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#resourcerequirements-v1-core)_ | Specifies the resources required by the Component.<br />It allows defining the CPU, memory requirements and limits for the Component's containers. |  |  |
| `volumeClaimTemplates` _ClusterComponentVolumeClaimTemplate array_ | Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component.<br />Each template specifies the desired characteristics of a persistent volume, such as storage class,<br />size, and access modes.<br />These templates are used to dynamically provision persistent volumes for the Component. |  |  |
| `services` _ClusterComponentService array_ | Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients. |  |  |
| `tls` _boolean_ | A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)<br />for secure communication.<br />When set to true, the Component will be configured to use TLS encryption for its network connections.<br />This ensures that the data transmitted between the Component and its clients or other Components is encrypted<br />and protected from unauthorized access.<br />If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys,<br />to properly set up the secure communication channel. |  |  |
| `issuer` _[Issuer](#issuer)_ | Specifies the configuration for the TLS certificates issuer.<br />It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.<br />The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.<br />Required when TLS is enabled. |  |  |
| `serviceAccountName` _string_ | Specifies the name of the ServiceAccount required by the running Component.<br />This ServiceAccount is used to grant necessary permissions for the Component's Pods to interact<br />with other Kubernetes resources, such as modifying Pod labels or sending events.<br /><br />Defaults:<br />If not specified, KubeBlocks automatically assigns a default ServiceAccount named "kb-{cluster.name}",<br />bound to a default role installed together with KubeBlocks.<br /><br />Future Changes:<br />Future versions might change the default ServiceAccount creation strategy to one per Component,<br />potentially revising the naming to "kb-{cluster.name}-{component.name}".<br /><br />Users can override the automatic ServiceAccount assignment by explicitly setting the name of<br />an existed ServiceAccount in this field. |  |  |
| `userResourceRefs` _[UserResourceRefs](#userresourcerefs)_ | Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes<br />in the Cluster's Pods.<br />This is useful in scenarios where users need to provide additional resources to the Cluster, such as:<br /><br />- Mounting custom scripts or configuration files during Cluster startup.<br />- Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster. |  |  |
| `instances` _InstanceTemplate array_ | Allows for the customization of configuration values for each instance within a Component.<br />An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).<br />While instances typically share a common configuration as defined in the BaseComponentSpec,<br />they can require unique settings in various scenarios:<br /><br />For example:<br />- A database Component might require different resource allocations for primary and secondary instances,<br />  with primaries needing more resources.<br />- During a rolling upgrade, a Component may first update the image for one or a few instances,<br />  and then update the remaining instances after verifying that the updated instances are functioning correctly.<br /><br />InstanceTemplate allows for specifying these unique configurations per instance.<br />Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),<br />starting with an ordinal of 0.<br />It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.<br /><br />The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component.<br />Any remaining replicas will be generated using the default template and will follow the default naming rules. |  |  |
| `offlineInstances` _string array_ | Specifies the names of instances to be transitioned to offline status.<br /><br />Marking an instance as offline results in the following:<br /><br />1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential<br />   future reuse or data recovery, but it is no longer actively used.<br />2. The ordinal number assigned to this instance is preserved, ensuring it remains unique<br />   and avoiding conflicts with new instances.<br /><br />Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining<br />ordinal consistency within the Cluster.<br />Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.<br />The administrator must manually manage the cleanup and removal of these resources when they are no longer needed. |  |  |
| `sidecars` _string array_ | Defines the sidecar containers that will be attached to the Component's main container. |  | MaxItems: 32 <br />MinItems: 1 <br /> |
| `monitorEnabled` _boolean_ | Determines whether metrics exporter information is annotated on the Component's headless Service.<br /><br />If set to true, the following annotations will be patched into the Service:<br /><br />- "monitor.kubeblocks.io/path"<br />- "monitor.kubeblocks.io/port"<br />- "monitor.kubeblocks.io/scheme"<br /><br />These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter. |  |  |


#### BaseSpec







_Appears in:_
- [ApeCloudMySQLSpec](#apecloudmysqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `services` _ClusterService array_ | Defines a list of additional Services that are exposed by a Cluster.<br />This field allows Services of selected Components,<br />alongside Services defined with ComponentService.<br /><br />Services defined here can be referenced by other clusters using the ServiceRefClusterSelector. |  |  |
| `schedulingPolicy` _[SchedulingPolicy](#schedulingpolicy)_ | Specifies the scheduling policy for the Cluster. |  |  |
| `backup` _[ClusterBackup](#clusterbackup)_ | Specifies the backup configuration of the Cluster. |  |  |
| `terminationPolicy` _[TerminationPolicyType](#terminationpolicytype)_ | Specifies the behavior when a Cluster is deleted.<br />It defines how resources, data, and backups associated with a Cluster are managed during termination.<br />Choose a policy based on the desired level of resource cleanup and data preservation:<br /><br />- `DoNotTerminate`: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.<br />- `Halt`: Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs),<br />  allowing for data preservation while stopping other operations.<br />- `Delete`: Extends the `Halt` policy by also removing PVCs, leading to a thorough cleanup while<br />  removing all persistent data.<br />- `WipeOut`: An aggressive policy that deletes all Cluster resources, including volume snapshots and<br />  backups in external storage.<br />  This results in complete data removal and should be used cautiously, primarily in non-production environments<br />  to avoid irreversible data loss.<br /><br />Warning: Choosing an inappropriate termination policy can result in data loss.<br />The `WipeOut` policy is particularly risky in production environments due to its irreversible nature. |  | Required: {} <br /> |


#### BaseStatus







_Appears in:_
- [ApeCloudMySQLStatus](#apecloudmysqlstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `observedGeneration` _integer_ | The most recent generation number of the Cluster object that has been observed by the controller. |  |  |
| `phase` _[ClusterPhase](#clusterphase)_ | The current phase of the Cluster includes:<br />`Creating`, `Running`, `Updating`, `Stopping`, `Stopped`, `Deleting`, `Failed`, `Abnormal`. |  | Enum: [Creating Running Updating Stopping Stopped Deleting Failed Abnormal] <br /> |
| `message` _string_ | Provides additional information about the current phase. |  |  |
| `components` _object (keys:string, values:[ClusterComponentStatus](#clustercomponentstatus))_ | Records the current status information of all Components within the Cluster. |  |  |
| `clusterDefGeneration` _integer_ | Represents the generation number of the referenced ClusterDefinition. |  |  |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#condition-v1-meta) array_ | Represents a list of detailed status of the Cluster object.<br />Each condition in the list provides real-time information about certain aspect of the Cluster object.<br /><br />This field is crucial for administrators and developers to monitor and respond to changes within the Cluster.<br />It provides a history of state transitions and a snapshot of the current state that can be used for<br />automated logic or direct inspection. |  |  |


#### SchedulingPolicy







_Appears in:_
- [ApeCloudMySQLSpec](#apecloudmysqlspec)
- [BaseComponentSpec](#basecomponentspec)
- [BaseSpec](#basespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `schedulerName` _string_ | If specified, the Pod will be dispatched by specified scheduler.<br />If not specified, the Pod will be dispatched by default scheduler. |  |  |
| `nodeSelector` _object (keys:string, values:string)_ | NodeSelector is a selector which must be true for the Pod to fit on a node.<br />Selector which must match a node's labels for the Pod to be scheduled on that node.<br />More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |  |  |
| `nodeName` _string_ | NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,<br />the scheduler simply schedules this Pod onto that node, assuming that it fits resource<br />requirements. |  |  |
| `affinity` _[Affinity](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#affinity-v1-core)_ | Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity. |  |  |
| `tolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#toleration-v1-core) array_ | Allows Pods to be scheduled onto nodes with matching taints.<br />Each toleration in the array allows the Pod to tolerate node taints based on<br />specified `key`, `value`, `effect`, and `operator`.<br /><br />- The `key`, `value`, and `effect` identify the taint that the toleration matches.<br />- The `operator` determines how the toleration matches the taint.<br /><br />Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes. |  |  |
| `topologySpreadConstraints` _[TopologySpreadConstraint](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#topologyspreadconstraint-v1-core) array_ | TopologySpreadConstraints describes how a group of Pods ought to spread across topology<br />domains. Scheduler will schedule Pods in a way which abides by the constraints.<br />All topologySpreadConstraints are ANDed. |  |  |


