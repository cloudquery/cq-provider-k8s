package resources

import (
	"github.com/cloudquery/faker/v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func fakeThroughPointers(t *testing.T, ptrs []interface{}) {
	for i, ptr := range ptrs {
		if err := faker.FakeData(ptr); err != nil {
			t.Fatalf("%v %v", i, ptr)
		}
	}
}

//nolint
func fakeDaemonSet(t *testing.T) appsv1.DaemonSet {
	var ds appsv1.DaemonSet
	ds.Spec.Template.Spec.Volumes = []corev1.Volume{fakeVolume(t)}
	fakeThroughPointers(t, []interface{}{
		&ds.TypeMeta,
		&ds.ObjectMeta,
		&ds.Status,
		&ds.ManagedFields,
		&ds.Annotations,
		&ds.Labels,
		&ds.OwnerReferences,
		&ds.Status,
		&ds.Spec.Selector,
		&ds.Spec.RevisionHistoryLimit,
	})

	ds.Spec.Template = fakePodTemplateSpec(t)
	return ds
}

func fakePodTemplateSpec(t *testing.T) corev1.PodTemplateSpec {
	var templateSpec corev1.PodTemplateSpec
	if err := faker.FakeDataSkipFields(&templateSpec, []string{"Spec"}); err != nil {
		t.Fatal(err)
	}
	templateSpec.Spec = fakePodSpec(t)
	return templateSpec
}

func fakeNode(t *testing.T) corev1.Node {
	// faker chokes on Node.Status.{Capacity,Allocatable} so doing it by hand
	var node corev1.Node
	ptrs := []interface{}{
		&node.TypeMeta,
		&node.ObjectMeta,
		&node.Spec,
		&node.Status.Phase,
		&node.Status.Conditions,
		&node.Status.Addresses,
		&node.Status.DaemonEndpoints,
		&node.Status.NodeInfo,
		&node.Status.Images,
		&node.Status.VolumesInUse,
		&node.Status.VolumesAttached,
		&node.Status.Config,
	}
	for i, ptr := range ptrs {
		if err := faker.FakeData(ptr); err != nil {
			t.Fatalf("%v %v", i, ptr)
		}
	}
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	node.Status.Capacity = rl
	node.Status.Allocatable = rl
	node.Spec.PodCIDR = "192.168.1.0/24"
	node.Spec.PodCIDRs = []string{"192.168.1.0/24"}
	node.Status.Addresses = []corev1.NodeAddress{
		{
			Type:    corev1.NodeHostName,
			Address: "testname",
		},
		{
			Type:    corev1.NodeInternalIP,
			Address: "fd00::1",
		},
		{
			Type:    corev1.NodeExternalIP,
			Address: "192.168.2.1",
		},
	}
	return node
}

func fakeVolume(t *testing.T) corev1.Volume {
	// faker chokes on volume.VolumeSource.Ephemeral
	var volume corev1.Volume
	fakeThroughPointers(t, []interface{}{
		&volume.Name,
		&volume.VolumeSource.HostPath,
		&volume.VolumeSource.EmptyDir,
		&volume.VolumeSource.GCEPersistentDisk,
		&volume.VolumeSource.AWSElasticBlockStore,
		&volume.VolumeSource.GitRepo,
		&volume.VolumeSource.Secret,
		&volume.VolumeSource.NFS,
		&volume.VolumeSource.ISCSI,
		&volume.VolumeSource.Glusterfs,
		&volume.VolumeSource.PersistentVolumeClaim,
		&volume.VolumeSource.RBD,
		&volume.VolumeSource.FlexVolume,
		&volume.VolumeSource.Cinder,
		&volume.VolumeSource.CephFS,
		&volume.VolumeSource.Flocker,
		&volume.VolumeSource.DownwardAPI,
		&volume.VolumeSource.FC,
		&volume.VolumeSource.AzureFile,
		&volume.VolumeSource.ConfigMap,
		&volume.VolumeSource.VsphereVolume,
		&volume.VolumeSource.Quobyte,
		&volume.VolumeSource.AzureDisk,
		&volume.VolumeSource.PhotonPersistentDisk,
		&volume.VolumeSource.Projected,
		&volume.VolumeSource.PortworxVolume,
		&volume.VolumeSource.ScaleIO,
		&volume.VolumeSource.StorageOS,
		&volume.VolumeSource.CSI,
		// &volume.VolumeSource.Ephemeral,
	})
	volume.Ephemeral = &corev1.EphemeralVolumeSource{}
	return volume
}

