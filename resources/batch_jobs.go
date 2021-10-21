package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BatchJobs() *schema.Table {
	return &schema.Table{
		Name:         "k8s_batch_jobs",
		Description:  "Job represents the configuration of a single job.",
		Resolver:     fetchBatchJobs,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
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
				Name:        "parallelism",
				Description: "Specifies the maximum desired number of pods the job should run at any given time",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Parallelism"),
			},
			{
				Name:        "completions",
				Description: "Specifies the desired number of successfully finished pods the job should be run with",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Completions"),
			},
			{
				Name:        "active_deadline_seconds",
				Description: "Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.ActiveDeadlineSeconds"),
			},
			{
				Name:        "backoff_limit",
				Description: "Specifies the number of retries before marking this job failed. Defaults to 6 +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.BackoffLimit"),
			},
			{
				Name:     "selector_match_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:        "manual_selector",
				Description: "manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.ManualSelector"),
			},
			{
				Name:     "template_object_meta_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Name"),
			},
			{
				Name:     "template_object_meta_generate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.GenerateName"),
			},
			{
				Name:     "template_object_meta_namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Namespace"),
			},
			{
				Name:     "template_object_meta_self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.SelfLink"),
			},
			{
				Name:     "template_object_meta_uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.UID"),
			},
			{
				Name:     "template_object_meta_resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.ResourceVersion"),
			},
			{
				Name:     "template_object_meta_generation",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Generation"),
			},
			{
				Name:     "template_object_meta_deletion_grace_period_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:     "template_object_meta_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Labels"),
			},
			{
				Name:     "template_object_meta_annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Annotations"),
			},
			{
				Name:     "template_object_meta_finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.Finalizers"),
			},
			{
				Name:     "template_object_meta_cluster_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Template.ObjectMeta.ClusterName"),
			},
			{
				Name:     "template_spec",
				Type:     schema.TypeJSON,
				Resolver: resolveBatchJobTemplateSpec,
			},
			{
				Name:        "ttl_seconds_after_finished",
				Description: "ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.TTLSecondsAfterFinished"),
			},
			{
				Name:        "completion_mode",
				Description: "CompletionMode specifies how Pod completions are tracked",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.CompletionMode"),
			},
			{
				Name:        "suspend",
				Description: "Suspend specifies whether the Job controller should create Pods or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Suspend"),
			},
			{
				Name:        "status_active",
				Description: "The number of actively running pods. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Active"),
			},
			{
				Name:        "status_succeeded",
				Description: "The number of pods which reached phase Succeeded. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Succeeded"),
			},
			{
				Name:        "status_failed",
				Description: "The number of pods which reached phase Failed. +optional",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Failed"),
			},
			{
				Name:        "status_completed_indexes",
				Description: "CompletedIndexes holds the completed indexes when .spec.completionMode = \"Indexed\" in a text format",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.CompletedIndexes"),
			},
			{
				Name:        "status_uncounted_terminated_pods_succeeded",
				Description: "Succeeded holds UIDs of succeeded Pods. +listType=set +optional",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Status.UncountedTerminatedPods.Succeeded"),
			},
			{
				Name:        "status_uncounted_terminated_pods_failed",
				Description: "Failed holds UIDs of failed Pods. +listType=set +optional",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Status.UncountedTerminatedPods.Failed"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_batch_job_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchBatchJobOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
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
				Name:        "k8s_batch_job_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchBatchJobManagedFields,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
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
				Name:        "k8s_batch_job_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchBatchJobSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
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
				Name:        "k8s_batch_job_template_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchBatchJobTemplateOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
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
				Name:        "k8s_batch_job_template_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchBatchJobTemplateManagedFields,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
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
				Name:        "k8s_batch_job_status_conditions",
				Description: "JobCondition describes current state of a job.",
				Resolver:    fetchBatchJobStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of job condition, Complete or Failed.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown.",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "(brief) reason for the condition's last transition. +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "Human readable message indicating details about last transition. +optional",
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
func fetchBatchJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	jobs := meta.(*client.Client).Services.Jobs
	result, err := jobs.List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	res <- result.Items
	return nil
}
func resolveBatchJobTemplateSpec(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	job, ok := resource.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", resource.Item)
	}
	b, err := json.Marshal(job.Spec.Template.Spec)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchBatchJobOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	res <- job.OwnerReferences
	return nil
}
func fetchBatchJobManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	res <- job.ManagedFields
	return nil
}
func fetchBatchJobSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	if job.Spec.Selector == nil {
		return nil
	}
	res <- job.Spec.Selector.MatchExpressions
	return nil
}
func fetchBatchJobTemplateOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	res <- job.Spec.Template.OwnerReferences
	return nil
}
func fetchBatchJobTemplateManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	res <- job.Spec.Template.ManagedFields
	return nil
}
func fetchBatchJobStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	res <- job.Status.Conditions
	return nil
}
