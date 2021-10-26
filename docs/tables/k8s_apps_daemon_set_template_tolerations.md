
# Table: k8s_apps_daemon_set_template_tolerations
The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|daemon_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)|
|key|text|Key is the taint key that the toleration applies to|
|operator|text|Operator represents a key's relationship to the value. Valid operators are Exists and Equal|
|value|text|Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. +optional|
|effect|text|Effect indicates the taint effect to match|
|toleration_seconds|bigint|TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint|