func fakeContainer(t *testing.T) corev1.Container {
	var c corev1.Container
	fakeThroughPointers(t, []interface{}{
		&c.Name,
		&c.Image,
		&c.Command,
		&c.Args,
		&c.WorkingDir,
		&c.Ports,
		&c.EnvFrom,
		&c.Env,
		// &c.Resources,
		&c.VolumeMounts,
		&c.VolumeDevices,
		// &c.LivenessProbe,
		// &c.ReadinessProbe,
		// &c.StartupProbe,
		// &c.Lifecycle,
		&c.TerminationMessagePath,
		&c.TerminationMessagePolicy,
		&c.ImagePullPolicy,
		&c.SecurityContext,
	})
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	c.Resources.Limits = rl
	c.Resources.Requests = rl
	c.LivenessProbe = &corev1.Probe{}
	c.ReadinessProbe = &corev1.Probe{}
	c.StartupProbe = &corev1.Probe{}
	c.Lifecycle = &corev1.Lifecycle{}
	return c
}

func fakeEphemeralContainer(t *testing.T) corev1.EphemeralContainer {
	var c corev1.EphemeralContainer
	fakeThroughPointers(t, []interface{}{
		&c.TargetContainerName,
		&c.Name,
		&c.Image,
		&c.Command,
		&c.Args,
		&c.WorkingDir,
		&c.Ports,
		&c.EnvFrom,
		&c.Env,
		// &c.Resources,
		&c.VolumeMounts,
		&c.VolumeDevices,
		// &c.LivenessProbe,
		// &c.ReadinessProbe,
		// &c.StartupProbe,
		// &c.Lifecycle,
		&c.TerminationMessagePath,
		&c.TerminationMessagePolicy,
		&c.ImagePullPolicy,
		&c.SecurityContext,
	})
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	c.Resources.Limits = rl
	c.Resources.Requests = rl
	c.LivenessProbe = &corev1.Probe{}
	c.ReadinessProbe = &corev1.Probe{}
	c.StartupProbe = &corev1.Probe{}
	c.Lifecycle = &corev1.Lifecycle{}
	return c
}

func fakePod(t *testing.T) corev1.Pod {
	var pod corev1.Pod
	pod.Spec.Volumes = []corev1.Volume{fakeVolume(t)}
	fakeThroughPointers(t, []interface{}{
		&pod.TypeMeta,
		&pod.ObjectMeta,
		&pod.Status,
	})
	pod.Spec = fakePodSpec(t)

	pod.Status.HostIP = "192.168.1.2"
	pod.Status.PodIP = "192.168.1.1"
	pod.Status.PodIPs = []corev1.PodIP{{IP: "192.168.1.1"}}
	return pod
}

func fakePodSpec(t *testing.T) corev1.PodSpec {
	var podSpec corev1.PodSpec
	if err := faker.FakeDataSkipFields(&podSpec, []string{
		"RestartPolicy",
		"Overhead",
		"InitContainers",
		"Containers",
		"EphemeralContainers",
		"Volumes",
		"DNSPolicy",
	}); err != nil {
		t.Fatal(err)
	}
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	podSpec.Overhead = rl
	podSpec.InitContainers = []corev1.Container{fakeContainer(t)}
	podSpec.Containers = []corev1.Container{fakeContainer(t)}
	podSpec.EphemeralContainers = []corev1.EphemeralContainer{fakeEphemeralContainer(t)}
	podSpec.Volumes = []corev1.Volume{fakeVolume(t)}
	podSpec.RestartPolicy = "test"

	return podSpec
}
