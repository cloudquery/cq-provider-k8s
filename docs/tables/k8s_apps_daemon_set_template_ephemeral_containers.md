
# Table: k8s_apps_daemon_set_template_ephemeral_containers
An EphemeralContainer is a container that may be added temporarily to an existing pod for user-initiated activities such as debugging
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|daemon_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)|
|name|text|Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.|
|image|text|Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images|
|command|text[]|Entrypoint array|
|args|text[]|Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment|
|working_dir|text|Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. +optional|
|env_from|jsonb|List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER|
|resources_limits|jsonb|Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional|
|resources_requests|jsonb|Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional|
|liveness_probe|jsonb|Probes are not allowed for ephemeral containers. +optional|
|readiness_probe|jsonb|Probes are not allowed for ephemeral containers. +optional|
|startup_probe|jsonb|Probes are not allowed for ephemeral containers. +optional|
|lifecycle|jsonb|Lifecycle is not allowed for ephemeral containers. +optional|
|termination_message_path|text|Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes|
|termination_message_policy|text|Indicate how the termination message should be populated|
|image_pull_policy|text|Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images +optional|
|security_context|jsonb|Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. +optional|
|stdin|boolean|Whether this container should allocate a buffer for stdin in the container runtime|
|stdin_once|boolean|Whether the container runtime should close the stdin channel after it has been opened by a single attach|
|tty|boolean|Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false. +optional|
|target_container_name|text|If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container is run in whatever namespaces are shared for the pod|
