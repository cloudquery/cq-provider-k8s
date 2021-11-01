
# Table: k8s_apps_stateful_sets
StatefulSet represents a set of pods with consistent identities. Identities are defined as:  - Network: A single stable DNS and hostname.  - Storage: As many VolumeClaims as requested. The StatefulSet guarantees that a given network identity will always map to the same storage identity.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|k8s_config_context|text|Name of the context from k8s configuration.|
|name|text||
|generate_name|text||
|namespace|text||
|self_link|text||
|uid|text||
|resource_version|text||
|generation|bigint||
|deletion_grace_period_seconds|bigint||
|labels|jsonb||
|annotations|jsonb||
|finalizers|text[]||
|cluster_name|text||
|replicas|integer|replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. TODO: Consider a rename of this field. +optional|
|selector_match_labels|jsonb||
|template_name|text||
|template_generate_name|text||
|template_namespace|text||
|template_self_link|text||
|template_uid|text||
|template_resource_version|text||
|template_generation|bigint||
|template_deletion_grace_period_seconds|bigint||
|template_labels|jsonb||
|template_annotations|jsonb||
|template_finalizers|text[]||
|template_cluster_name|text||
|template_restart_policy|text||
|template_termination_grace_period_seconds|bigint||
|template_active_deadline_seconds|bigint||
|template_dns_policy|text||
|template_node_selector|jsonb||
|template_service_account_name|text||
|template_deprecated_service_account|text||
|template_automount_service_account_token|boolean||
|template_node_name|text||
|template_host_network|boolean||
|template_host_pid|boolean||
|template_host_ipc|boolean||
|template_share_process_namespace|boolean||
|template_security_context|jsonb||
|template_hostname|text||
|template_subdomain|text||
|template_affinity|jsonb||
|template_scheduler_name|text||
|template_priority_class_name|text||
|template_priority|integer||
|template_dns_config|jsonb||
|template_readiness_gates|jsonb||
|template_runtime_class_name|text||
|template_enable_service_links|boolean||
|template_preemption_policy|text||
|template_overhead|jsonb||
|template_topology_spread_constraints|jsonb||
|template_set_hostname_as_fqdn|boolean||
|service_name|text|serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set|
|pod_management_policy|text|podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down|
|update_strategy_type|text|Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate. +optional|
|update_strategy_rolling_update_partition|integer|Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0. +optional|
|revision_history_limit|integer|revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history|
|min_ready_seconds|integer|Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate. +optional|
|status_observed_generation|bigint|observedGeneration is the most recent generation observed for this StatefulSet|
|status_replicas|integer|replicas is the number of Pods created by the StatefulSet controller.|
|status_ready_replicas|integer|readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.|
|status_current_replicas|integer|currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.|
|status_updated_replicas|integer|updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.|
|status_current_revision|text|currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).|
|status_update_revision|text|updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)|
|status_collision_count|integer|collisionCount is the count of hash collisions for the StatefulSet|
|status_available_replicas|integer|Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset. This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate. Remove omitempty when graduating to beta +optional|
