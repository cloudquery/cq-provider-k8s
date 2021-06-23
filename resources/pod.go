package resources

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Pod() *schema.Table {
	return &schema.Table{
		Name:     "k8s_pods",
		Resolver: fetchPods,
		Columns: append(
			objectCommonColumns(),
			schema.Column{
				Name: "spec_restart_policy",
				Type: schema.TypeString,
			},
		),
		Relations: []*schema.Table{
			{
				Name:     "k8s_pod_spec_volumes",
				Resolver: fetchPodVolumes,
				Columns: []schema.Column{
					{
						Name:     "pod_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "source",
						Type:     schema.TypeJSON,
						Resolver: resolveVolumeSource,
					},
				},
			},
			{
				Name:     "k8s_pod_spec_containers",
				Resolver: fetchPodContainers,
				Columns: []schema.Column{
					{
						Name:     "pod_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "image",
						Type: schema.TypeString,
					},
					{
						Name: "command",
						Type: schema.TypeStringArray,
					},
					{
						Name: "args",
						Type: schema.TypeStringArray,
					},
					{
						Name: "working_dir",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_limits",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("Resources.Limits"),
					},
					{
						Name:     "resource_requests",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("Resources.Requests"),
					},
					{
						Name:     "liveness_probe",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("LivenessProbe"),
					},
					{
						Name:     "readiness_probe",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("ReadinessProbe"),
					},
					{
						Name:     "startup_probe",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("StartupProbe"),
					},
					{
						Name:     "lifecicle_post_start",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("Lifecycle.PostStart"),
					},
					{
						Name:     "lifecicle_pre_stop",
						Type:     schema.TypeJSON,
						Resolver: PathToJSONResolver("Lifecycle.PreStart"),
					},
					{
						Name: "termination_message_path",
						Type: schema.TypeString,
					},
					{
						Name: "termination_message_policy",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "k8s_pod_spec_container_ports",
						Resolver: fetchPodContainerPorts,
						Columns: []schema.Column{
							{
								Name:     "container_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "name",
								Type: schema.TypeString,
							},
							{
								Name: "host_port",
								Type: schema.TypeInt,
							},
							{
								Name: "container_port",
								Type: schema.TypeInt,
							},
							{
								Name: "protocol",
								Type: schema.TypeString,
							},
							{
								Name: "host_ip",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "k8s_pod_spec_container_env_from_secrets",
						Resolver: fetchPodContainerEnvFromSecrets,
						Columns: []schema.Column{
							{
								Name:     "container_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "prefix",
								Type: schema.TypeString,
							},
							{
								Name:     "secret_name",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SecretRef.Name"),
							},
							{
								Name:     "optional",
								Type:     schema.TypeBool,
								Resolver: schema.PathResolver("SecretRef.Optional"),
							},
						},
					},
					{
						Name:     "k8s_pod_spec_container_env_from_config_maps",
						Resolver: fetchPodContainerEnvFromConfigMaps,
						Columns: []schema.Column{
							{
								Name:     "container_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "prefix",
								Type: schema.TypeString,
							},
							{
								Name:     "config_map_name",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ConfigMapRef.Name"),
							},
							{
								Name:     "optional",
								Type:     schema.TypeBool,
								Resolver: schema.PathResolver("ConfigMapRef.Optional"),
							},
						},
					},
					{
						Name:     "k8s_pod_spec_container_envs",
						Resolver: PathResolver("Env"),
						Columns: []schema.Column{
							{
								Name:     "container_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "name",
								Type: schema.TypeString,
							},
							{
								Name: "value",
								Type: schema.TypeString,
							},
							{
								Name:     "value_from",
								Type:     schema.TypeJSON,
								Resolver: PathToJSONResolver("ValueFrom"),
							},
						},
					},
					{
						Name:     "k8s_pod_spec_container_volume_mounts",
						Resolver: PathResolver("VolumeMounts"),
						Columns: []schema.Column{
							{
								Name:     "container_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "name",
								Type: schema.TypeString,
							},
							{
								Name: "read_only",
								Type: schema.TypeBool,
							},
							{
								Name: "mount_path",
								Type: schema.TypeString,
							},
							{
								Name: "sub_path",
								Type: schema.TypeString,
							},
							{
								Name: "mount_propagation",
								Type: schema.TypeString,
							},
							{
								Name: "sub_path_expr",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "k8s_pod_spec_container_volume_devices",
						Resolver: PathResolver("VolumeDevices"),
						Columns: []schema.Column{
							{
								Name:     "container_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "name",
								Type: schema.TypeString,
							},
							{
								Name: "device_path",
								Type: schema.TypeBool,
							},
						},
					},
				},
			},
		},
	}
}

func fetchPods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	pods, err := c.KClient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	res <- pods.Items
	return nil
}

func fetchPodVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	pod := parent.Item.(v1.Pod)
	res <- pod.Spec.Volumes
	return nil
}

func resolveVolumeSource(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	volume := resource.Item.(v1.Volume)
	json, err := json.Marshal(volume.VolumeSource)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, string(json))
}

func fetchPodContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	pod := parent.Item.(v1.Pod)
	res <- pod.Spec.Containers
	return nil
}

func fetchPodContainerPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	container := parent.Item.(v1.Container)
	res <- container.Ports
	return nil
}

func fetchPodContainerEnvFromSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	container := parent.Item.(v1.Container)
	for _, envFrom := range container.EnvFrom {
		if envFrom.SecretRef != nil {
			res <- envFrom
		}
	}
	return nil
}

func fetchPodContainerEnvFromConfigMaps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	container := parent.Item.(v1.Container)
	for _, envFrom := range container.EnvFrom {
		if envFrom.ConfigMapRef != nil {
			res <- envFrom
		}
	}
	return nil
}
