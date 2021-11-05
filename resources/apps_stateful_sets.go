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

func AppsStatefulSets() *schema.Table {
	return &schema.Table{
		Name:         "k8s_apps_stateful_sets",
		Description:  "StatefulSet represents a set of pods with consistent identities. Identities are defined as:  - Network: A single stable DNS and hostname.  - Storage: As many VolumeClaims as requested. The StatefulSet guarantees that a given network identity will always map to the same storage identity.",
		Resolver:     fetchAppsStatefulSets,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:     "generate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.GenerateName"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.SelfLink"),
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ObjectMeta.Finalizers"),
			},
			{
				Name:     "cluster_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "replicas",
				Description: "replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. TODO: Consider a rename of this field. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Replicas"),
			},
			{
				Name:     "selector_match_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:     "template_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Name"),
			},
			{
				Name:     "template_generate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.GenerateName"),
			},
			{
				Name:     "template_namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Namespace"),
			},
			{
				Name:     "template_self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.SelfLink"),
			},
			{
				Name:     "template_uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.UID"),
			},
			{
				Name:     "template_resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.ResourceVersion"),
			},
			{
				Name:     "template_generation",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Generation"),
			},
			{
				Name:     "template_deletion_grace_period_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:     "template_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Labels"),
			},
			{
				Name:     "template_annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Annotations"),
			},
			{
				Name:     "template_finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Finalizers"),
			},
			{
				Name:     "template_cluster_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.ClusterName"),
			},
			{
				Name:     "template_restart_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.RestartPolicy"),
			},
			{
				Name:     "template_termination_grace_period_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Template.Spec.TerminationGracePeriodSeconds"),
			},
			{
				Name:     "template_active_deadline_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Template.Spec.ActiveDeadlineSeconds"),
			},
			{
				Name:     "template_dns_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.DNSPolicy"),
			},
			{
				Name:     "template_node_selector",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template.Spec.NodeSelector"),
			},
			{
				Name:     "template_service_account_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.ServiceAccountName"),
			},
			{
				Name:     "template_deprecated_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.DeprecatedServiceAccount"),
			},
			{
				Name:     "template_automount_service_account_token",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.AutomountServiceAccountToken"),
			},
			{
				Name:     "template_node_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.NodeName"),
			},
			{
				Name:     "template_host_network",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.HostNetwork"),
			},
			{
				Name:     "template_host_pid",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.HostPID"),
			},
			{
				Name:     "template_host_ipc",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.HostIPC"),
			},
			{
				Name:     "template_share_process_namespace",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.ShareProcessNamespace"),
			},
			{
				Name:     "template_security_context",
				Type:     schema.TypeJSON,
				Resolver: resolveAppsStatefulSetTemplateSpecSecurityContext,
			},
			{
				Name:     "template_hostname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.Hostname"),
			},
			{
				Name:     "template_subdomain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.Subdomain"),
			},
			{
				Name:     "template_affinity",
				Type:     schema.TypeJSON,
				Resolver: resolveAppsStatefulSetTemplateSpecAffinity,
			},
			{
				Name:     "template_scheduler_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.SchedulerName"),
			},
			{
				Name:     "template_priority_class_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.PriorityClassName"),
			},
			{
				Name:     "template_priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Template.Spec.Priority"),
			},
			{
				Name:     "template_dns_config",
				Type:     schema.TypeJSON,
				Resolver: resolveAppsStatefulSetTemplateSpecDNSConfig,
			},
			{
				Name:     "template_readiness_gates",
				Type:     schema.TypeJSON,
				Resolver: resolveAppsStatefulSetTemplateSpecReadinessGates,
			},
			{
				Name:     "template_runtime_class_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.RuntimeClassName"),
			},
			{
				Name:     "template_enable_service_links",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.EnableServiceLinks"),
			},
			{
				Name:     "template_preemption_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.Spec.PreemptionPolicy"),
			},
			{
				Name:     "template_overhead",
				Type:     schema.TypeJSON,
				Resolver: resolveAppsStatefulSetTemplateSpecOverhead,
			},
			{
				Name:     "template_topology_spread_constraints",
				Type:     schema.TypeJSON,
				Resolver: resolveAppsStatefulSetTemplateSpecTopologySpreadConstraints,
			},
			{
				Name:     "template_set_hostname_as_fqdn",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Template.Spec.SetHostnameAsFQDN"),
			},
			{
				Name:        "service_name",
				Description: "serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ServiceName"),
			},
			{
				Name:        "pod_management_policy",
				Description: "podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.PodManagementPolicy"),
			},
			{
				Name:        "update_strategy_type",
				Description: "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.UpdateStrategy.Type"),
			},
			{
				Name:        "update_strategy_rolling_update_partition",
				Description: "Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.Partition"),
			},
			{
				Name:        "revision_history_limit",
				Description: "revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.RevisionHistoryLimit"),
			},
			{
				Name:        "min_ready_seconds",
				Description: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:        "status_observed_generation",
				Description: "observedGeneration is the most recent generation observed for this StatefulSet",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_replicas",
				Description: "replicas is the number of Pods created by the StatefulSet controller.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Replicas"),
			},
			{
				Name:        "status_ready_replicas",
				Description: "readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.ReadyReplicas"),
			},
			{
				Name:        "status_current_replicas",
				Description: "currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CurrentReplicas"),
			},
			{
				Name:        "status_updated_replicas",
				Description: "updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UpdatedReplicas"),
			},
			{
				Name:        "status_current_revision",
				Description: "currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.CurrentRevision"),
			},
			{
				Name:        "status_update_revision",
				Description: "updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.UpdateRevision"),
			},
			{
				Name:        "status_collision_count",
				Description: "collisionCount is the count of hash collisions for the StatefulSet",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CollisionCount"),
			},
			{
				Name:        "status_available_replicas",
				Description: "Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset. This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate. Remove omitempty when graduating to beta +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.AvailableReplicas"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_apps_stateful_set_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchAppsStatefulSetOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchAppsStatefulSetManagedFields,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchAppsStatefulSetSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_template_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchAppsStatefulSetTemplateOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_template_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchAppsStatefulSetTemplateManagedFields,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_template_volumes",
				Description: "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
				Resolver:    fetchAppsStatefulSetTemplateSpecVolumes,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						Type:        schema.TypeString,
					},
					{
						Name:        "host_path",
						Description: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.HostPath.Path"),
					},
					{
						Name:        "host_path_type",
						Description: "Type for HostPath Volume Defaults to \"\" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.HostPath.Type"),
					},
					{
						Name:        "empty_dir_medium",
						Description: "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VolumeSource.EmptyDir.Medium"),
					},
					{
						Name:     "empty_dir_size_limit_format",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VolumeSource.EmptyDir.SizeLimit.Format"),
					},
					{
						Name:        "gce_persistent_disk",
						Description: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.GCEPersistentDisk"),
					},
					{
						Name:        "aws_elastic_block_store",
						Description: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeAWSElasticBlockStore,
					},
					{
						Name:        "git_repo",
						Description: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "secret",
						Description: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "nfs",
						Description: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeNfs,
					},
					{
						Name:        "iscsi",
						Description: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeIscsi,
					},
					{
						Name:        "glusterfs",
						Description: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "persistent_volume_claim",
						Description: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "rbd",
						Description: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeRbd,
					},
					{
						Name:        "flex_volume",
						Description: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin. +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "cinder",
						Description: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "ceph_fs",
						Description: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.CephFS"),
					},
					{
						Name:        "flocker",
						Description: "Flocker represents a Flocker volume attached to a kubelet's host machine",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "downward_api",
						Description: "DownwardAPI represents downward API about the pod that should populate this volume +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeDownwardAPI,
					},
					{
						Name:        "fc",
						Description: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod. +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.FC"),
					},
					{
						Name:        "azure_file",
						Description: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod. +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "config_map",
						Description: "ConfigMap represents a configMap that should populate this volume +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "vsphere_volume",
						Description: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "quobyte",
						Description: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "azure_disk",
						Description: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod. +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "photon_persistent_disk",
						Description: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "projected",
						Description: "Items for all in one resources secrets, configmaps, and downward API",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "portworx_volume",
						Description: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine +optional",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "scale_io",
						Description: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes. +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VolumeSource.ScaleIO"),
					},
					{
						Name:        "storage_os",
						Description: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeStorageOs,
					},
					{
						Name:        "csi",
						Description: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature). +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecVolumeCsi,
					},
					{
						Name:        "ephemeral",
						Description: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.  This is a beta feature and only available when the GenericEphemeralVolume feature gate is enabled.  +optional",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "k8s_apps_stateful_set_template_containers",
				Description: "A single application container that you want to run within a pod.",
				Resolver:    fetchAppsStatefulSetTemplateSpecContainers,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
						Name:        "env_from",
						Description: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerEnvFrom,
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
						Name:        "volume_mounts",
						Description: "Pod volumes to mount into the container's filesystem. Cannot be updated. +optional +patchMergeKey=mountPath +patchStrategy=merge",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerVolumeMounts,
					},
					{
						Name:        "volume_devices",
						Description: "volumeDevices is the list of block devices to be used by the container. +patchMergeKey=devicePath +patchStrategy=merge +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerVolumeDevices,
					},
					{
						Name:        "liveness_probe",
						Description: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerLivenessProbe,
					},
					{
						Name:        "readiness_probe",
						Description: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerReadinessProbe,
					},
					{
						Name:        "startup_probe",
						Description: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerStartupProbe,
					},
					{
						Name:        "lifecycle",
						Description: "Actions that the management system should take in response to container lifecycle events. Cannot be updated. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerLifecycle,
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
						Name:        "security_context",
						Description: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecContainerSecurityContext,
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
						Name:        "k8s_apps_stateful_set_template_container_ports",
						Description: "ContainerPort represents a network port in a single container.",
						Resolver:    fetchAppsStatefulSetTemplateSpecContainerPorts,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_template_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_template_containers table (FK)",
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
						Name:        "k8s_apps_stateful_set_template_container_envs",
						Description: "EnvVar represents an environment variable present in a Container.",
						Resolver:    fetchAppsStatefulSetTemplateSpecContainerEnvs,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_template_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_template_containers table (FK)",
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
				},
			},
			{
				Name:        "k8s_apps_stateful_set_template_ephemeral_containers",
				Description: "An EphemeralContainer is a container that may be added temporarily to an existing pod for user-initiated activities such as debugging",
				Resolver:    fetchAppsStatefulSetTemplateSpecEphemeralContainers,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Name"),
					},
					{
						Name:        "image",
						Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Image"),
					},
					{
						Name:        "command",
						Description: "Entrypoint array",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Command"),
					},
					{
						Name:        "args",
						Description: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Args"),
					},
					{
						Name:        "working_dir",
						Description: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.WorkingDir"),
					},
					{
						Name:        "env_from",
						Description: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerEnvFrom,
					},
					{
						Name:        "resources_limits",
						Description: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Resources.Limits"),
					},
					{
						Name:        "resources_requests",
						Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Resources.Requests"),
					},
					{
						Name:        "volume_mounts",
						Description: "Pod volumes to mount into the container's filesystem. Cannot be updated. +optional +patchMergeKey=mountPath +patchStrategy=merge",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerVolumeMounts,
					},
					{
						Name:        "volume_devices",
						Description: "volumeDevices is the list of block devices to be used by the container. +patchMergeKey=devicePath +patchStrategy=merge +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerVolumeDevices,
					},
					{
						Name:        "liveness_probe",
						Description: "Probes are not allowed for ephemeral containers. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerLivenessProbe,
					},
					{
						Name:        "readiness_probe",
						Description: "Probes are not allowed for ephemeral containers. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerReadinessProbe,
					},
					{
						Name:        "startup_probe",
						Description: "Probes are not allowed for ephemeral containers. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerStartupProbe,
					},
					{
						Name:        "lifecycle",
						Description: "Lifecycle is not allowed for ephemeral containers. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerLifecycle,
					},
					{
						Name:        "termination_message_path",
						Description: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.TerminationMessagePath"),
					},
					{
						Name:        "termination_message_policy",
						Description: "Indicate how the termination message should be populated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.TerminationMessagePolicy"),
					},
					{
						Name:        "image_pull_policy",
						Description: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.ImagePullPolicy"),
					},
					{
						Name:        "security_context",
						Description: "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecEphemeralContainerSecurityContext,
					},
					{
						Name:        "stdin",
						Description: "Whether this container should allocate a buffer for stdin in the container runtime",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.Stdin"),
					},
					{
						Name:        "stdin_once",
						Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EphemeralContainerCommon.StdinOnce"),
					},
					{
						Name:        "tty",
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
						Name:        "k8s_apps_stateful_set_template_ephemeral_container_ports",
						Description: "ContainerPort represents a network port in a single container.",
						Resolver:    fetchAppsStatefulSetTemplateSpecEphemeralContainerPorts,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_template_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_template_ephemeral_containers table (FK)",
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
						Name:        "k8s_apps_stateful_set_template_ephemeral_container_envs",
						Description: "EnvVar represents an environment variable present in a Container.",
						Resolver:    fetchAppsStatefulSetTemplateSpecEphemeralContainerEnvs,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_template_ephemeral_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_template_ephemeral_containers table (FK)",
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
				},
			},
			{
				Name:        "k8s_apps_stateful_set_template_image_pull_secrets",
				Description: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace. +structType=atomic",
				Resolver:    fetchAppsStatefulSetTemplateSpecImagePullSecrets,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_template_tolerations",
				Description: "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.",
				Resolver:    fetchAppsStatefulSetTemplateSpecTolerations,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_template_host_aliases",
				Description: "HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.",
				Resolver:    fetchAppsStatefulSetTemplateSpecHostAliases,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_volume_claim_templates",
				Description: "PersistentVolumeClaim is a user's request for and claim to a persistent volume",
				Resolver:    fetchAppsStatefulSetVolumeClaimTemplates,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "type_meta_kind",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TypeMeta.Kind"),
					},
					{
						Name:     "type_meta_api_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TypeMeta.APIVersion"),
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.Name"),
					},
					{
						Name:     "generate_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.GenerateName"),
					},
					{
						Name:     "namespace",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.Namespace"),
					},
					{
						Name:     "self_link",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.SelfLink"),
					},
					{
						Name:     "uid",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.UID"),
					},
					{
						Name:     "resource_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.ResourceVersion"),
					},
					{
						Name:     "generation",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ObjectMeta.Generation"),
					},
					{
						Name:     "deletion_grace_period_seconds",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
					},
					{
						Name:     "labels",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("ObjectMeta.Labels"),
					},
					{
						Name:     "annotations",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("ObjectMeta.Annotations"),
					},
					{
						Name:     "finalizers",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("ObjectMeta.Finalizers"),
					},
					{
						Name:     "cluster_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectMeta.ClusterName"),
					},
					{
						Name:        "access_modes",
						Description: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Spec.AccessModes"),
					},
					{
						Name:     "selector_match_labels",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Spec.Selector.MatchLabels"),
					},
					{
						Name:     "selector_match_expressions",
						Type:     schema.TypeJSON,
						Resolver: resolveAppsStatefulSetVolumeClaimTemplateSelectorMatchExpressions,
					},
					{
						Name:        "resources_limits",
						Description: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Spec.Resources.Limits"),
					},
					{
						Name:        "resources_requests",
						Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Spec.Resources.Requests"),
					},
					{
						Name:        "volume_name",
						Description: "VolumeName is the binding reference to the PersistentVolume backing this claim. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.VolumeName"),
					},
					{
						Name:        "storage_class_name",
						Description: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.StorageClassName"),
					},
					{
						Name:        "volume_mode",
						Description: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.VolumeMode"),
					},
					{
						Name:        "data_source_api_group",
						Description: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.DataSource.APIGroup"),
					},
					{
						Name:        "data_source_kind",
						Description: "Kind is the type of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.DataSource.Kind"),
					},
					{
						Name:        "data_source_name",
						Description: "Name is the name of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.DataSource.Name"),
					},
					{
						Name:        "data_source_ref_api_group",
						Description: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.DataSourceRef.APIGroup"),
					},
					{
						Name:        "data_source_ref_kind",
						Description: "Kind is the type of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.DataSourceRef.Kind"),
					},
					{
						Name:        "data_source_ref_name",
						Description: "Name is the name of resource being referenced",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.DataSourceRef.Name"),
					},
					{
						Name:        "status_phase",
						Description: "Phase represents the current phase of PersistentVolumeClaim. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.Phase"),
					},
					{
						Name:        "status_access_modes",
						Description: "AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Status.AccessModes"),
					},
					{
						Name:        "status_capacity",
						Description: "Represents the actual resources of the underlying volume. +optional",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Status.Capacity"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_apps_stateful_set_volume_claim_template_owner_references",
						Description: "OwnerReference contains enough information to let you identify an owning object",
						Resolver:    fetchAppsStatefulSetVolumeClaimTemplateOwnerReferences,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_volume_claim_template_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_volume_claim_templates table (FK)",
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
						Name:        "k8s_apps_stateful_set_volume_claim_template_managed_fields",
						Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
						Resolver:    fetchAppsStatefulSetVolumeClaimTemplateManagedFields,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_volume_claim_template_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_volume_claim_templates table (FK)",
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
						Name:        "k8s_apps_stateful_set_volume_claim_template_status_conditions",
						Description: "PersistentVolumeClaimCondition contails details about state of pvc",
						Resolver:    fetchAppsStatefulSetVolumeClaimTemplateStatusConditions,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_volume_claim_template_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_volume_claim_templates table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
							{
								Name: "status",
								Type: schema.TypeString,
							},
							{
								Name:        "reason",
								Description: "Unique, this should be a short, machine understandable string that gives the reason for condition's last transition",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "Human-readable message indicating details about last transition. +optional",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "k8s_apps_stateful_set_status_conditions",
				Description: "StatefulSetCondition describes the state of a statefulset at a certain point.",
				Resolver:    fetchAppsStatefulSetStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of statefulset condition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown.",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "The reason for the condition's last transition. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "A human readable message indicating details about the transition. +optional",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "k8s_apps_stateful_set_template_init_containers",
				Description: "A single application container that you want to run within a pod.",
				Resolver:    fetchAppsStatefulSetTemplateSpecInitContainers,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
						Name:        "env_from",
						Description: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerEnvFrom,
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
						Name:        "volume_mounts",
						Description: "Pod volumes to mount into the container's filesystem. Cannot be updated. +optional +patchMergeKey=mountPath +patchStrategy=merge",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerVolumeMounts,
					},
					{
						Name:        "volume_devices",
						Description: "volumeDevices is the list of block devices to be used by the container. +patchMergeKey=devicePath +patchStrategy=merge +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerVolumeDevices,
					},
					{
						Name:        "liveness_probe",
						Description: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerLivenessProbe,
					},
					{
						Name:        "readiness_probe",
						Description: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerReadinessProbe,
					},
					{
						Name:        "startup_probe",
						Description: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerStartupProbe,
					},
					{
						Name:        "lifecycle",
						Description: "Actions that the management system should take in response to container lifecycle events. Cannot be updated. +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerLifecycle,
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
						Name:        "security_context",
						Description: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ +optional",
						Type:        schema.TypeJSON,
						Resolver:    resolveAppsStatefulSetTemplateSpecInitContainerSecurityContext,
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
						Name:        "k8s_apps_stateful_set_template_init_container_ports",
						Description: "ContainerPort represents a network port in a single container.",
						Resolver:    fetchAppsStatefulSetTemplateSpecInitContainerPorts,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_template_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_template_init_containers table (FK)",
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
						Name:        "k8s_apps_stateful_set_template_init_container_envs",
						Description: "EnvVar represents an environment variable present in a Container.",
						Resolver:    fetchAppsStatefulSetTemplateSpecInitContainerEnvs,
						Columns: []schema.Column{
							{
								Name:        "stateful_set_template_init_container_cq_id",
								Description: "Unique CloudQuery ID of k8s_apps_stateful_set_template_init_containers table (FK)",
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
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAppsStatefulSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services.StatefulSets
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
func resolveAppsStatefulSetTemplateSpecSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
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
func resolveAppsStatefulSetTemplateSpecAffinity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
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
func resolveAppsStatefulSetTemplateSpecDNSConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
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
func resolveAppsStatefulSetTemplateSpecReadinessGates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template.Spec.ReadinessGates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsStatefulSetTemplateSpecOverhead(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template.Spec.Overhead)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsStatefulSetTemplateSpecTopologySpreadConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template.Spec.TopologySpreadConstraints)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchAppsStatefulSetOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.OwnerReferences
	return nil
}
func fetchAppsStatefulSetManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.ManagedFields
	return nil
}
func fetchAppsStatefulSetSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	if p.Spec.Selector == nil {
		return nil
	}
	res <- p.Spec.Selector.MatchExpressions
	return nil
}
func fetchAppsStatefulSetTemplateOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.Template.OwnerReferences
	return nil
}
func fetchAppsStatefulSetTemplateManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.Template.ManagedFields
	return nil
}
func fetchAppsStatefulSetTemplateSpecVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.Volumes
	return nil
}
func resolveAppsStatefulSetTemplateSpecVolumeAWSElasticBlockStore(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecVolumeNfs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Volume)
	if !ok {
		return fmt.Errorf("not a corev1.Volume instance: %T", resource.Item)
	}

	if p.NFS == nil {
		return nil
	}

	b, err := json.Marshal(p.NFS)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsStatefulSetTemplateSpecVolumeIscsi(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecVolumeRbd(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecVolumeDownwardAPI(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecVolumeStorageOs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecVolumeCsi(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func fetchAppsStatefulSetTemplateSpecContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.Containers
	return nil
}
func resolveAppsStatefulSetTemplateSpecContainerEnvFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerLivenessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerReadinessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerStartupProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerLifecycle(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecContainerSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func fetchAppsStatefulSetTemplateSpecContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchAppsStatefulSetTemplateSpecContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Env
	return nil
}
func fetchAppsStatefulSetTemplateSpecEphemeralContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.EphemeralContainers
	return nil
}
func resolveAppsStatefulSetTemplateSpecEphemeralContainerEnvFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerLivenessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerReadinessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerStartupProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerLifecycle(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecEphemeralContainerSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func fetchAppsStatefulSetTemplateSpecEphemeralContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchAppsStatefulSetTemplateSpecEphemeralContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.EphemeralContainer)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}
	res <- p.Env
	return nil
}
func fetchAppsStatefulSetTemplateSpecImagePullSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}

	res <- p.Spec.Template.Spec.ImagePullSecrets
	return nil
}
func fetchAppsStatefulSetTemplateSpecTolerations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}

	res <- p.Spec.Template.Spec.Tolerations
	return nil
}
func fetchAppsStatefulSetTemplateSpecHostAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}

	res <- p.Spec.Template.Spec.HostAliases
	return nil
}
func fetchAppsStatefulSetVolumeClaimTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.VolumeClaimTemplates
	return nil
}
func resolveAppsStatefulSetVolumeClaimTemplateSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.PersistentVolumeClaim)
	if !ok {
		return fmt.Errorf("not a corev1.PersistentVolumeClaim instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.Spec.Selector.MatchExpressions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)

}
func fetchAppsStatefulSetVolumeClaimTemplateOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.PersistentVolumeClaim)
	if !ok {
		return fmt.Errorf("not a corev1.PersistentVolumeClaim instance: %T", parent.Item)
	}

	res <- p.OwnerReferences
	return nil
}
func fetchAppsStatefulSetVolumeClaimTemplateManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.PersistentVolumeClaim)
	if !ok {
		return fmt.Errorf("not a corev1.PersistentVolumeClaim instance: %T", parent.Item)
	}

	res <- p.ManagedFields
	return nil
}
func fetchAppsStatefulSetVolumeClaimTemplateStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.PersistentVolumeClaim)
	if !ok {
		return fmt.Errorf("not a corev1.PersistentVolumeClaim instance: %T", parent.Item)
	}

	res <- p.Status.Conditions
	return nil
}
func fetchAppsStatefulSetStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}

	res <- p.Status.Conditions
	return nil
}
func fetchAppsStatefulSetTemplateSpecInitContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	res <- p.Spec.Template.Spec.InitContainers
	return nil
}
func resolveAppsStatefulSetTemplateSpecInitContainerEnvFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerVolumeDevices(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerLivenessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerReadinessProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerStartupProbe(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerLifecycle(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveAppsStatefulSetTemplateSpecInitContainerSecurityContext(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func fetchAppsStatefulSetTemplateSpecInitContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}

	res <- p.Ports
	return nil
}
func fetchAppsStatefulSetTemplateSpecInitContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(corev1.Container)
	if !ok {
		return fmt.Errorf("not a corev1.Container instance: %T", parent.Item)
	}

	res <- p.Env
	return nil
}
