
# Table: k8s_apps_stateful_set_volume_claim_template_managed_fields
ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stateful_set_volume_claim_template_cq_id|uuid|Unique CloudQuery ID of k8s_apps_stateful_set_volume_claim_templates table (FK)|
|manager|text|Manager is an identifier of the workflow managing these fields.|
|operation|text|Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.|
|api_version|text|APIVersion defines the version of this resource that this field set applies to|
|fields_type|text|FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: "FieldsV1"|
|subresource|text|Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource|
