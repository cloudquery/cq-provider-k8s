
# Table: k8s_apps_daemon_sets
DaemonSet represents the configuration of a daemon set.
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
|update_strategy_type|text|Type of daemon set update|
|update_strategy_rolling_update_max_unavailable_type|bigint||
|update_strategy_rolling_update_max_unavailable_int_val|integer||
|update_strategy_rolling_update_max_unavailable_str_val|text||
|update_strategy_rolling_update_max_surge_type|bigint||
|update_strategy_rolling_update_max_surge_int_val|integer||
|update_strategy_rolling_update_max_surge_str_val|text||
|min_ready_seconds|integer|The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available|
|revision_history_limit|integer|The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. +optional|
|status_current_number_scheduled|integer|The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/|
|status_number_misscheduled|integer|The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/|
|status_desired_number_scheduled|integer|The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/|
|status_number_ready|integer|The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.|
|status_observed_generation|bigint|The most recent generation observed by the daemon set controller. +optional|
|status_updated_number_scheduled|integer|The total number of nodes that are running updated daemon pod +optional|
|status_number_available|integer|The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds) +optional|
|status_number_unavailable|integer|The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds) +optional|
|status_collision_count|integer|Count of hash collisions for the DaemonSet|
