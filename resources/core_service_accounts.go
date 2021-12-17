package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CoreServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_service_accounts",
		Description:  "ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets",
		Resolver:     fetchCoreServiceAccounts,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.APIVersion"),
			},
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
				Resolver:    resolveCoreServiceAccountsOwnerReferences,
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
				Resolver:    resolveCoreServiceAccountsManagedFields,
			},
			{
				Name:        "automount_service_account_token",
				Description: "AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level. +optional",
				Type:        schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_core_service_account_secrets",
				Description: "ObjectReference contains enough information to let you inspect or modify the referred object. --- New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.  1",
				Resolver:    fetchCoreServiceAccountSecrets,
				Columns: []schema.Column{
					{
						Name:        "service_account_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_service_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "namespace",
						Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "uid",
						Description: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UID"),
					},
					{
						Name:        "api_version",
						Description: "API version of the referent. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("APIVersion"),
					},
					{
						Name:        "resource_version",
						Description: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency +optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "field_path",
						Description: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: \"spec.containers{name}\" (where \"name\" refers to the name of the container that triggered the event) or if no container name is specified \"spec.containers[2]\" (container with index 2 in this pod)",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "k8s_core_service_account_image_pull_secrets",
				Description: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace. +structType=atomic",
				Resolver:    fetchCoreServiceAccountImagePullSecrets,
				Columns: []schema.Column{
					{
						Name:        "service_account_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_service_accounts table (FK)",
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
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCoreServiceAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client).Services().ServiceAccounts
	opts := metav1.ListOptions{}
	for {
		result, err := c.List(ctx, opts)
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

func resolveCoreServiceAccountsOwnerReferences(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.ServiceAccount)
	if !ok {
		return fmt.Errorf("not a corev1.ServiceAccount instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreServiceAccountsManagedFields(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.ServiceAccount)
	if !ok {
		return fmt.Errorf("not a corev1.ServiceAccount instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchCoreServiceAccountSecrets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	serviceAccount, ok := parent.Item.(corev1.ServiceAccount)
	if !ok {
		return fmt.Errorf("not a corev1.ServiceAccount instance: %T", parent.Item)
	}
	res <- serviceAccount.Secrets
	return nil
}

func fetchCoreServiceAccountImagePullSecrets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	serviceAccount, ok := parent.Item.(corev1.ServiceAccount)
	if !ok {
		return fmt.Errorf("not a corev1.ServiceAccount instance: %T", parent.Item)
	}
	res <- serviceAccount.ImagePullSecrets
	return nil
}
