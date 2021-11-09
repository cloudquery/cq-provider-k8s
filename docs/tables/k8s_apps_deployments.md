
# Table: k8s_apps_deployments
Deployment enables declarative updates for Pods and ReplicaSets.
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
|replicas|integer|Number of desired pods|
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
|strategy_type|text|Type of deployment|
|strategy_rolling_update_max_unavailable_type|bigint||
|strategy_rolling_update_max_unavailable_int_val|integer||
|strategy_rolling_update_max_unavailable_str_val|text||
|strategy_rolling_update_max_surge_type|bigint||
|strategy_rolling_update_max_surge_int_val|integer||
|strategy_rolling_update_max_surge_str_val|text||
|min_ready_seconds|integer|Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) +optional|
|revision_history_limit|integer|The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. +optional|
|paused|boolean|Indicates that the deployment is paused. +optional|
|progress_deadline_seconds|integer|The maximum time in seconds for a deployment to make progress before it is considered to be failed|
|status_observed_generation|bigint|The generation observed by the deployment controller. +optional|
|status_replicas|integer|Total number of non-terminated pods targeted by this deployment (their labels match the selector). +optional|
|status_updated_replicas|integer|Total number of non-terminated pods targeted by this deployment that have the desired template spec. +optional|
|status_ready_replicas|integer|Total number of ready pods targeted by this deployment. +optional|
|status_available_replicas|integer|Total number of available pods (ready for at least minReadySeconds) targeted by this deployment. +optional|
|status_unavailable_replicas|integer|Total number of unavailable pods targeted by this deployment|
|status_collision_count|integer|Count of hash collisions for the Deployment|
