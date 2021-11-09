package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
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

func TestBatchCronJobs(t *testing.T) {
	k8sTestHelper(t, BatchCronJobs(), createBatchCronJobs)
}
