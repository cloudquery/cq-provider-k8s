
# Table: k8s_batch_cron_job_templates
JobTemplateSpec describes the data a Job should have when created from a template
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cron_job_cq_id|uuid|Unique CloudQuery ID of k8s_batch_cron_jobs table (FK)|
|name|text|Unique name within a namespace.|
|namespace|text|Namespace defines the space within which each name must be unique.|
|uid|text|UID is the unique in time and space value for this object.|
|resource_version|text|An opaque value that represents the internal version of this object.|
|generation|bigint|A sequence number representing a specific generation of the desired state.|
|deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate.|
|labels|jsonb|Map of string keys and values that can be used to organize and categorize objects.|
|annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools.|
|owner_references|jsonb|List of objects depended by this object.|
|finalizers|text[]|List of finalizers|
|cluster_name|text|The name of the cluster which the object belongs to.|
|parallelism|integer|Specifies the maximum desired number of pods the job should run at any given time|
|completions|integer|Specifies the desired number of successfully finished pods the job should be run with|
|active_deadline_seconds|bigint|Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it.|
|backoff_limit|integer|Specifies the number of retries before marking this job failed.|
|selector|jsonb|A label query over pods that should match the pod count.|
|manual_selector|boolean|manualSelector controls generation of pod labels and pod selectors.|
|ttl_seconds_after_finished|integer|Limits the lifetime of a Job that has finished execution.|
|completion_mode|text|CompletionMode specifies how Pod completions are tracked|
|suspend|boolean|Suspend specifies whether the Job controller should create Pods or not|
