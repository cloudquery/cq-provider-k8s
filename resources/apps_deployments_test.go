package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func createAppsDeployments(t *testing.T, ctrl *gomock.Controller) client.Services {
	deploymentsClient := mocks.NewMockDeploymentsClient(ctrl)
	deploymentsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.DeploymentList{Items: []appsv1.Deployment{fakeAppsDeployment(t)}}, nil,
	)
	return client.Services{
		Deployments: deploymentsClient,
	}
}

func fakeAppsDeployment(t *testing.T) appsv1.Deployment {
	var deployment appsv1.Deployment
	fakeThroughPointers(t, []interface{}{
		&deployment.TypeMeta,
		&deployment.ObjectMeta,
		&deployment.Status,
		&deployment.ManagedFields,
		&deployment.Annotations,
		&deployment.Labels,
		&deployment.OwnerReferences,
		&deployment.Status,
		&deployment.Spec.Selector,
		&deployment.Spec.RevisionHistoryLimit,
	})

	deployment.Spec.Template = fakePodTemplateSpec(t)

	return deployment
}

func TestAppsDeployments(t *testing.T) {
	k8sTestHelper(t, AppsDeployments(), createAppsDeployments)
}
