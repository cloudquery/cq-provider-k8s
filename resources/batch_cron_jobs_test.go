package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createBatchCronJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	cronJobs := mocks.NewMockCronJobsClient(ctrl)

	cronJobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&batchv1.CronJobList{Items: []batchv1.CronJob{fakeCronJob(t)}},
		nil,
	)
	return client.Services{
		CronJobs: cronJobs,
	}
}

func fakeCronJob(t *testing.T) batchv1.CronJob {
	var job batchv1.CronJob
	fakeThroughPointers(t, []interface{}{
		&job.TypeMeta,
		&job.ObjectMeta,
		&job.Spec.Schedule,
		&job.Spec.StartingDeadlineSeconds,
		&job.Spec.ConcurrencyPolicy,
		&job.Spec.Suspend,
		&job.Spec.SuccessfulJobsHistoryLimit,
		&job.Spec.FailedJobsHistoryLimit,
		&job.Spec.JobTemplate.ObjectMeta,
		&job.Spec.JobTemplate.Spec.Parallelism,
		&job.Spec.JobTemplate.Spec.Completions,
		&job.Spec.JobTemplate.Spec.ActiveDeadlineSeconds,
		&job.Spec.JobTemplate.Spec.BackoffLimit,
		&job.Spec.JobTemplate.Spec.Selector,
		&job.Spec.JobTemplate.Spec.ManualSelector,
		&job.Spec.JobTemplate.Spec.TTLSecondsAfterFinished,
		&job.Spec.JobTemplate.Spec.CompletionMode,
		&job.Spec.JobTemplate.Spec.Suspend,
	})
	job.Spec.JobTemplate.Spec.Template = fakePodTemplateSpec(t)
	return job
}

func fakePodTemplateSpec(t *testing.T) corev1.PodTemplateSpec {
	var pod corev1.PodTemplateSpec
	pod.Spec.Volumes = []corev1.Volume{fakeVolume(t)}
	fakeThroughPointers(t, []interface{}{
		&pod.ObjectMeta,

		// &pod.Spec.InitContainers,
		// &pod.Spec.Containers,
		// &pod.Spec.EphemeralContainers,
		&pod.Spec.RestartPolicy,
		&pod.Spec.TerminationGracePeriodSeconds,
		&pod.Spec.ActiveDeadlineSeconds,
		&pod.Spec.DNSPolicy,
		&pod.Spec.NodeSelector,
		&pod.Spec.ServiceAccountName,
		&pod.Spec.AutomountServiceAccountToken,
		&pod.Spec.NodeName,
		&pod.Spec.HostNetwork,
		&pod.Spec.HostPID,
		&pod.Spec.HostIPC,
		&pod.Spec.ShareProcessNamespace,
		&pod.Spec.SecurityContext,
		&pod.Spec.ImagePullSecrets,
		&pod.Spec.Hostname,
		&pod.Spec.Subdomain,
		&pod.Spec.Affinity,
		&pod.Spec.SchedulerName,
		&pod.Spec.Tolerations,
		&pod.Spec.HostAliases,
		&pod.Spec.PriorityClassName,
		&pod.Spec.Priority,
		&pod.Spec.DNSConfig,
		&pod.Spec.ReadinessGates,
		&pod.Spec.RuntimeClassName,
		&pod.Spec.EnableServiceLinks,
		&pod.Spec.PreemptionPolicy,
		// &pod.Spec.Overhead,
		&pod.Spec.TopologySpreadConstraints,
		&pod.Spec.SetHostnameAsFQDN,
	})
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	pod.Spec.Overhead = rl

	pod.Spec.InitContainers = []corev1.Container{fakeContainer(t)}
	pod.Spec.Containers = []corev1.Container{fakeContainer(t)}
	pod.Spec.EphemeralContainers = []corev1.EphemeralContainer{fakeEphemeralContainer(t)}

	return pod
}

func TestBatchCronJobs(t *testing.T) {
	k8sTestHelper(t, BatchCronJobs(), createBatchCronJobs)
}
