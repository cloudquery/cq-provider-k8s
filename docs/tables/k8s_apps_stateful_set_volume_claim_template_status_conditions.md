
# Table: k8s_apps_stateful_set_volume_claim_template_status_conditions
PersistentVolumeClaimCondition contails details about state of pvc
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stateful_set_volume_claim_template_cq_id|uuid|Unique CloudQuery ID of k8s_apps_stateful_set_volume_claim_templates table (FK)|
|type|text||
|status|text||
|reason|text|Unique, this should be a short, machine understandable string that gives the reason for condition's last transition|
|message|text|Human-readable message indicating details about last transition. +optional|
