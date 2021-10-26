package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func createAppsDaemonSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	daemonSetsClient := mocks.NewMockDaemonSetsClient(ctrl)
	daemonSetsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.DaemonSetList{Items: []appsv1.DaemonSet{fakeDaemonSet(t)}}, nil,
	)
	return client.Services{
		DaemonSets: daemonSetsClient,
	}
}

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
	fakeThroughPointers(t, []interface{}{
		&templateSpec.Annotations,
		&templateSpec.Name,
		&templateSpec.GenerateName,
		&templateSpec.Namespace,
		&templateSpec.SelfLink,
		&templateSpec.UID,
		&templateSpec.ResourceVersion,
		&templateSpec.Generation,
		&templateSpec.DeletionGracePeriodSeconds,
		&templateSpec.Labels,
		&templateSpec.Finalizers,
		&templateSpec.ClusterName,
		&templateSpec.OwnerReferences,
		&templateSpec.ManagedFields,
	})
	templateSpec.Spec = fakePodSpec(t)
	return templateSpec
}

func TestAppsDaemonSets(t *testing.T) {
	k8sTestHelper(t, AppsDaemonSets(), createAppsDaemonSets)
}
