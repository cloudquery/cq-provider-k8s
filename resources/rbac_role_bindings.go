package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RbacRoleBindings() *schema.Table {
	return &schema.Table{
		Name:        "k8s_rbac_role_bindings",
		Description: "RoleBinding references a role, but does not contain it",
		Resolver:    fetchRbacRoleBindings,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:     "api_version",
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
				Name:        "role_ref_api_group",
				Description: "APIGroup is the group for the resource being referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleRef.APIGroup"),
			},
			{
				Name:        "role_ref_kind",
				Description: "Kind is the type of resource being referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleRef.Kind"),
			},
			{
				Name:        "role_ref_name",
				Description: "Name is the name of resource being referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleRef.Name"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_rbac_role_binding_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchRbacRoleBindingOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "role_binding_cq_id",
						Description: "Unique CloudQuery ID of k8s_rbac_role_bindings table (FK)",
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
				Name:        "k8s_rbac_role_binding_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchRbacRoleBindingManagedFields,
				Columns: []schema.Column{
					{
						Name:        "role_binding_cq_id",
						Description: "Unique CloudQuery ID of k8s_rbac_role_bindings table (FK)",
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
				Name:        "k8s_rbac_role_binding_subjects",
				Description: "Subject contains a reference to the object or user identities a role binding applies to",
				Resolver:    fetchRbacRoleBindingSubjects,
				Columns: []schema.Column{
					{
						Name:        "role_binding_cq_id",
						Description: "Unique CloudQuery ID of k8s_rbac_role_bindings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Kind of object being referenced",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_group",
						Description: "APIGroup holds the API group of the referenced subject. Defaults to \"\" for ServiceAccount subjects. Defaults to \"rbac.authorization.k8s.io\" for User and Group subjects. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("APIGroup"),
					},
					{
						Name:        "name",
						Description: "Name of the object being referenced.",
						Type:        schema.TypeString,
					},
					{
						Name:        "namespace",
						Description: "Namespace of the referenced object",
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
func fetchRbacRoleBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services.RoleBindings
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
	return nil
}
func fetchRbacRoleBindingOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	role, ok := parent.Item.(rbacv1.RoleBinding)
	if !ok {
		return fmt.Errorf("not a rbacv1.RoleBinding instance: %T", parent.Item)
	}
	res <- role.OwnerReferences
	return nil
}
func fetchRbacRoleBindingManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	role, ok := parent.Item.(rbacv1.RoleBinding)
	if !ok {
		return fmt.Errorf("not a rbacv1.RoleBinding instance: %T", parent.Item)
	}
	res <- role.ManagedFields
	return nil
}
func fetchRbacRoleBindingSubjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	role, ok := parent.Item.(rbacv1.RoleBinding)
	if !ok {
		return fmt.Errorf("not a rbacv1.RoleBinding instance: %T", parent.Item)
	}
	res <- role.Subjects
	return nil
}
