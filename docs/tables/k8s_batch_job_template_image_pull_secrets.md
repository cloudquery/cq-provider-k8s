
# Table: k8s_batch_job_template_image_pull_secrets
LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace. +structType=atomic
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|job_cq_id|uuid|Unique CloudQuery ID of k8s_batch_jobs table (FK)|
|name|text|Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields|
