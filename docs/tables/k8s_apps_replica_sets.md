
# Table: k8s_apps_replica_sets
ReplicaSet ensures that a specified number of pod replicas are running at any given time.
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
|owner_references|jsonb||
|finalizers|text[]||
|cluster_name|text||
|managed_fields|jsonb||
|replicas|integer|Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller +optional|
|min_ready_seconds|integer|Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) +optional|
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
|status_replicas|integer|Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller|
|status_fully_labeled_replicas|integer|The number of pods that have labels matching the labels of the pod template of the replicaset. +optional|
|status_ready_replicas|integer|The number of ready replicas for this replica set. +optional|
|status_available_replicas|integer|The number of available replicas (ready for at least minReadySeconds) for this replica set. +optional|
|status_observed_generation|bigint|ObservedGeneration reflects the generation of the most recently observed ReplicaSet. +optional|
