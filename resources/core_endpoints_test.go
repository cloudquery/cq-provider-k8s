package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreEndpoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	endpoints := mocks.NewMockEndpointsClient(ctrl)
	e := corev1.Endpoints{}
	if err := faker.FakeData(&e); err != nil {
		t.Fatal(err)
	}
	e.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}
	endpoints.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.EndpointsList{Items: []corev1.Endpoints{e}}, nil,
	)
	return client.Services{
		Endpoints: endpoints,
	}
}

func TestCoreEndpoints(t *testing.T) {
	k8sTestHelper(t, CoreEndpoints(), createCoreEndpoints)
}
