
# Table: k8s_batch_jobs
Job represents the configuration of a single job.
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
|parallelism|integer|Specifies the maximum desired number of pods the job should run at any given time|
|completions|integer|Specifies the desired number of successfully finished pods the job should be run with|
|active_deadline_seconds|bigint|Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer|
|backoff_limit|integer|Specifies the number of retries before marking this job failed. Defaults to 6 +optional|
|selector_match_labels|jsonb||
|manual_selector|boolean|manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template|
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
|ttl_seconds_after_finished|integer|ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed)|
|completion_mode|text|CompletionMode specifies how Pod completions are tracked|
|suspend|boolean|Suspend specifies whether the Job controller should create Pods or not|
|status_active|integer|The number of actively running pods. +optional|
|status_succeeded|integer|The number of pods which reached phase Succeeded. +optional|
|status_failed|integer|The number of pods which reached phase Failed. +optional|
|status_completed_indexes|text|CompletedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format|
|status_uncounted_terminated_pods_succeeded|text[]|Succeeded holds UIDs of succeeded Pods. +listType=set +optional|
|status_uncounted_terminated_pods_failed|text[]|Failed holds UIDs of failed Pods. +listType=set +optional|
