
# Table: k8s_batch_cron_jobs
CronJob represents the configuration of a single cron job.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|k8s_config_context|text|Name of the context from k8s configuration.|
|kind|text|Kind is a string value representing the REST resource this object represents.|
|api_version|text|Defines the versioned schema of this representation of an object.|
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
|schedule|text|The schedule in Cron format.|
|starting_deadline_seconds|bigint|Optional deadline in seconds for starting the job if it misses scheduled time for any reason|
|concurrency_policy|text|Specifies how to treat concurrent executions of a Job.|
|suspend|boolean|This flag tells the controller to suspend subsequent executions, it does not apply to already started executions|
|successful_jobs_history_limit|integer|The number of successful finished jobs to retain|
|failed_jobs_history_limit|integer|The number of failed finished jobs to retain|
