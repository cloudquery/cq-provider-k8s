
# Table: k8s_batch_job_template_spec_init_containers
A single application container that you want to run within a pod.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|job_cq_id|uuid|Unique CloudQuery ID of k8s_batch_jobs table (FK)|
|name|text|Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.|
|image|text|Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets. +optional|
|command|text[]|Entrypoint array|
|args|text[]|Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment|
|working_dir|text|Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. +optional|
|env_from|jsonb|List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER|
|resources_limits|jsonb|Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional|
|resources_requests|jsonb|Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional|
|liveness_probe|jsonb|Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional|
|readiness_probe|jsonb|Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional|
|startup_probe|jsonb|StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional|
|lifecycle|jsonb|Actions that the management system should take in response to container lifecycle events. Cannot be updated. +optional|
|termination_message_path|text|Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes|
|termination_message_policy|text|Indicate how the termination message should be populated|
|image_pull_policy|text|Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images +optional|
|security_context|jsonb|SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ +optional|
|stdin|boolean|Whether this container should allocate a buffer for stdin in the container runtime|
|stdin_once|boolean|Whether the container runtime should close the stdin channel after it has been opened by a single attach|
|tty|boolean|Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false. +optional|
