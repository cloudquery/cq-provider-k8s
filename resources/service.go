package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Service() *schema.Table {
	return &schema.Table{
		Name:     "k8s_services",
		Resolver: fetchServices,
		Columns: append(objectCommonColumns(),
			schema.Column{
				Name:     "spec_selector",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Selector"),
			},
			schema.Column{
				Name:     "spec_cluster_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ClusterIP"),
			},
			schema.Column{
				Name:     "spec_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Type"),
			},
			schema.Column{
				Name:     "spec_external_ips",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.ExternalIPs"),
			},
			schema.Column{
				Name:     "spec_session_affinity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.SessionAffinity"),
			},
			schema.Column{
				Name:     "spec_load_balancer_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.LoadBalancerIP"),
			},
			schema.Column{
				Name:     "spec_load_balancer_source_ranges",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.LoadBalancerSourceRanges"),
			},
			schema.Column{
				Name:     "spec_external_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ExternalName"),
			},
			schema.Column{
				Name:     "spec_external_traffic_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ExternalTrafficPolicy"),
			},
			schema.Column{
				Name:     "spec_health_check_node_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.HealthCheckNodePort"),
			},
			schema.Column{
				Name:     "spec_publish_not_ready_addresses",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.PublishNotReadyAddresses"),
			},
			schema.Column{
				Name:     "spec_session_affinity_config_client_ip_timeout_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds"),
			},
			schema.Column{
				Name:     "spec_ip_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.IPFamily"),
			},
			schema.Column{
				Name:     "spec_topology_keys",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.TopologyKeys"),
			},
		),
		Relations: []*schema.Table{
			{
				Name:     "k8s_service_spec_ports",
				Resolver: fetchServicePorts,
				Columns: []schema.Column{
					{
						Name:     "service_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "protocol",
						Type: schema.TypeString,
					},
					{
						Name: "app_protocol",
						Type: schema.TypeString,
					},
					{
						Name: "port",
						Type: schema.TypeInt,
					},
					{
						Name:     "target_port",
						Type:     schema.TypeString,
						Resolver: resolveServicePortTargetPort,
					},
					{
						Name: "node_port",
						Type: schema.TypeInt,
					},
				},
			},
			{
				Name:     "k8s_service_status_load_balancer_ingress",
				Resolver: fetchServiceStatus,
				Columns: []schema.Column{
					{
						Name:     "service_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "IP",
						Type: schema.TypeString,
					},
					{
						Name: "Hostname",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	services, err := c.KClient.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Println(services.Items)
	res <- services.Items
	return nil
}

func fetchServicePorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(v1.Service)
	ports := service.Spec.Ports
	res <- ports
	return nil
}

func fetchServiceStatus(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(v1.Service)
	status := service.Status
	res <- status.LoadBalancer.Ingress
	return nil
}

func resolveServicePortTargetPort(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	port := resource.Item.(v1.ServicePort)
	target_port := ""

	if port.TargetPort.Type == intstr.String {
		target_port = port.TargetPort.StrVal
	} else {
		target_port = fmt.Sprintf("%d", port.TargetPort.IntVal)
	}
	return resource.Set(c.Name, target_port)
}
