
# Table: k8s_apps_daemon_set_template_owner_references
OwnerReference contains enough information to let you identify an owning object
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|daemon_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)|
|api_version|text|API version of the referent.|
|kind|text|Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds|
|name|text|Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names|
|uid|text|UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids|
|controller|boolean|If true, this reference points to the managing controller. +optional|
|block_owner_deletion|boolean|If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned. +optional|
