package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AppsDeployments() *schema.Table {
	return &schema.Table{
		Name:         "k8s_apps_deployments",
		Description:  "Deployment enables declarative updates for Pods and ReplicaSets.",
		Resolver:     fetchAppsDeployments,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			{
				Name:        "name",
				Description: "Name must be unique within a namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:        "generate_name",
				Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.GenerateName"),
			},
			{
				Name:        "namespace",
				Description: "Namespace defines the space within which each name must be unique",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:        "self_link",
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.SelfLink"),
			},
			{
				Name:        "uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:        "resource_version",
				Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "generation",
				Description: "A sequence number representing a specific generation of the desired state. Populated by the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:        "deletion_grace_period_seconds",
				Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:        "labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:        "annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:        "owner_references",
				Description: "List of objects depended by this object",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsOwnerReferences,
			},
			{
				Name:        "finalizers",
				Description: "Must be empty before the object is deleted from the registry",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "managed_fields",
				Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsManagedFields,
			},
			{
				Name:        "replicas",
				Description: "Number of desired pods",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Replicas"),
			},
			{
				Name:        "selector_match_labels",
				Description: "matchLabels is a map of {key,value} pairs",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:        "template_name",
				Description: "Name must be unique within a namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.Name"),
			},
			{
				Name:        "template_generate_name",
				Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.GenerateName"),
			},
			{
				Name:        "template_namespace",
				Description: "Namespace defines the space within which each name must be unique",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.Namespace"),
			},
			{
				Name:        "template_self_link",
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.SelfLink"),
			},
			{
				Name:        "template_uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.UID"),
			},
			{
				Name:        "template_resource_version",
				Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "template_generation",
				Description: "A sequence number representing a specific generation of the desired state. Populated by the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.Generation"),
			},
			{
				Name:        "template_deletion_grace_period_seconds",
				Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:        "template_labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.Labels"),
			},
			{
				Name:        "template_annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.Annotations"),
			},
			{
				Name:        "template_finalizers",
				Description: "Must be empty before the object is deleted from the registry",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.Finalizers"),
			},
			{
				Name:        "template_cluster_name",
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.ObjectMeta.ClusterName"),
			},
			{
				Name:        "template_spec_restart_policy",
				Description: "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.RestartPolicy"),
			},
			{
				Name:        "template_spec_termination_grace_period_seconds",
				Description: "Optional duration in seconds the pod needs to terminate gracefully",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.Spec.TerminationGracePeriodSeconds"),
			},
			{
				Name:        "template_spec_active_deadline_seconds",
				Description: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer. +optional",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.Spec.ActiveDeadlineSeconds"),
			},
			{
				Name:        "template_spec_dns_policy",
				Description: "Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.DNSPolicy"),
			},
			{
				Name:        "template_spec_node_selector",
				Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ +optional +mapType=atomic",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Template.Spec.NodeSelector"),
			},
			{
				Name:        "template_spec_service_account_name",
				Description: "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/ +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.ServiceAccountName"),
			},
			{
				Name:        "template_spec_deprecated_service_account",
				Description: "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead. +k8s:conversion-gen=false +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.DeprecatedServiceAccount"),
			},
			{
				Name:        "template_spec_automount_service_account_token",
				Description: "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted. +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.AutomountServiceAccountToken"),
			},
			{
				Name:        "template_spec_node_name",
				Description: "NodeName is a request to schedule this pod onto a specific node",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.NodeName"),
			},
			{
				Name:        "template_spec_host_network",
				Description: "Host networking requested for this pod",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.HostNetwork"),
			},
			{
				Name:        "template_spec_host_pid",
				Description: "Use the host's pid namespace. Optional: Default to false. +k8s:conversion-gen=false +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.HostPID"),
			},
			{
				Name:        "template_spec_host_ipc",
				Description: "Use the host's ipc namespace. Optional: Default to false. +k8s:conversion-gen=false +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.HostIPC"),
			},
			{
				Name:        "template_spec_share_process_namespace",
				Description: "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. +k8s:conversion-gen=false +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.ShareProcessNamespace"),
			},
			{
				Name:        "template_spec_security_context",
				Description: "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplateSpecSecurityContext,
			},
			{
				Name:        "template_spec_hostname",
				Description: "Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.Hostname"),
			},
			{
				Name:        "template_spec_subdomain",
				Description: "If specified, the fully qualified Pod hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\". If not specified, the pod will not have a domainname at all. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.Subdomain"),
			},
			{
				Name:        "template_spec_affinity",
				Description: "If specified, the pod's scheduling constraints +optional",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplateSpecAffinity,
			},
			{
				Name:        "template_spec_scheduler_name",
				Description: "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.SchedulerName"),
			},
			{
				Name:        "template_spec_priority_class_name",
				Description: "If specified, indicates the pod's priority",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.PriorityClassName"),
			},
			{
				Name:        "template_spec_priority",
				Description: "The priority value",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Template.Spec.Priority"),
			},
			{
				Name:        "template_spec_dns_config",
				Description: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. +optional",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplateSpecDnsConfig,
			},
			{
				Name:        "template_spec_readiness_gates",
				Description: "If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to \"True\" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates +optional",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplateSpecReadinessGates,
			},
			{
				Name:        "template_spec_runtime_class_name",
				Description: "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.RuntimeClassName"),
			},
			{
				Name:        "template_spec_enable_service_links",
				Description: "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true. +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.EnableServiceLinks"),
			},
			{
				Name:        "template_spec_preemption_policy",
				Description: "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.PreemptionPolicy"),
			},
			{
				Name:        "template_spec_overhead",
				Description: "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplateSpecOverhead,
			},
			{
				Name:        "template_spec_topology_spread_constraints",
				Description: "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplateSpecTopologySpreadConstraints,
			},
			{
				Name:        "template_spec_set_hostname_as_fqdn",
				Description: "If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\\\SYSTEM\\\\CurrentControlSet\\\\Services\\\\Tcpip\\\\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false. +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Template.Spec.SetHostnameAsFQDN"),
			},
			{
				Name:        "strategy_type",
				Description: "Type of deployment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Strategy.Type"),
			},
			{
				Name:     "strategy_rolling_update_max_unavailable_type",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxUnavailable.Type"),
			},
			{
				Name:     "strategy_rolling_update_max_unavailable_int_val",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxUnavailable.IntVal"),
			},
			{
				Name:     "strategy_rolling_update_max_unavailable_str_val",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal"),
			},
			{
				Name:     "strategy_rolling_update_max_surge_type",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxSurge.Type"),
			},
			{
				Name:     "strategy_rolling_update_max_surge_int_val",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxSurge.IntVal"),
			},
			{
				Name:     "strategy_rolling_update_max_surge_str_val",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxSurge.StrVal"),
			},
			{
				Name:        "min_ready_seconds",
				Description: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:        "revision_history_limit",
				Description: "The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.RevisionHistoryLimit"),
			},
			{
				Name:        "paused",
				Description: "Indicates that the deployment is paused. +optional",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Paused"),
			},
			{
				Name:        "progress_deadline_seconds",
				Description: "The maximum time in seconds for a deployment to make progress before it is considered to be failed",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.ProgressDeadlineSeconds"),
			},
			{
				Name:        "status_observed_generation",
				Description: "The generation observed by the deployment controller. +optional",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_replicas",
				Description: "Total number of non-terminated pods targeted by this deployment (their labels match the selector). +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Replicas"),
			},
			{
				Name:        "status_updated_replicas",
				Description: "Total number of non-terminated pods targeted by this deployment that have the desired template spec. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UpdatedReplicas"),
			},
			{
				Name:        "status_ready_replicas",
				Description: "Total number of ready pods targeted by this deployment. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.ReadyReplicas"),
			},
			{
				Name:        "status_available_replicas",
				Description: "Total number of available pods (ready for at least minReadySeconds) targeted by this deployment. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.AvailableReplicas"),
			},
			{
				Name:        "status_unavailable_replicas",
				Description: "Total number of unavailable pods targeted by this deployment",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UnavailableReplicas"),
			},
			{
				Name:        "status_collision_count",
				Description: "Count of hash collisions for the Deployment",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CollisionCount"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_apps_deployment_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchAppsDeploymentSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "values is an array of string values",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchAppsDeploymentTemplateOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_version",
						Description: "API version of the referent.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("APIVersion"),
					},
					{
						Name:        "kind",
						Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
						Type:        schema.TypeString,
					},
					{
						Name:        "uid",
						Description: "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UID"),
					},
					{
						Name:        "controller",
						Description: "If true, this reference points to the managing controller. +optional",
						Type:        schema.TypeBool,
					},
					{
						Name:        "block_owner_deletion",
						Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned. +optional",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchAppsDeploymentTemplateManagedFields,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "manager",
						Description: "Manager is an identifier of the workflow managing these fields.",
						Type:        schema.TypeString,
					},
					{
						Name:        "operation",
						Description: "Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_version",
						Description: "APIVersion defines the version of this resource that this field set applies to",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("APIVersion"),
					},
					{
						Name:        "fields_type",
						Description: "FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"",
						Type:        schema.TypeString,
					},
					{
						Name:        "subresource",
						Description: "Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_volumes",
				Description: "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
				Resolver:    fetchAppsDeploymentTemplateSpecVolumes,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						Type:        schema.TypeString,
					},
					{
						Name:        "volume_source_host_path",
						Description: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.HostPath.Path"),
					},
					{
						Name:        "volume_source_host_path_type",
						Description: "Type for HostPath Volume Defaults to \"\" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.HostPath.Type"),
					},
					{
						Name:        "volume_source_empty_dir_medium",
						Description: "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.EmptyDir.Medium"),
					},
					{
						Name:     "volume_source_empty_dir_size_limit_format",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VolumeSource.EmptyDir.SizeLimit.Format"),
					},
					{
						Name:        "volume_source_g_c_e_persistent_disk_p_d_name",
						Description: "Unique name of the PD resource in GCE",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.GCEPersistentDisk.PDName"),
					},
					{
						Name:        "volume_source_g_c_e_persistent_disk_f_s_type",
						Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.GCEPersistentDisk.FSType"),
					},
					{
						Name:        "volume_source_g_c_e_persistent_disk_partition",
						Description: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.GCEPersistentDisk.Partition"),
					},
					{
						Name:        "volume_source_g_c_e_persistent_disk_read_only",
						Description: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.GCEPersistentDisk.ReadOnly"),
					},
					{
						Name:        "volume_source_aws_elastic_block_store_volume_id",
						Description: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AWSElasticBlockStore.VolumeID"),
					},
					{
						Name:        "volume_source_aws_elastic_block_store_f_s_type",
						Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AWSElasticBlockStore.FSType"),
					},
					{
						Name:        "volume_source_aws_elastic_block_store_partition",
						Description: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.AWSElasticBlockStore.Partition"),
					},
					{
						Name:        "volume_source_aws_elastic_block_store_read_only",
						Description: "Specify \"true\" to force and set the ReadOnly property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.AWSElasticBlockStore.ReadOnly"),
					},
					{
						Name:        "volume_source_git_repo_repository",
						Description: "Repository URL",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.GitRepo.Repository"),
					},
					{
						Name:        "volume_source_git_repo_revision",
						Description: "Commit hash for the specified revision. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.GitRepo.Revision"),
					},
					{
						Name:        "volume_source_git_repo_directory",
						Description: "Target directory name. Must not contain or start with '..'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.GitRepo.Directory"),
					},
					{
						Name:        "volume_source_secret_secret_name",
						Description: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Secret.SecretName"),
					},
					{
						Name:        "volume_source_secret_default_mode",
						Description: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.Secret.DefaultMode"),
					},
					{
						Name:        "volume_source_secret_optional",
						Description: "Specify whether the Secret or its keys must be defined +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.Secret.Optional"),
					},
					{
						Name:        "volume_source_nfs_server",
						Description: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.NFS.Server"),
					},
					{
						Name:        "volume_source_nfs_path",
						Description: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.NFS.Path"),
					},
					{
						Name:        "volume_source_nfs_read_only",
						Description: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.NFS.ReadOnly"),
					},
					{
						Name:        "volume_source_iscsi_target_portal",
						Description: "iSCSI Target Portal",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.TargetPortal"),
					},
					{
						Name:        "volume_source_iscsi_i_q_n",
						Description: "Target iSCSI Qualified Name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.IQN"),
					},
					{
						Name:        "volume_source_iscsi_lun",
						Description: "iSCSI Target Lun number.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.Lun"),
					},
					{
						Name:        "volume_source_iscsi_iscsi_interface",
						Description: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp). +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.ISCSIInterface"),
					},
					{
						Name:        "volume_source_iscsi_f_s_type",
						Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.FSType"),
					},
					{
						Name:        "volume_source_iscsi_read_only",
						Description: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.ReadOnly"),
					},
					{
						Name:        "volume_source_iscsi_portals",
						Description: "iSCSI Target Portal List",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.Portals"),
					},
					{
						Name:        "volume_source_iscsi_discovery_c_h_a_p_auth",
						Description: "whether support iSCSI Discovery CHAP authentication +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.DiscoveryCHAPAuth"),
					},
					{
						Name:        "volume_source_iscsi_session_c_h_a_p_auth",
						Description: "whether support iSCSI Session CHAP authentication +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.SessionCHAPAuth"),
					},
					{
						Name:        "volume_source_iscsi_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.SecretRef.Name"),
					},
					{
						Name:        "volume_source_iscsi_initiator_name",
						Description: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ISCSI.InitiatorName"),
					},
					{
						Name:        "volume_source_glusterfs_endpoints_name",
						Description: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Glusterfs.EndpointsName"),
					},
					{
						Name:        "volume_source_glusterfs_path",
						Description: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Glusterfs.Path"),
					},
					{
						Name:        "volume_source_glusterfs_read_only",
						Description: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.Glusterfs.ReadOnly"),
					},
					{
						Name:        "volume_source_persistent_volume_claim_claim_name",
						Description: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.PersistentVolumeClaim.ClaimName"),
					},
					{
						Name:        "volume_source_persistent_volume_claim_read_only",
						Description: "Will force the ReadOnly setting in VolumeMounts. Default false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.PersistentVolumeClaim.ReadOnly"),
					},
					{
						Name:        "volume_source_rbd_ceph_monitors",
						Description: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.RBD.CephMonitors"),
					},
					{
						Name:        "volume_source_rbd_rbd_image",
						Description: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.RBD.RBDImage"),
					},
					{
						Name:        "volume_source_rbd_f_s_type",
						Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.RBD.FSType"),
					},
					{
						Name:        "volume_source_rbd_rbd_pool",
						Description: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.RBD.RBDPool"),
					},
					{
						Name:        "volume_source_rbd_rados_user",
						Description: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.RBD.RadosUser"),
					},
					{
						Name:        "volume_source_rbd_keyring",
						Description: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.RBD.Keyring"),
					},
					{
						Name:        "volume_source_rbd_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.RBD.SecretRef.Name"),
					},
					{
						Name:        "volume_source_rbd_read_only",
						Description: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.RBD.ReadOnly"),
					},
					{
						Name:        "volume_source_flex_volume_driver",
						Description: "Driver is the name of the driver to use for this volume.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.FlexVolume.Driver"),
					},
					{
						Name:        "volume_source_flex_volume_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.FlexVolume.FSType"),
					},
					{
						Name:        "volume_source_flex_volume_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.FlexVolume.SecretRef.Name"),
					},
					{
						Name:        "volume_source_flex_volume_read_only",
						Description: "Optional: Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.FlexVolume.ReadOnly"),
					},
					{
						Name:        "volume_source_flex_volume_options",
						Description: "Optional: Extra command options if any. +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.FlexVolume.Options"),
					},
					{
						Name:        "volume_source_cinder_volume_id",
						Description: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Cinder.VolumeID"),
					},
					{
						Name:        "volume_source_cinder_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Cinder.FSType"),
					},
					{
						Name:        "volume_source_cinder_read_only",
						Description: "Optional: Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.Cinder.ReadOnly"),
					},
					{
						Name:        "volume_source_cinder_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Cinder.SecretRef.Name"),
					},
					{
						Name:        "volume_source_ceph_f_s_monitors",
						Description: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.CephFS.Monitors"),
					},
					{
						Name:        "volume_source_ceph_f_s_path",
						Description: "Optional: Used as the mounted root, rather than the full Ceph tree, default is / +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CephFS.Path"),
					},
					{
						Name:        "volume_source_ceph_f_s_user",
						Description: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CephFS.User"),
					},
					{
						Name:        "volume_source_ceph_f_s_secret_file",
						Description: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CephFS.SecretFile"),
					},
					{
						Name:        "volume_source_ceph_f_s_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CephFS.SecretRef.Name"),
					},
					{
						Name:        "volume_source_ceph_f_s_read_only",
						Description: "Optional: Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.CephFS.ReadOnly"),
					},
					{
						Name:        "volume_source_flocker_dataset_name",
						Description: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Flocker.DatasetName"),
					},
					{
						Name:        "volume_source_flocker_dataset_uuid",
						Description: "UUID of the dataset",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Flocker.DatasetUUID"),
					},
					{
						Name:        "volume_source_downward_api_default_mode",
						Description: "Optional: mode bits to use on created files by default",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.DownwardAPI.DefaultMode"),
					},
					{
						Name:        "volume_source_f_c_target_w_w_ns",
						Description: "Optional: FC target worldwide names (WWNs) +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.FC.TargetWWNs"),
					},
					{
						Name:        "volume_source_f_c_lun",
						Description: "Optional: FC target lun number +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.FC.Lun"),
					},
					{
						Name:        "volume_source_f_c_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.FC.FSType"),
					},
					{
						Name:        "volume_source_f_c_read_only",
						Description: "Optional: Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.FC.ReadOnly"),
					},
					{
						Name:        "volume_source_f_c_w_w_ids",
						Description: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously. +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.FC.WWIDs"),
					},
					{
						Name:        "volume_source_azure_file_secret_name",
						Description: "the name of secret that contains Azure Storage Account Name and Key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureFile.SecretName"),
					},
					{
						Name:        "volume_source_azure_file_share_name",
						Description: "Share Name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureFile.ShareName"),
					},
					{
						Name:        "volume_source_azure_file_read_only",
						Description: "Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.AzureFile.ReadOnly"),
					},
					{
						Name:        "volume_source_config_map_local_object_reference_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ConfigMap.LocalObjectReference.Name"),
					},
					{
						Name:        "volume_source_config_map_default_mode",
						Description: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.ConfigMap.DefaultMode"),
					},
					{
						Name:        "volume_source_config_map_optional",
						Description: "Specify whether the ConfigMap or its keys must be defined +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.ConfigMap.Optional"),
					},
					{
						Name:        "volume_source_vsphere_volume_volume_path",
						Description: "Path that identifies vSphere volume vmdk",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.VsphereVolume.VolumePath"),
					},
					{
						Name:        "volume_source_vsphere_volume_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.VsphereVolume.FSType"),
					},
					{
						Name:        "volume_source_vsphere_volume_storage_policy_name",
						Description: "Storage Policy Based Management (SPBM) profile name. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.VsphereVolume.StoragePolicyName"),
					},
					{
						Name:        "volume_source_vsphere_volume_storage_policy_id",
						Description: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.VsphereVolume.StoragePolicyID"),
					},
					{
						Name:        "volume_source_quobyte_registry",
						Description: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Quobyte.Registry"),
					},
					{
						Name:        "volume_source_quobyte_volume",
						Description: "Volume is a string that references an already created Quobyte volume by name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Quobyte.Volume"),
					},
					{
						Name:        "volume_source_quobyte_read_only",
						Description: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.Quobyte.ReadOnly"),
					},
					{
						Name:        "volume_source_quobyte_user",
						Description: "User to map volume access to Defaults to serivceaccount user +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Quobyte.User"),
					},
					{
						Name:        "volume_source_quobyte_group",
						Description: "Group to map volume access to Default is no group +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Quobyte.Group"),
					},
					{
						Name:        "volume_source_quobyte_tenant",
						Description: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Quobyte.Tenant"),
					},
					{
						Name:        "volume_source_azure_disk_disk_name",
						Description: "The Name of the data disk in the blob storage",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureDisk.DiskName"),
					},
					{
						Name:        "volume_source_azure_disk_data_disk_uri",
						Description: "The URI the data disk in the blob storage",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureDisk.DataDiskURI"),
					},
					{
						Name:        "volume_source_azure_disk_caching_mode",
						Description: "Host Caching mode: None, Read Only, Read Write. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureDisk.CachingMode"),
					},
					{
						Name:        "volume_source_azure_disk_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureDisk.FSType"),
					},
					{
						Name:        "volume_source_azure_disk_read_only",
						Description: "Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.AzureDisk.ReadOnly"),
					},
					{
						Name:        "volume_source_azure_disk_kind",
						Description: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.AzureDisk.Kind"),
					},
					{
						Name:        "volume_source_photon_persistent_disk_pd_id",
						Description: "ID that identifies Photon Controller persistent disk",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.PhotonPersistentDisk.PdID"),
					},
					{
						Name:        "volume_source_photon_persistent_disk_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.PhotonPersistentDisk.FSType"),
					},
					{
						Name:        "volume_source_projected_default_mode",
						Description: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSource.Projected.DefaultMode"),
					},
					{
						Name:        "volume_source_portworx_volume_volume_id",
						Description: "VolumeID uniquely identifies a Portworx volume",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.PortworxVolume.VolumeID"),
					},
					{
						Name:        "volume_source_portworx_volume_f_s_type",
						Description: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.PortworxVolume.FSType"),
					},
					{
						Name:        "volume_source_portworx_volume_read_only",
						Description: "Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.PortworxVolume.ReadOnly"),
					},
					{
						Name:        "volume_source_scale_i_o_gateway",
						Description: "The host address of the ScaleIO API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.Gateway"),
					},
					{
						Name:        "volume_source_scale_i_o_system",
						Description: "The name of the storage system as configured in ScaleIO.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.System"),
					},
					{
						Name:        "volume_source_scale_i_os_ecret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.SecretRef.Name"),
					},
					{
						Name:        "volume_source_scale_i_o_ssl_enabled",
						Description: "Flag to enable/disable SSL communication with Gateway, default false +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.SSLEnabled"),
					},
					{
						Name:        "volume_source_scale_i_o_protection_domain",
						Description: "The name of the ScaleIO Protection Domain for the configured storage. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.ProtectionDomain"),
					},
					{
						Name:        "volume_source_scale_i_o_storage_pool",
						Description: "The ScaleIO Storage Pool associated with the protection domain. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.StoragePool"),
					},
					{
						Name:        "volume_source_scale_i_o_storage_mode",
						Description: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.StorageMode"),
					},
					{
						Name:        "volume_source_scale_i_o_volume_name",
						Description: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.VolumeName"),
					},
					{
						Name:        "volume_source_scale_i_o_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.FSType"),
					},
					{
						Name:        "volume_source_scale_i_o_read_only",
						Description: "Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO.ReadOnly"),
					},
					{
						Name:        "volume_source_storage_os_volume_name",
						Description: "VolumeName is the human-readable name of the StorageOS volume",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.StorageOS.VolumeName"),
					},
					{
						Name:        "volume_source_storage_os_volume_namespace",
						Description: "VolumeNamespace specifies the scope of the volume within StorageOS",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.StorageOS.VolumeNamespace"),
					},
					{
						Name:        "volume_source_storage_os_f_s_type",
						Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.StorageOS.FSType"),
					},
					{
						Name:        "volume_source_storage_os_read_only",
						Description: "Defaults to false (read/write)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.StorageOS.ReadOnly"),
					},
					{
						Name:        "volume_source_storage_os_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.StorageOS.SecretRef.Name"),
					},
					{
						Name:        "volume_source_csi_driver",
						Description: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CSI.Driver"),
					},
					{
						Name:        "volume_source_csi_read_only",
						Description: "Specifies a read-only configuration for the volume. Defaults to false (read/write). +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VolumeSource.CSI.ReadOnly"),
					},
					{
						Name:        "volume_source_csi_f_s_type",
						Description: "Filesystem type to mount",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CSI.FSType"),
					},
					{
						Name:        "volume_source_csi_volume_attributes",
						Description: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.CSI.VolumeAttributes"),
					},
					{
						Name:        "volume_source_csi_node_publish_secret_ref_name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.CSI.NodePublishSecretRef.Name"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_name",
						Description: "Name must be unique within a namespace",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.Name"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_generate_name",
						Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.GenerateName"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_namespace",
						Description: "Namespace defines the space within which each name must be unique",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.Namespace"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_self_link",
						Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.SelfLink"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_uid",
						Description: "UID is the unique in time and space value for this object",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.UID"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_resource_version",
						Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.ResourceVersion"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_generation",
						Description: "A sequence number representing a specific generation of the desired state. Populated by the system",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.Generation"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_deletion_grace_period_seconds",
						Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.DeletionGracePeriodSeconds"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_labels",
						Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.Labels"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_annotations",
						Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.Annotations"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_finalizers",
						Description: "Must be empty before the object is deleted from the registry",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.Finalizers"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_object_meta_cluster_name",
						Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.ObjectMeta.ClusterName"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_access_modes",
						Description: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.AccessModes"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_selector_match_labels",
						Description: "matchLabels is a map of {key,value} pairs",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Selector.MatchLabels"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_resources_limits",
						Description: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Limits"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_resources_requests",
						Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Requests"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_volume_name",
						Description: "VolumeName is the binding reference to the PersistentVolume backing this claim. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.VolumeName"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_storage_class_name",
						Description: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.StorageClassName"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_volume_mode",
						Description: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.VolumeMode"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_data_source_api_group",
						Description: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.DataSource.APIGroup"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_data_source_kind",
						Description: "Kind is the type of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.DataSource.Kind"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_data_source_name",
						Description: "Name is the name of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.DataSource.Name"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_data_source_ref_api_group",
						Description: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.DataSourceRef.APIGroup"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_data_source_ref_kind",
						Description: "Kind is the type of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.DataSourceRef.Kind"),
					},
					{
						Name:        "volume_source_ephemeral_volume_claim_template_spec_data_source_ref_name",
						Description: "Name is the name of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.DataSourceRef.Name"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_secret_items",
						Description: "Maps a string key to a path within a volume.",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceSecretItems,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key",
								Description: "The key to project.",
								Type:        schema.TypeString,
							},
							{
								Name:        "path",
								Description: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
								Type:        schema.TypeString,
							},
							{
								Name:        "mode",
								Description: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
								Type:        schema.TypeInt,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_downward_api_items",
						Description: "DownwardAPIVolumeFile represents information to create the file containing the pod field",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceDownwardApiItems,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "path",
								Description: "Required: Path is  the relative path name of the file to be created",
								Type:        schema.TypeString,
							},
							{
								Name:        "field_ref_api_version",
								Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\". +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("FieldRef.APIVersion"),
							},
							{
								Name:        "field_ref_field_path",
								Description: "Path of the field to select in the specified API version.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("FieldRef.FieldPath"),
							},
							{
								Name:        "resource_field_ref_container_name",
								Description: "Container name: required for volumes, optional for env vars +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ResourceFieldRef.ContainerName"),
							},
							{
								Name:        "resource_field_ref_resource",
								Description: "Required: resource to select",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ResourceFieldRef.Resource"),
							},
							{
								Name:     "resource_field_ref_divisor_format",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ResourceFieldRef.Divisor.Format"),
							},
							{
								Name:        "mode",
								Description: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
								Type:        schema.TypeInt,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_config_map_items",
						Description: "Maps a string key to a path within a volume.",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceConfigMapItems,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key",
								Description: "The key to project.",
								Type:        schema.TypeString,
							},
							{
								Name:        "path",
								Description: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
								Type:        schema.TypeString,
							},
							{
								Name:        "mode",
								Description: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
								Type:        schema.TypeInt,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_projected_sources",
						Description: "Projection that may be projected along with other supported volume types",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSources,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "secret_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Secret.LocalObjectReference.Name"),
							},
							{
								Name:        "secret_optional",
								Description: "Specify whether the Secret or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Secret.Optional"),
							},
							{
								Name:        "config_map_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ConfigMap.LocalObjectReference.Name"),
							},
							{
								Name:        "config_map_optional",
								Description: "Specify whether the ConfigMap or its keys must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ConfigMap.Optional"),
							},
							{
								Name:        "service_account_token_audience",
								Description: "Audience is the intended audience of the token",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ServiceAccountToken.Audience"),
							},
							{
								Name:        "service_account_token_expiration_seconds",
								Description: "ExpirationSeconds is the requested duration of validity of the service account token",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("ServiceAccountToken.ExpirationSeconds"),
							},
							{
								Name:        "service_account_token_path",
								Description: "Path is the path relative to the mount point of the file to project the token into.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ServiceAccountToken.Path"),
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "k8s_apps_deployment_template_spec_volume_volume_source_projected_source_secret_items",
								Description: "Maps a string key to a path within a volume.",
								Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSourceSecretItems,
								Columns: []schema.Column{
									{
										Name:        "deployment_template_spec_volume_volume_source_projected_source_cq_id",
										Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volume_volume_source_projected_sources table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "key",
										Description: "The key to project.",
										Type:        schema.TypeString,
									},
									{
										Name:        "path",
										Description: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
										Type:        schema.TypeString,
									},
									{
										Name:        "mode",
										Description: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
										Type:        schema.TypeInt,
									},
								},
							},
							{
								Name:        "k8s_apps_deployment_template_spec_volume_volume_source_projected_source_downward_api_items",
								Description: "DownwardAPIVolumeFile represents information to create the file containing the pod field",
								Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSourceDownwardApiItems,
								Columns: []schema.Column{
									{
										Name:        "deployment_template_spec_volume_volume_source_projected_source_cq_id",
										Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volume_volume_source_projected_sources table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "path",
										Description: "Required: Path is  the relative path name of the file to be created",
										Type:        schema.TypeString,
									},
									{
										Name:        "field_ref_api_version",
										Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\". +optional",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("FieldRef.APIVersion"),
									},
									{
										Name:        "field_ref_field_path",
										Description: "Path of the field to select in the specified API version.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("FieldRef.FieldPath"),
									},
									{
										Name:        "resource_field_ref_container_name",
										Description: "Container name: required for volumes, optional for env vars +optional",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("ResourceFieldRef.ContainerName"),
									},
									{
										Name:        "resource_field_ref_resource",
										Description: "Required: resource to select",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("ResourceFieldRef.Resource"),
									},
									{
										Name:     "resource_field_ref_divisor_format",
										Type:     schema.TypeString,
										Resolver: schema.PathResolver("ResourceFieldRef.Divisor.Format"),
									},
									{
										Name:        "mode",
										Description: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
										Type:        schema.TypeInt,
									},
								},
							},
							{
								Name:        "k8s_apps_deployment_template_spec_volume_volume_source_projected_source_config_map_items",
								Description: "Maps a string key to a path within a volume.",
								Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSourceConfigMapItems,
								Columns: []schema.Column{
									{
										Name:        "deployment_template_spec_volume_volume_source_projected_source_cq_id",
										Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volume_volume_source_projected_sources table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "key",
										Description: "The key to project.",
										Type:        schema.TypeString,
									},
									{
										Name:        "path",
										Description: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
										Type:        schema.TypeString,
									},
									{
										Name:        "mode",
										Description: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional",
										Type:        schema.TypeInt,
									},
								},
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_ephemeral_volume_claim_template_object_meta_owner_references",
						Description: "OwnerReference contains enough information to let you identify an owning object",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceEphemeralVolumeClaimTemplateObjectMetaOwnerReferences,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "api_version",
								Description: "API version of the referent.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("APIVersion"),
							},
							{
								Name:        "kind",
								Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Type:        schema.TypeString,
							},
							{
								Name:        "name",
								Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
								Type:        schema.TypeString,
							},
							{
								Name:        "uid",
								Description: "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("UID"),
							},
							{
								Name:        "controller",
								Description: "If true, this reference points to the managing controller. +optional",
								Type:        schema.TypeBool,
							},
							{
								Name:        "block_owner_deletion",
								Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned. +optional",
								Type:        schema.TypeBool,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_ephemeral_volume_claim_template_object_meta_managed_fields",
						Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceEphemeralVolumeClaimTemplateObjectMetaManagedFields,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "manager",
								Description: "Manager is an identifier of the workflow managing these fields.",
								Type:        schema.TypeString,
							},
							{
								Name:        "operation",
								Description: "Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.",
								Type:        schema.TypeString,
							},
							{
								Name:        "api_version",
								Description: "APIVersion defines the version of this resource that this field set applies to",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("APIVersion"),
							},
							{
								Name:        "fields_type",
								Description: "FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"",
								Type:        schema.TypeString,
							},
							{
								Name:        "subresource",
								Description: "Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_volume_volume_source_ephemeral_volume_claim_template_spec_selector_match_expressions",
						Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
						Resolver:    fetchAppsDeploymentTemplateSpecVolumeVolumeSourceEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_volume_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key",
								Description: "key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge",
								Type:        schema.TypeString,
							},
							{
								Name:        "operator",
								Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
								Type:        schema.TypeString,
							},
							{
								Name:        "values",
								Description: "values is an array of string values",
								Type:        schema.TypeStringArray,
							},
						},
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_init_containers",
				Description: "A single application container that you want to run within a pod.",
				Resolver:    fetchAppsDeploymentTemplateSpecInitContainers,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image",
						Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "command",
						Description: "Entrypoint array",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "args",
						Description: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "working_dir",
						Description: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "resources_limits",
						Description: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Resources.Limits"),
					},
					{
						Name:        "resources_requests",
						Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Resources.Requests"),
					},
					{
						Name:        "liveness_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.Exec.Command"),
					},
					{
						Name:        "liveness_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "liveness_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "liveness_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "liveness_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LivenessProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "liveness_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "liveness_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "liveness_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "liveness_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "liveness_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LivenessProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "liveness_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "liveness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "liveness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.TimeoutSeconds"),
					},
					{
						Name:        "liveness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.PeriodSeconds"),
					},
					{
						Name:        "liveness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.SuccessThreshold"),
					},
					{
						Name:        "liveness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.FailureThreshold"),
					},
					{
						Name:        "liveness_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "readiness_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.Exec.Command"),
					},
					{
						Name:        "readiness_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "readiness_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "readiness_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "readiness_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "readiness_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "readiness_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "readiness_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "readiness_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "readiness_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "readiness_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "readiness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "readiness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.TimeoutSeconds"),
					},
					{
						Name:        "readiness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.PeriodSeconds"),
					},
					{
						Name:        "readiness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.SuccessThreshold"),
					},
					{
						Name:        "readiness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.FailureThreshold"),
					},
					{
						Name:        "readiness_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "startup_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("StartupProbe.Handler.Exec.Command"),
					},
					{
						Name:        "startup_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "startup_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "startup_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "startup_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("StartupProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "startup_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "startup_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "startup_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "startup_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "startup_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("StartupProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "startup_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "startup_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.InitialDelaySeconds"),
					},
					{
						Name:        "startup_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.TimeoutSeconds"),
					},
					{
						Name:        "startup_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.PeriodSeconds"),
					},
					{
						Name:        "startup_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.SuccessThreshold"),
					},
					{
						Name:        "startup_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.FailureThreshold"),
					},
					{
						Name:        "startup_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "lifecycle_post_start_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.Exec.Command"),
					},
					{
						Name:        "lifecycle_post_start_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.HTTPGet.Path"),
					},
					{
						Name:     "lifecycle_post_start_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.HTTPGet.Port.Type"),
					},
					{
						Name:     "lifecycle_post_start_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "lifecycle_post_start_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PostStart.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "lifecycle_post_start_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.HTTPGet.Host"),
					},
					{
						Name:        "lifecycle_post_start_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.HTTPGet.Scheme"),
					},
					{
						Name:     "lifecycle_post_start_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.TCPSocket.Port.Type"),
					},
					{
						Name:     "lifecycle_post_start_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "lifecycle_post_start_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PostStart.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "lifecycle_post_start_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.TCPSocket.Host"),
					},
					{
						Name:        "lifecycle_pre_stop_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.Exec.Command"),
					},
					{
						Name:        "lifecycle_pre_stop_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.HTTPGet.Path"),
					},
					{
						Name:     "lifecycle_pre_stop_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.HTTPGet.Port.Type"),
					},
					{
						Name:     "lifecycle_pre_stop_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "lifecycle_pre_stop_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PreStop.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "lifecycle_pre_stop_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.HTTPGet.Host"),
					},
					{
						Name:        "lifecycle_pre_stop_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.HTTPGet.Scheme"),
					},
					{
						Name:     "lifecycle_pre_stop_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.TCPSocket.Port.Type"),
					},
					{
						Name:     "lifecycle_pre_stop_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "lifecycle_pre_stop_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PreStop.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "lifecycle_pre_stop_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.TCPSocket.Host"),
					},
					{
						Name:        "termination_message_path",
						Description: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes",
						Type:        schema.TypeString,
					},
					{
						Name:        "termination_message_policy",
						Description: "Indicate how the termination message should be populated",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_pull_policy",
						Description: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "security_context_capabilities_add",
						Description: "Added capabilities +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityContext.Capabilities.Add"),
					},
					{
						Name:        "security_context_capabilities_drop",
						Description: "Removed capabilities +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityContext.Capabilities.Drop"),
					},
					{
						Name:        "security_context_privileged",
						Description: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.Privileged"),
					},
					{
						Name:        "security_context_s_e_linux_options_user",
						Description: "User is a SELinux user label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.User"),
					},
					{
						Name:        "security_context_s_e_linux_options_role",
						Description: "Role is a SELinux role label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.Role"),
					},
					{
						Name:        "security_context_s_e_linux_options_type",
						Description: "Type is a SELinux type label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.Type"),
					},
					{
						Name:        "security_context_s_e_linux_options_level",
						Description: "Level is SELinux level label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.Level"),
					},
					{
						Name:        "security_context_windows_options_g_m_s_a_credential_spec_name",
						Description: "GMSACredentialSpecName is the name of the GMSA credential spec to use. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.GMSACredentialSpecName"),
					},
					{
						Name:        "security_context_windows_options_g_m_s_a_credential_spec",
						Description: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.GMSACredentialSpec"),
					},
					{
						Name:        "security_context_windows_options_run_as_user_name",
						Description: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.RunAsUserName"),
					},
					{
						Name:        "security_context_windows_options_host_process",
						Description: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.HostProcess"),
					},
					{
						Name:        "security_context_run_as_user",
						Description: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SecurityContext.RunAsUser"),
					},
					{
						Name:        "security_context_run_as_group",
						Description: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SecurityContext.RunAsGroup"),
					},
					{
						Name:        "security_context_run_as_non_root",
						Description: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.RunAsNonRoot"),
					},
					{
						Name:        "security_context_read_only_root_filesystem",
						Description: "Whether this container has a read-only root filesystem. Default is false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.ReadOnlyRootFilesystem"),
					},
					{
						Name:        "security_context_allow_privilege_escalation",
						Description: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.AllowPrivilegeEscalation"),
					},
					{
						Name:        "security_context_proc_mount",
						Description: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.ProcMount"),
					},
					{
						Name:        "security_context_seccomp_profile_type",
						Description: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied. +unionDiscriminator",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SeccompProfile.Type"),
					},
					{
						Name:        "security_context_seccomp_profile_localhost_profile",
						Description: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is \"Localhost\". +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SeccompProfile.LocalhostProfile"),
					},
					{
						Name:        "stdin",
						Description: "Whether this container should allocate a buffer for stdin in the container runtime",
						Type:        schema.TypeBool,
					},
					{
						Name:        "stdin_once",
						Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
						Type:        schema.TypeBool,
					},
					{
						Name:        "tty",
						Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TTY"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_ports",
						Description: "ContainerPort represents a network port in a single container.",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerPorts,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_port",
								Description: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this. +optional",
								Type:        schema.TypeInt,
							},
							{
								Name:        "container_port",
								Description: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "protocol",
								Description: "Protocol for port",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_ip",
								Description: "What host IP to bind the external port to. +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("HostIP"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_env_from",
						Description: "EnvFromSource represents the source of a set of ConfigMaps",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerEnvFroms,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "prefix",
								Description: "An optional identifier to prepend to each key in the ConfigMap",
								Type:        schema.TypeString,
							},
							{
								Name:        "config_map_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ConfigMapRef.LocalObjectReference.Name"),
							},
							{
								Name:        "config_map_ref_optional",
								Description: "Specify whether the ConfigMap must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ConfigMapRef.Optional"),
							},
							{
								Name:        "secret_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SecretRef.LocalObjectReference.Name"),
							},
							{
								Name:        "secret_ref_optional",
								Description: "Specify whether the Secret must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("SecretRef.Optional"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_env",
						Description: "EnvVar represents an environment variable present in a Container.",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerEnvs,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Name of the environment variable",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables",
								Type:        schema.TypeString,
							},
							{
								Name:        "value_from_field_ref_api_version",
								Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\". +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.FieldRef.APIVersion"),
							},
							{
								Name:        "value_from_field_ref_field_path",
								Description: "Path of the field to select in the specified API version.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.FieldRef.FieldPath"),
							},
							{
								Name:        "value_from_resource_field_ref_container_name",
								Description: "Container name: required for volumes, optional for env vars +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.ContainerName"),
							},
							{
								Name:        "value_from_resource_field_ref_resource",
								Description: "Required: resource to select",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.Resource"),
							},
							{
								Name:     "value_from_resource_field_ref_divisor_format",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ValueFrom.ResourceFieldRef.Divisor.Format"),
							},
							{
								Name:        "value_from_config_map_key_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_config_map_key_ref_key",
								Description: "The key to select.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
							},
							{
								Name:        "value_from_config_map_key_ref_optional",
								Description: "Specify whether the ConfigMap or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
							},
							{
								Name:        "value_from_secret_key_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_secret_key_ref_key",
								Description: "The key of the secret to select from",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
							},
							{
								Name:        "value_from_secret_key_ref_optional",
								Description: "Specify whether the Secret or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_volume_mounts",
						Description: "VolumeMount describes a mounting of a Volume within a container.",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerVolumeMounts,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "This must match the Name of a Volume.",
								Type:        schema.TypeString,
							},
							{
								Name:        "read_only",
								Description: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional",
								Type:        schema.TypeBool,
							},
							{
								Name:        "mount_path",
								Description: "Path within the container at which the volume should be mounted",
								Type:        schema.TypeString,
							},
							{
								Name:        "sub_path",
								Description: "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root). +optional",
								Type:        schema.TypeString,
							},
							{
								Name:        "mount_propagation",
								Description: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. +optional",
								Type:        schema.TypeString,
							},
							{
								Name:        "sub_path_expr",
								Description: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to \"\" (volume's root). SubPathExpr and SubPath are mutually exclusive. +optional",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_volume_devices",
						Description: "volumeDevice describes a mapping of a raw block device within a container.",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerVolumeDevices,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "name must match the name of a persistentVolumeClaim in the pod",
								Type:        schema.TypeString,
							},
							{
								Name:        "device_path",
								Description: "devicePath is the path inside of the container that the device will be mapped to.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_liveness_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerLivenessProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_readiness_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerReadinessProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_startup_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerStartupProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_lifecycle_post_start_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerLifecyclePostStartHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_init_container_lifecycle_pre_stop_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_init_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_containers",
				Description: "A single application container that you want to run within a pod.",
				Resolver:    fetchAppsDeploymentTemplateSpecContainers,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image",
						Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "command",
						Description: "Entrypoint array",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "args",
						Description: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "working_dir",
						Description: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "resources_limits",
						Description: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Resources.Limits"),
					},
					{
						Name:        "resources_requests",
						Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Resources.Requests"),
					},
					{
						Name:        "liveness_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.Exec.Command"),
					},
					{
						Name:        "liveness_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "liveness_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "liveness_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "liveness_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LivenessProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "liveness_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "liveness_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "liveness_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "liveness_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("LivenessProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "liveness_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LivenessProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "liveness_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "liveness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "liveness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.TimeoutSeconds"),
					},
					{
						Name:        "liveness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.PeriodSeconds"),
					},
					{
						Name:        "liveness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.SuccessThreshold"),
					},
					{
						Name:        "liveness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LivenessProbe.FailureThreshold"),
					},
					{
						Name:        "liveness_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "readiness_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.Exec.Command"),
					},
					{
						Name:        "readiness_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "readiness_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "readiness_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "readiness_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "readiness_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "readiness_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "readiness_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "readiness_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "readiness_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "readiness_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "readiness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "readiness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.TimeoutSeconds"),
					},
					{
						Name:        "readiness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.PeriodSeconds"),
					},
					{
						Name:        "readiness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.SuccessThreshold"),
					},
					{
						Name:        "readiness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ReadinessProbe.FailureThreshold"),
					},
					{
						Name:        "readiness_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "startup_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("StartupProbe.Handler.Exec.Command"),
					},
					{
						Name:        "startup_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "startup_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "startup_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "startup_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("StartupProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "startup_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "startup_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "startup_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "startup_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("StartupProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "startup_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("StartupProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "startup_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "startup_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.InitialDelaySeconds"),
					},
					{
						Name:        "startup_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.TimeoutSeconds"),
					},
					{
						Name:        "startup_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.PeriodSeconds"),
					},
					{
						Name:        "startup_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.SuccessThreshold"),
					},
					{
						Name:        "startup_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("StartupProbe.FailureThreshold"),
					},
					{
						Name:        "startup_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "lifecycle_post_start_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.Exec.Command"),
					},
					{
						Name:        "lifecycle_post_start_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.HTTPGet.Path"),
					},
					{
						Name:     "lifecycle_post_start_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.HTTPGet.Port.Type"),
					},
					{
						Name:     "lifecycle_post_start_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "lifecycle_post_start_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PostStart.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "lifecycle_post_start_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.HTTPGet.Host"),
					},
					{
						Name:        "lifecycle_post_start_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.HTTPGet.Scheme"),
					},
					{
						Name:     "lifecycle_post_start_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.TCPSocket.Port.Type"),
					},
					{
						Name:     "lifecycle_post_start_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PostStart.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "lifecycle_post_start_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PostStart.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "lifecycle_post_start_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PostStart.TCPSocket.Host"),
					},
					{
						Name:        "lifecycle_pre_stop_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.Exec.Command"),
					},
					{
						Name:        "lifecycle_pre_stop_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.HTTPGet.Path"),
					},
					{
						Name:     "lifecycle_pre_stop_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.HTTPGet.Port.Type"),
					},
					{
						Name:     "lifecycle_pre_stop_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "lifecycle_pre_stop_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PreStop.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "lifecycle_pre_stop_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.HTTPGet.Host"),
					},
					{
						Name:        "lifecycle_pre_stop_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.HTTPGet.Scheme"),
					},
					{
						Name:     "lifecycle_pre_stop_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.TCPSocket.Port.Type"),
					},
					{
						Name:     "lifecycle_pre_stop_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Lifecycle.PreStop.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "lifecycle_pre_stop_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Lifecycle.PreStop.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "lifecycle_pre_stop_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lifecycle.PreStop.TCPSocket.Host"),
					},
					{
						Name:        "termination_message_path",
						Description: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes",
						Type:        schema.TypeString,
					},
					{
						Name:        "termination_message_policy",
						Description: "Indicate how the termination message should be populated",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_pull_policy",
						Description: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "security_context_capabilities_add",
						Description: "Added capabilities +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityContext.Capabilities.Add"),
					},
					{
						Name:        "security_context_capabilities_drop",
						Description: "Removed capabilities +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityContext.Capabilities.Drop"),
					},
					{
						Name:        "security_context_privileged",
						Description: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.Privileged"),
					},
					{
						Name:        "security_context_s_e_linux_options_user",
						Description: "User is a SELinux user label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.User"),
					},
					{
						Name:        "security_context_s_e_linux_options_role",
						Description: "Role is a SELinux role label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.Role"),
					},
					{
						Name:        "security_context_s_e_linux_options_type",
						Description: "Type is a SELinux type label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.Type"),
					},
					{
						Name:        "security_context_s_e_linux_options_level",
						Description: "Level is SELinux level label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SELinuxOptions.Level"),
					},
					{
						Name:        "security_context_windows_options_g_m_s_a_credential_spec_name",
						Description: "GMSACredentialSpecName is the name of the GMSA credential spec to use. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.GMSACredentialSpecName"),
					},
					{
						Name:        "security_context_windows_options_g_m_s_a_credential_spec",
						Description: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.GMSACredentialSpec"),
					},
					{
						Name:        "security_context_windows_options_run_as_user_name",
						Description: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.RunAsUserName"),
					},
					{
						Name:        "security_context_windows_options_host_process",
						Description: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.WindowsOptions.HostProcess"),
					},
					{
						Name:        "security_context_run_as_user",
						Description: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SecurityContext.RunAsUser"),
					},
					{
						Name:        "security_context_run_as_group",
						Description: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SecurityContext.RunAsGroup"),
					},
					{
						Name:        "security_context_run_as_non_root",
						Description: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.RunAsNonRoot"),
					},
					{
						Name:        "security_context_read_only_root_filesystem",
						Description: "Whether this container has a read-only root filesystem. Default is false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.ReadOnlyRootFilesystem"),
					},
					{
						Name:        "security_context_allow_privilege_escalation",
						Description: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SecurityContext.AllowPrivilegeEscalation"),
					},
					{
						Name:        "security_context_proc_mount",
						Description: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.ProcMount"),
					},
					{
						Name:        "security_context_seccomp_profile_type",
						Description: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied. +unionDiscriminator",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SeccompProfile.Type"),
					},
					{
						Name:        "security_context_seccomp_profile_localhost_profile",
						Description: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is \"Localhost\". +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityContext.SeccompProfile.LocalhostProfile"),
					},
					{
						Name:        "stdin",
						Description: "Whether this container should allocate a buffer for stdin in the container runtime",
						Type:        schema.TypeBool,
					},
					{
						Name:        "stdin_once",
						Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
						Type:        schema.TypeBool,
					},
					{
						Name:        "tty",
						Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TTY"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_apps_deployment_template_spec_container_ports",
						Description: "ContainerPort represents a network port in a single container.",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerPorts,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_port",
								Description: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this. +optional",
								Type:        schema.TypeInt,
							},
							{
								Name:        "container_port",
								Description: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "protocol",
								Description: "Protocol for port",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_ip",
								Description: "What host IP to bind the external port to. +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("HostIP"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_env_from",
						Description: "EnvFromSource represents the source of a set of ConfigMaps",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerEnvFroms,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "prefix",
								Description: "An optional identifier to prepend to each key in the ConfigMap",
								Type:        schema.TypeString,
							},
							{
								Name:        "config_map_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ConfigMapRef.LocalObjectReference.Name"),
							},
							{
								Name:        "config_map_ref_optional",
								Description: "Specify whether the ConfigMap must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ConfigMapRef.Optional"),
							},
							{
								Name:        "secret_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SecretRef.LocalObjectReference.Name"),
							},
							{
								Name:        "secret_ref_optional",
								Description: "Specify whether the Secret must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("SecretRef.Optional"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_env",
						Description: "EnvVar represents an environment variable present in a Container.",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerEnvs,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Name of the environment variable",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables",
								Type:        schema.TypeString,
							},
							{
								Name:        "value_from_field_ref_api_version",
								Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\". +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.FieldRef.APIVersion"),
							},
							{
								Name:        "value_from_field_ref_field_path",
								Description: "Path of the field to select in the specified API version.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.FieldRef.FieldPath"),
							},
							{
								Name:        "value_from_resource_field_ref_container_name",
								Description: "Container name: required for volumes, optional for env vars +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.ContainerName"),
							},
							{
								Name:        "value_from_resource_field_ref_resource",
								Description: "Required: resource to select",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.Resource"),
							},
							{
								Name:     "value_from_resource_field_ref_divisor_format",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ValueFrom.ResourceFieldRef.Divisor.Format"),
							},
							{
								Name:        "value_from_config_map_key_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_config_map_key_ref_key",
								Description: "The key to select.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
							},
							{
								Name:        "value_from_config_map_key_ref_optional",
								Description: "Specify whether the ConfigMap or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
							},
							{
								Name:        "value_from_secret_key_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_secret_key_ref_key",
								Description: "The key of the secret to select from",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
							},
							{
								Name:        "value_from_secret_key_ref_optional",
								Description: "Specify whether the Secret or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_volume_mounts",
						Description: "VolumeMount describes a mounting of a Volume within a container.",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerVolumeMounts,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "This must match the Name of a Volume.",
								Type:        schema.TypeString,
							},
							{
								Name:        "read_only",
								Description: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional",
								Type:        schema.TypeBool,
							},
							{
								Name:        "mount_path",
								Description: "Path within the container at which the volume should be mounted",
								Type:        schema.TypeString,
							},
							{
								Name:        "sub_path",
								Description: "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root). +optional",
								Type:        schema.TypeString,
							},
							{
								Name:        "mount_propagation",
								Description: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. +optional",
								Type:        schema.TypeString,
							},
							{
								Name:        "sub_path_expr",
								Description: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to \"\" (volume's root). SubPathExpr and SubPath are mutually exclusive. +optional",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_volume_devices",
						Description: "volumeDevice describes a mapping of a raw block device within a container.",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerVolumeDevices,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "name must match the name of a persistentVolumeClaim in the pod",
								Type:        schema.TypeString,
							},
							{
								Name:        "device_path",
								Description: "devicePath is the path inside of the container that the device will be mapped to.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_liveness_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerLivenessProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_readiness_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerReadinessProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_startup_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerStartupProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_lifecycle_post_start_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerLifecyclePostStartHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_container_lifecycle_pre_stop_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_ephemeral_containers",
				Description: "An EphemeralContainer is a container that may be added temporarily to an existing pod for user-initiated activities such as debugging",
				Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainers,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ephemeral_container_common_name",
						Description: "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Name"),
					},
					{
						Name:        "ephemeral_container_common_image",
						Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Image"),
					},
					{
						Name:        "ephemeral_container_common_command",
						Description: "Entrypoint array",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Command"),
					},
					{
						Name:        "ephemeral_container_common_args",
						Description: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Args"),
					},
					{
						Name:        "ephemeral_container_common_working_dir",
						Description: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.WorkingDir"),
					},
					{
						Name:        "ephemeral_container_common_resources_limits",
						Description: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Resources.Limits"),
					},
					{
						Name:        "ephemeral_container_common_resources_requests",
						Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Resources.Requests"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.Exec.Command"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "ephemeral_container_common_liveness_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_liveness_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_liveness_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "ephemeral_container_common_liveness_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_liveness_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_liveness_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.TimeoutSeconds"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.PeriodSeconds"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.SuccessThreshold"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.FailureThreshold"),
					},
					{
						Name:        "ephemeral_container_common_liveness_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.LivenessProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.Exec.Command"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "ephemeral_container_common_readiness_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_readiness_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_readiness_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "ephemeral_container_common_readiness_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_readiness_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_readiness_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.TimeoutSeconds"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.PeriodSeconds"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.SuccessThreshold"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.FailureThreshold"),
					},
					{
						Name:        "ephemeral_container_common_readiness_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ReadinessProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_handler_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.Exec.Command"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_handler_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.HTTPGet.Path"),
					},
					{
						Name:     "ephemeral_container_common_startup_probe_handler_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.HTTPGet.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_startup_probe_handler_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_startup_probe_handler_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_handler_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.HTTPGet.Host"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_handler_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.HTTPGet.Scheme"),
					},
					{
						Name:     "ephemeral_container_common_startup_probe_handler_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.TCPSocket.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_startup_probe_handler_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_startup_probe_handler_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_handler_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.Handler.TCPSocket.Host"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.InitialDelaySeconds"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out. Defaults to 1 second",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.TimeoutSeconds"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe. Default to 10 seconds",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.PeriodSeconds"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.SuccessThreshold"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.FailureThreshold"),
					},
					{
						Name:        "ephemeral_container_common_startup_probe_termination_grace_period_seconds",
						Description: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StartupProbe.TerminationGracePeriodSeconds"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_post_start_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.Exec.Command"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_post_start_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet.Path"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_post_start_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_post_start_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_post_start_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_post_start_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet.Host"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_post_start_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet.Scheme"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_post_start_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.TCPSocket.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_post_start_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_post_start_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_post_start_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PostStart.TCPSocket.Host"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_pre_stop_exec_command",
						Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.Exec.Command"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_pre_stop_http_get_path",
						Description: "Path to access on the HTTP server. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet.Path"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_pre_stop_http_get_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_pre_stop_http_get_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_pre_stop_http_get_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_pre_stop_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet.Host"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_pre_stop_http_get_scheme",
						Description: "Scheme to use for connecting to the host. Defaults to HTTP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet.Scheme"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_pre_stop_tcp_socket_port_type",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.TCPSocket.Port.Type"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_pre_stop_tcp_socket_port_int_val",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.TCPSocket.Port.IntVal"),
					},
					{
						Name:     "ephemeral_container_common_lifecycle_pre_stop_tcp_socket_port_str_val",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.TCPSocket.Port.StrVal"),
					},
					{
						Name:        "ephemeral_container_common_lifecycle_pre_stop_tcp_socket_host",
						Description: "Optional: Host name to connect to, defaults to the pod IP. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Lifecycle.PreStop.TCPSocket.Host"),
					},
					{
						Name:        "ephemeral_container_common_termination_message_path",
						Description: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.TerminationMessagePath"),
					},
					{
						Name:        "ephemeral_container_common_termination_message_policy",
						Description: "Indicate how the termination message should be populated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.TerminationMessagePolicy"),
					},
					{
						Name:        "ephemeral_container_common_image_pull_policy",
						Description: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ImagePullPolicy"),
					},
					{
						Name:        "ephemeral_container_common_security_context_capabilities_add",
						Description: "Added capabilities +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.Capabilities.Add"),
					},
					{
						Name:        "ephemeral_container_common_security_context_capabilities_drop",
						Description: "Removed capabilities +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.Capabilities.Drop"),
					},
					{
						Name:        "ephemeral_container_common_security_context_privileged",
						Description: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.Privileged"),
					},
					{
						Name:        "ephemeral_container_common_security_context_s_e_linux_options_user",
						Description: "User is a SELinux user label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.SELinuxOptions.User"),
					},
					{
						Name:        "ephemeral_container_common_security_context_s_e_linux_options_role",
						Description: "Role is a SELinux role label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.SELinuxOptions.Role"),
					},
					{
						Name:        "ephemeral_container_common_security_context_s_e_linux_options_type",
						Description: "Type is a SELinux type label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.SELinuxOptions.Type"),
					},
					{
						Name:        "ephemeral_container_common_security_context_s_e_linux_options_level",
						Description: "Level is SELinux level label that applies to the container. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.SELinuxOptions.Level"),
					},
					{
						Name:        "ephemeral_container_common_security_context_windows_options_g_m_s_a_credential_spec_name",
						Description: "GMSACredentialSpecName is the name of the GMSA credential spec to use. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.WindowsOptions.GMSACredentialSpecName"),
					},
					{
						Name:        "ephemeral_container_common_security_context_windows_options_g_m_s_a_credential_spec",
						Description: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.WindowsOptions.GMSACredentialSpec"),
					},
					{
						Name:        "ephemeral_container_common_security_context_windows_options_run_as_user_name",
						Description: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.WindowsOptions.RunAsUserName"),
					},
					{
						Name:        "ephemeral_container_common_security_context_windows_options_host_process",
						Description: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.WindowsOptions.HostProcess"),
					},
					{
						Name:        "ephemeral_container_common_security_context_run_as_user",
						Description: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.RunAsUser"),
					},
					{
						Name:        "ephemeral_container_common_security_context_run_as_group",
						Description: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.RunAsGroup"),
					},
					{
						Name:        "ephemeral_container_common_security_context_run_as_non_root",
						Description: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.RunAsNonRoot"),
					},
					{
						Name:        "ephemeral_container_common_security_context_read_only_root_filesystem",
						Description: "Whether this container has a read-only root filesystem. Default is false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.ReadOnlyRootFilesystem"),
					},
					{
						Name:        "ephemeral_container_common_security_context_allow_privilege_escalation",
						Description: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.AllowPrivilegeEscalation"),
					},
					{
						Name:        "ephemeral_container_common_security_context_proc_mount",
						Description: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.ProcMount"),
					},
					{
						Name:        "ephemeral_container_common_security_context_seccomp_profile_type",
						Description: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied. +unionDiscriminator",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.SeccompProfile.Type"),
					},
					{
						Name:        "ephemeral_container_common_security_context_seccomp_profile_localhost_profile",
						Description: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is \"Localhost\". +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.SecurityContext.SeccompProfile.LocalhostProfile"),
					},
					{
						Name:        "ephemeral_container_common_stdin",
						Description: "Whether this container should allocate a buffer for stdin in the container runtime",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Stdin"),
					},
					{
						Name:        "ephemeral_container_common_stdin_once",
						Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StdinOnce"),
					},
					{
						Name:        "ephemeral_container_common_tty",
						Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false. +optional",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.TTY"),
					},
					{
						Name:        "target_container_name",
						Description: "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container is run in whatever namespaces are shared for the pod",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_ports",
						Description: "ContainerPort represents a network port in a single container.",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonPorts,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_port",
								Description: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this. +optional",
								Type:        schema.TypeInt,
							},
							{
								Name:        "container_port",
								Description: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "protocol",
								Description: "Protocol for port",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_ip",
								Description: "What host IP to bind the external port to. +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("HostIP"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_env_from",
						Description: "EnvFromSource represents the source of a set of ConfigMaps",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonEnvFroms,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "prefix",
								Description: "An optional identifier to prepend to each key in the ConfigMap",
								Type:        schema.TypeString,
							},
							{
								Name:        "config_map_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ConfigMapRef.LocalObjectReference.Name"),
							},
							{
								Name:        "config_map_ref_optional",
								Description: "Specify whether the ConfigMap must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ConfigMapRef.Optional"),
							},
							{
								Name:        "secret_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SecretRef.LocalObjectReference.Name"),
							},
							{
								Name:        "secret_ref_optional",
								Description: "Specify whether the Secret must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("SecretRef.Optional"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_env",
						Description: "EnvVar represents an environment variable present in a Container.",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonEnvs,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Name of the environment variable",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables",
								Type:        schema.TypeString,
							},
							{
								Name:        "value_from_field_ref_api_version",
								Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\". +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.FieldRef.APIVersion"),
							},
							{
								Name:        "value_from_field_ref_field_path",
								Description: "Path of the field to select in the specified API version.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.FieldRef.FieldPath"),
							},
							{
								Name:        "value_from_resource_field_ref_container_name",
								Description: "Container name: required for volumes, optional for env vars +optional",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.ContainerName"),
							},
							{
								Name:        "value_from_resource_field_ref_resource",
								Description: "Required: resource to select",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.Resource"),
							},
							{
								Name:     "value_from_resource_field_ref_divisor_format",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ValueFrom.ResourceFieldRef.Divisor.Format"),
							},
							{
								Name:        "value_from_config_map_key_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_config_map_key_ref_key",
								Description: "The key to select.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
							},
							{
								Name:        "value_from_config_map_key_ref_optional",
								Description: "Specify whether the ConfigMap or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
							},
							{
								Name:        "value_from_secret_key_ref_local_object_reference_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_secret_key_ref_key",
								Description: "The key of the secret to select from",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
							},
							{
								Name:        "value_from_secret_key_ref_optional",
								Description: "Specify whether the Secret or its key must be defined +optional",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_volume_mounts",
						Description: "VolumeMount describes a mounting of a Volume within a container.",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonVolumeMounts,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "This must match the Name of a Volume.",
								Type:        schema.TypeString,
							},
							{
								Name:        "read_only",
								Description: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional",
								Type:        schema.TypeBool,
							},
							{
								Name:        "mount_path",
								Description: "Path within the container at which the volume should be mounted",
								Type:        schema.TypeString,
							},
							{
								Name:        "sub_path",
								Description: "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root). +optional",
								Type:        schema.TypeString,
							},
							{
								Name:        "mount_propagation",
								Description: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. +optional",
								Type:        schema.TypeString,
							},
							{
								Name:        "sub_path_expr",
								Description: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to \"\" (volume's root). SubPathExpr and SubPath are mutually exclusive. +optional",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_volume_devices",
						Description: "volumeDevice describes a mapping of a raw block device within a container.",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonVolumeDevices,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "name must match the name of a persistentVolumeClaim in the pod",
								Type:        schema.TypeString,
							},
							{
								Name:        "device_path",
								Description: "devicePath is the path inside of the container that the device will be mapped to.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_liveness_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonLivenessProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_readiness_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonReadinessProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_startup_probe_handler_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonStartupProbeHandlerHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_lifecycle_post_start_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonLifecyclePostStartHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "k8s_apps_deployment_template_spec_ephemeral_container_ephemeral_container_common_lifecycle_pre_stop_http_get_http_headers",
						Description: "HTTPHeader describes a custom header to be used in HTTP probes",
						Resolver:    fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonLifecyclePreStopHttpGetHttpHeaders,
						Columns: []schema.Column{
							{
								Name:        "deployment_template_spec_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_deployment_template_spec_ephemeral_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The header field name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The header field value",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_image_pull_secrets",
				Description: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace. +structType=atomic",
				Resolver:    fetchAppsDeploymentTemplateSpecImagePullSecrets,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_tolerations",
				Description: "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.",
				Resolver:    fetchAppsDeploymentTemplateSpecTolerations,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "Key is the taint key that the toleration applies to",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "effect",
						Description: "Effect indicates the taint effect to match",
						Type:        schema.TypeString,
					},
					{
						Name:        "toleration_seconds",
						Description: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint",
						Type:        schema.TypeBigInt,
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_template_spec_host_aliases",
				Description: "HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.",
				Resolver:    fetchAppsDeploymentTemplateSpecHostAliases,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ip",
						Description: "IP address of the host file entry.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IP"),
					},
					{
						Name:        "hostnames",
						Description: "Hostnames for the above IP address.",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "k8s_apps_deployment_status_conditions",
				Description: "DeploymentCondition describes the state of a deployment at a certain point.",
				Resolver:    fetchAppsDeploymentStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of deployment condition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown.",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "The reason for the condition's last transition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "A human readable message indicating details about the transition.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAppsDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services().Deployments
	opts := metav1.ListOptions{}
	for {
		result, err := client.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
func resolveAppsDeploymentsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsTemplateSpecSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsTemplateSpecAffinity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsTemplateSpecDnsConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsTemplateSpecReadinessGates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsTemplateSpecOverhead(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveAppsDeploymentsTemplateSpecTopologySpreadConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchAppsDeploymentSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	deployment, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	if deployment.Spec.Selector == nil {
		return nil
	}
	res <- deployment.Spec.Selector.MatchExpressions
	return nil
}
func fetchAppsDeploymentTemplateOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	deployment, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- deployment.Spec.Template.OwnerReferences
	return nil
}
func fetchAppsDeploymentTemplateManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	deployment, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- deployment.Spec.Template.ManagedFields
	return nil
}
func fetchAppsDeploymentTemplateSpecVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	deployment, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- deployment.Spec.Template.Spec.Volumes
	return nil
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceSecretItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceDownwardApiItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceConfigMapItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSourceSecretItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSourceDownwardApiItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceProjectedSourceConfigMapItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceEphemeralVolumeClaimTemplateObjectMetaOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceEphemeralVolumeClaimTemplateObjectMetaManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecVolumeVolumeSourceEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.InitContainers
	return nil
}
func fetchAppsDeploymentTemplateSpecInitContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}

	res <- p.Ports
	return nil
}
func fetchAppsDeploymentTemplateSpecInitContainerEnvFroms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}

	res <- p.Env
	return nil
}
func fetchAppsDeploymentTemplateSpecInitContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerLivenessProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerReadinessProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerStartupProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerLifecyclePostStartHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.Containers
	return nil
}
func fetchAppsDeploymentTemplateSpecContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchAppsDeploymentTemplateSpecContainerEnvFroms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Env
	return nil
}
func fetchAppsDeploymentTemplateSpecContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerLivenessProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerReadinessProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerStartupProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerLifecyclePostStartHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.EphemeralContainers
	return nil
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonEnvFroms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonVolumeMounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonVolumeDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonLivenessProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonReadinessProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonStartupProbeHandlerHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonLifecyclePostStartHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEphemeralContainerCommonLifecyclePreStopHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchAppsDeploymentTemplateSpecImagePullSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}

	res <- p.Spec.Template.Spec.ImagePullSecrets
	return nil
}
func fetchAppsDeploymentTemplateSpecTolerations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}

	res <- p.Spec.Template.Spec.Tolerations
	return nil
}
func fetchAppsDeploymentTemplateSpecHostAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}

	res <- p.Spec.Template.Spec.HostAliases
	return nil
}
func fetchAppsDeploymentStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}

	res <- p.Status.Conditions
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func resolveAppsDeploymentOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	if p.Spec.Template.Spec.SecurityContext == nil {
		return nil
	}
	b, err := json.Marshal(p.Spec.Template.Spec.SecurityContext)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecAffinity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	if p.Spec.Template.Spec.Affinity == nil {
		return nil
	}
	b, err := json.Marshal(p.Spec.Template.Spec.Affinity)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecDNSConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	if p.Spec.Template.Spec.DNSConfig == nil {
		return nil
	}
	b, err := json.Marshal(p.Spec.Template.Spec.DNSConfig)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecReadinessGates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.Spec.Template.Spec.ReadinessGates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecOverhead(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template.Spec.Overhead)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecTopologySpreadConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template.Spec.TopologySpreadConstraints)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeAWSElasticBlockStore(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.AWSElasticBlockStore == nil {
		return nil
	}
	b, err := json.Marshal(p.AWSElasticBlockStore)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeNfs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.AWSElasticBlockStore == nil {
		return nil
	}
	b, err := json.Marshal(p.AWSElasticBlockStore)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeIscsi(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.ISCSI == nil {
		return nil
	}
	b, err := json.Marshal(p.ISCSI)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeRbd(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.RBD == nil {
		return nil
	}
	b, err := json.Marshal(p.RBD)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeDownwardAPI(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.DownwardAPI == nil {
		return nil
	}
	b, err := json.Marshal(p.DownwardAPI)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeStorageOs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.StorageOS == nil {
		return nil
	}
	b, err := json.Marshal(p.StorageOS)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecVolumeCsi(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}
	if p.CSI == nil {
		return nil
	}
	b, err := json.Marshal(p.CSI)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerEnvFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.EnvFrom)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.VolumeMounts)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.VolumeDevices)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerLivenessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.LivenessProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.LivenessProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerReadinessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.ReadinessProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.ReadinessProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerStartupProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.StartupProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.StartupProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerLifecycle(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.Lifecycle == nil {
		return nil
	}
	b, err := json.Marshal(p.Lifecycle)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecContainerSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.SecurityContext == nil {
		return nil
	}
	b, err := json.Marshal(p.SecurityContext)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerEnvFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.EnvFrom)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.EphemeralContainer instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.VolumeMounts)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.EphemeralContainer instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.VolumeDevices)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerLivenessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.LivenessProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.LivenessProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerReadinessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.ReadinessProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.ReadinessProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerStartupProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.StartupProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.StartupProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerLifecycle(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.Lifecycle == nil {
		return nil
	}
	b, err := json.Marshal(p.Lifecycle)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecEphemeralContainerSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}
	if p.SecurityContext == nil {
		return nil
	}
	b, err := json.Marshal(p.SecurityContext)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchAppsDeploymentTemplateSpecEphemeralContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Env
	return nil
}
func resolveAppsDeploymentTemplateSpecInitContainerEnvFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.EnvFrom)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.VolumeMounts)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.VolumeDevices)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerLivenessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	if p.LivenessProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.EnvFrom)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerReadinessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	if p.ReadinessProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.ReadinessProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerStartupProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	if p.StartupProbe == nil {
		return nil
	}
	b, err := json.Marshal(p.StartupProbe)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerLifecycle(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	if p.Lifecycle == nil {
		return nil
	}
	b, err := json.Marshal(p.Lifecycle)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentTemplateSpecInitContainerSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", resource.Item)
	}

	if p.SecurityContext == nil {
		return nil
	}
	b, err := json.Marshal(p.SecurityContext)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
