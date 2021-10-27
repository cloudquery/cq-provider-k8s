package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func createAppsDeployments(t *testing.T, ctrl *gomock.Controller) client.Services {
	nodes := mocks.NewMockNodesClient(ctrl)
	nodes.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.NodeList{Items: []corev1.Node{fakeNode(t)}}, nil,
	)
	return client.Services{
		Nodes: nodes,
	}
}

func fakeAppsDeployments(t *testing.T) appsv1.Deployment {
	// faker chokes on Node.Status.{Capacity,Allocatable} so doing it by hand
	var deployment appsv1.Deployment
	ptrs := []interface{}{
		&deployment.TypeMeta,
		&deployment.ObjectMeta,
		&deployment.Spec,
		&deployment.Status,
	}
	for i, ptr := range ptrs {
		if err := faker.FakeData(ptr); err != nil {
			t.Fatalf("%v %v", i, ptr)
		}
	}
	//rl := make(corev1.ResourceList)
	//rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	//deployment.Status.Capacity = rl
	//node.Status.Allocatable = rl
	//node.Spec.PodCIDR = "192.168.1.0/24"
	//node.Spec.PodCIDRs = []string{"192.168.1.0/24"}
	//node.Status.Addresses = []corev1.NodeAddress{
	//	{
	//		Type:    corev1.NodeHostName,
	//		Address: "testname",
	//	},
	//	{
	//		Type:    corev1.NodeInternalIP,
	//		Address: "fd00::1",
	//	},
	//	{
	//		Type:    corev1.NodeExternalIP,
	//		Address: "192.168.2.1",
	//	},
	//}
	return deployment
}

func TestAppsDeployments(t *testing.T) {
	k8sTestHelper(t, AppsDeployments(), createAppsDeployments)
}
