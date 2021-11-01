
# Table: k8s_apps_stateful_set_volume_claim_templates
PersistentVolumeClaim is a user's request for and claim to a persistent volume
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stateful_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)|
|type_meta_kind|text||
|type_meta_api_version|text||
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
|access_modes|text[]|AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional|
|selector_match_labels|jsonb||
|selector_match_expressions|jsonb||
|resources_limits|jsonb|Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional|
|resources_requests|jsonb|Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional|
|volume_name|text|VolumeName is the binding reference to the PersistentVolume backing this claim. +optional|
|storage_class_name|text|Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 +optional|
|volume_mode|text|volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. +optional|
|data_source_api_group|text|APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. +optional|
|data_source_kind|text|Kind is the type of resource being referenced|
|data_source_name|text|Name is the name of resource being referenced|
|data_source_ref_api_group|text|APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. +optional|
|data_source_ref_kind|text|Kind is the type of resource being referenced|
|data_source_ref_name|text|Name is the name of resource being referenced|
|status_phase|text|Phase represents the current phase of PersistentVolumeClaim. +optional|
|status_access_modes|text[]|AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional|
|status_capacity|jsonb|Represents the actual resources of the underlying volume. +optional|
