
# Table: k8s_batch_jobs
Job represents the configuration of a single job.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|context|text|Name of the context from k8s configuration.|
|owner_references|jsonb||
|name|text|Name must be unique within a namespace|
|generate_name|text|GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed|
|namespace|text|Namespace defines the space within which each name must be unique|
|self_link|text|SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional|
|uid|text|UID is the unique in time and space value for this object|
|resource_version|text|An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed|
|generation|bigint|A sequence number representing a specific generation of the desired state. Populated by the system|
|deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate before it will be removed from the system|
|labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|finalizers|text[]|Must be empty before the object is deleted from the registry|
|cluster_name|text|The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional|
|managed_fields|jsonb|ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow|
|parallelism|integer|Specifies the maximum desired number of pods the job should run at any given time|
|completions|integer|Specifies the desired number of successfully finished pods the job should be run with|
|active_deadline_seconds|bigint|Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer|
|backoff_limit|integer|Specifies the number of retries before marking this job failed. Defaults to 6 +optional|
|selector_match_labels|jsonb|matchLabels is a map of {key,value} pairs|
|manual_selector|boolean|manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template|
|template_name|text|Name must be unique within a namespace|
|template_generate_name|text|GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed|
|template_namespace|text|Namespace defines the space within which each name must be unique|
|template_self_link|text|SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional|
|template_uid|text|UID is the unique in time and space value for this object|
|template_resource_version|text|An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed|
|template_generation|bigint|A sequence number representing a specific generation of the desired state. Populated by the system|
|template_deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate before it will be removed from the system|
|template_labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|template_annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|template_owner_references|jsonb|List of objects depended by this object|
|template_finalizers|text[]|Must be empty before the object is deleted from the registry|
|template_cluster_name|text|The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional|
|template_managed_fields|jsonb|ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow|
|template_restart_policy|text|Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy +optional|
|template_termination_grace_period_seconds|bigint|Optional duration in seconds the pod needs to terminate gracefully|
|template_active_deadline_seconds|bigint|Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer. +optional|
|template_dns_policy|text|Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. +optional|
|template_node_selector|jsonb|NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ +optional +mapType=atomic|
|template_service_account_name|text|ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/ +optional|
|template_deprecated_service_account|text|DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead. +k8s:conversion-gen=false +optional|
|template_automount_service_account_token|boolean|AutomountServiceAccountToken indicates whether a service account token should be automatically mounted. +optional|
|template_node_name|text|NodeName is a request to schedule this pod onto a specific node|
|template_host_network|boolean|Host networking requested for this pod|
|template_host_pid|boolean|Use the host's pid namespace. Optional: Default to false. +k8s:conversion-gen=false +optional|
|template_host_ipc|boolean|Use the host's ipc namespace. Optional: Default to false. +k8s:conversion-gen=false +optional|
|template_share_process_namespace|boolean|Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. +k8s:conversion-gen=false +optional|
|template_security_context|jsonb|SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty|
|template_hostname|text|Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value. +optional|
|template_subdomain|text|If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all. +optional|
|template_affinity|jsonb|If specified, the pod's scheduling constraints +optional|
|template_scheduler_name|text|If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler. +optional|
|template_priority_class_name|text|If specified, indicates the pod's priority|
|template_priority|integer|The priority value|
|template_dns_config|jsonb|Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. +optional|
|template_readiness_gates|jsonb|If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates +optional|
|template_runtime_class_name|text|RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod|
|template_enable_service_links|boolean|EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true. +optional|
|template_preemption_policy|text|PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate. +optional|
|template_overhead|jsonb|Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller|
|template_topology_spread_constraints|jsonb|TopologySpreadConstraints describes how a group of pods ought to spread across topology domains|
|template_set_hostname_as_fqdn|boolean|If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false. +optional|
|ttl_seconds_after_finished|integer|ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed)|
|completion_mode|text|CompletionMode specifies how Pod completions are tracked|
|suspend|boolean|Suspend specifies whether the Job controller should create Pods or not|
|status_active|integer|The number of actively running pods. +optional|
|status_succeeded|integer|The number of pods which reached phase Succeeded. +optional|
|status_failed|integer|The number of pods which reached phase Failed. +optional|
|status_completed_indexes|text|CompletedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format|
|status_uncounted_terminated_pods_succeeded|text[]|Succeeded holds UIDs of succeeded Pods. +listType=set +optional|
|status_uncounted_terminated_pods_failed|text[]|Failed holds UIDs of failed Pods. +listType=set +optional|
