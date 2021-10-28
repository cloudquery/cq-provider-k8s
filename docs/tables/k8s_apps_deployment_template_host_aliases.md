
# Table: k8s_apps_deployment_template_host_aliases
HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|deployment_cq_id|uuid|Unique CloudQuery ID of k8s_apps_deployments table (FK)|
|ip|text|IP address of the host file entry.|
|hostnames|text[]|Hostnames for the above IP address.|
