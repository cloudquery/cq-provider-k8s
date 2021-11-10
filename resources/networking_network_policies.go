package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NetworkingNetworkPolicies() *schema.Table {
	return &schema.Table{
		Name:        "k8s_networking_network_policies",
		Description: "NetworkPolicy describes what network traffic is allowed for a set of Pods",
		Resolver:    fetchNetworkingNetworkPolicies,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.",
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
				Name:        "finalizers",
				Description: "Must be empty before the object is deleted from the registry",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "pod_selector_match_labels",
				Description: "matchLabels is a map of {key,value} pairs",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.PodSelector.MatchLabels"),
			},
			{
				Name:        "policy_types",
				Description: "List of rule types that the NetworkPolicy relates to. Valid options are [\"Ingress\"], [\"Egress\"], or [\"Ingress\", \"Egress\"]. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ \"Egress\" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include \"Egress\" (since such a policy would not include an Egress section and would otherwise default to just [ \"Ingress\" ]). This field is beta-level in 1.8",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.PolicyTypes"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_networking_network_policy_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchNetworkingNetworkPolicyOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
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
						Description: "If true, this reference points to the managing controller.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "block_owner_deletion",
						Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "k8s_networking_network_policy_managed_fields",
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Resolver:    fetchNetworkingNetworkPolicyManagedFields,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
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
				Name:        "k8s_networking_network_policy_pod_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchNetworkingNetworkPolicyPodSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
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
				Name:        "k8s_networking_network_policy_ingress",
				Description: "NetworkPolicyIngressRule describes a particular set of traffic that is allowed to the pods matched by a NetworkPolicySpec's podSelector",
				Resolver:    fetchNetworkingNetworkPolicyIngresses,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_networking_network_policy_ingress_ports",
						Description: "NetworkPolicyPort describes a port to allow traffic on",
						Resolver:    fetchNetworkingNetworkPolicyIngressPorts,
						Columns: []schema.Column{
							{
								Name:        "network_policy_ingress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_ingress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "protocol",
								Description: "The protocol (TCP, UDP, or SCTP) which traffic must match",
								Type:        schema.TypeString,
							},
							{
								Name:     "port_type",
								Type:     schema.TypeBigInt,
								Resolver: schema.PathResolver("Port.Type"),
							},
							{
								Name:     "port_int_val",
								Type:     schema.TypeInt,
								Resolver: schema.PathResolver("Port.IntVal"),
							},
							{
								Name:     "port_str_val",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Port.StrVal"),
							},
							{
								Name:        "end_port",
								Description: "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy",
								Type:        schema.TypeInt,
							},
						},
					},
					{
						Name:        "k8s_networking_network_policy_ingress_from",
						Description: "NetworkPolicyPeer describes a peer to allow traffic to/from",
						Resolver:    fetchNetworkingNetworkPolicyIngressFroms,
						Columns: []schema.Column{
							{
								Name:        "network_policy_ingress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_ingress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "pod_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("PodSelector.MatchLabels"),
							},
							{
								Name:        "namespace_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("NamespaceSelector.MatchLabels"),
							},
							{
								Name:        "ip_block_c_id_r",
								Description: "CIDR is a string representing the IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("IPBlock.CIDR"),
							},
							{
								Name:        "ip_block_except",
								Description: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\" Except values will be rejected if they are outside the CIDR range",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("IPBlock.Except"),
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "k8s_networking_network_policy_ingress_from_pod_selector_match_expressions",
								Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
								Resolver:    fetchNetworkingNetworkPolicyIngressFromPodSelectorMatchExpressions,
								Columns: []schema.Column{
									{
										Name:        "network_policy_ingress_from_cq_id",
										Description: "Unique CloudQuery ID of k8s_networking_network_policy_ingress_from table (FK)",
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
								Name:        "k8s_networking_network_policy_ingress_from_namespace_selector_match_expressions",
								Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
								Resolver:    fetchNetworkingNetworkPolicyIngressFromNamespaceSelectorMatchExpressions,
								Columns: []schema.Column{
									{
										Name:        "network_policy_ingress_from_cq_id",
										Description: "Unique CloudQuery ID of k8s_networking_network_policy_ingress_from table (FK)",
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
				},
			},
			{
				Name:        "k8s_networking_network_policy_egress",
				Description: "NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched by a NetworkPolicySpec's podSelector",
				Resolver:    fetchNetworkingNetworkPolicyEgresses,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_networking_network_policy_egress_ports",
						Description: "NetworkPolicyPort describes a port to allow traffic on",
						Resolver:    fetchNetworkingNetworkPolicyEgressPorts,
						Columns: []schema.Column{
							{
								Name:        "network_policy_egress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_egress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "protocol",
								Description: "The protocol (TCP, UDP, or SCTP) which traffic must match",
								Type:        schema.TypeString,
							},
							{
								Name:     "port_type",
								Type:     schema.TypeBigInt,
								Resolver: schema.PathResolver("Port.Type"),
							},
							{
								Name:     "port_int_val",
								Type:     schema.TypeInt,
								Resolver: schema.PathResolver("Port.IntVal"),
							},
							{
								Name:     "port_str_val",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Port.StrVal"),
							},
							{
								Name:        "end_port",
								Description: "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy",
								Type:        schema.TypeInt,
							},
						},
					},
					{
						Name:        "k8s_networking_network_policy_egress_to",
						Description: "NetworkPolicyPeer describes a peer to allow traffic to/from",
						Resolver:    fetchNetworkingNetworkPolicyEgressTos,
						Columns: []schema.Column{
							{
								Name:        "network_policy_egress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_egress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "pod_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("PodSelector.MatchLabels"),
							},
							{
								Name:        "namespace_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("NamespaceSelector.MatchLabels"),
							},
							{
								Name:        "ip_block_c_id_r",
								Description: "CIDR is a string representing the IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("IPBlock.CIDR"),
							},
							{
								Name:        "ip_block_except",
								Description: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\" Except values will be rejected if they are outside the CIDR range",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("IPBlock.Except"),
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "k8s_networking_network_policy_egress_to_pod_selector_match_expressions",
								Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
								Resolver:    fetchNetworkingNetworkPolicyEgressToPodSelectorMatchExpressions,
								Columns: []schema.Column{
									{
										Name:        "network_policy_egress_to_cq_id",
										Description: "Unique CloudQuery ID of k8s_networking_network_policy_egress_to table (FK)",
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
								Name:        "k8s_networking_network_policy_egress_to_namespace_selector_match_expressions",
								Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
								Resolver:    fetchNetworkingNetworkPolicyEgressToNamespaceSelectorMatchExpressions,
								Columns: []schema.Column{
									{
										Name:        "network_policy_egress_to_cq_id",
										Description: "Unique CloudQuery ID of k8s_networking_network_policy_egress_to table (FK)",
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
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkingNetworkPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services.NetworkPolicies
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
func fetchNetworkingNetworkPolicyOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.OwnerReferences
	return nil
}
func fetchNetworkingNetworkPolicyManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.ManagedFields
	return nil
}
func fetchNetworkingNetworkPolicyPodSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.Spec.PodSelector.MatchExpressions
	return nil
}
func fetchNetworkingNetworkPolicyIngresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.Spec.Ingress
	return nil
}
func fetchNetworkingNetworkPolicyIngressPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyIngressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchNetworkingNetworkPolicyIngressFroms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyIngressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.From
	return nil
}
func fetchNetworkingNetworkPolicyIngressFromPodSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", parent.Item)
	}
	if p.PodSelector == nil {
		return nil
	}
	res <- p.PodSelector.MatchExpressions
	return nil
}
func fetchNetworkingNetworkPolicyIngressFromNamespaceSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", parent.Item)
	}
	if p.NamespaceSelector == nil {
		return nil
	}
	res <- p.NamespaceSelector.MatchExpressions
	return nil
}
func fetchNetworkingNetworkPolicyEgresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.Spec.Egress
	return nil
}
func fetchNetworkingNetworkPolicyEgressPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyEgressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchNetworkingNetworkPolicyEgressTos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyEgressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.To
	return nil
}
func fetchNetworkingNetworkPolicyEgressToPodSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", parent.Item)
	}
	if p.PodSelector == nil {
		return nil
	}
	res <- p.PodSelector.MatchExpressions
	return nil
}
func fetchNetworkingNetworkPolicyEgressToNamespaceSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", parent.Item)
	}
	if p.NamespaceSelector == nil {
		return nil
	}
	res <- p.NamespaceSelector.MatchExpressions
	return nil
}
