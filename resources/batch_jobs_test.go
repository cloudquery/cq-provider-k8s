package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func createBatchJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	jobs := mocks.NewMockJobsClient(ctrl)
	j := batchv1.Job{}
	fakeSkipFields(t, &j, []string{
		"Spec.Template.Spec.Overhead",
		"Spec.Template.Spec.DNSPolicy",
		"Spec.Template.Spec.Volumes",
		"Spec.Template.Spec.EphemeralContainers",
		"Spec.Template.Spec.Containers",
		"Spec.Template.Spec.InitContainers",
		"Spec.Template.Spec.RestartPolicy"})
	if err := faker.FakeData(&j.Spec.Template.Spec.RestartPolicy); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeData(&j.Spec.Template.Spec.TopologySpreadConstraints); err != nil {
		t.Fatal(err)
	}
	j.Spec.Template.Spec.Containers = []v1.Container{fakeContainer(t)}
	j.Spec.Template.Spec.InitContainers = []v1.Container{fakeContainer(t)}
	j.Spec.Template.Spec.EphemeralContainers = []v1.EphemeralContainer{fakeEphemeralContainer(t)}
	j.Spec.Template.Spec.Volumes = []v1.Volume{fakeVolume(t)}
	j.Spec.Template.Spec.Overhead = map[v1.ResourceName]resource.Quantity{
		"test": {Format: "test"},
	}
	jobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&batchv1.JobList{Items: []batchv1.Job{j}}, nil,
	)
	return client.Services{
		Jobs: jobs,
	}
}

func TestBatchJobs(t *testing.T) {
	k8sTestHelper(t, BatchJobs(), createBatchJobs)
}
