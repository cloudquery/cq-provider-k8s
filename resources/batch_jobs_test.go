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

/*
   Volumes                       []Volume                   `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
    InitContainers                []Container                `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
    Containers                    []Container                `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
    EphemeralContainers           []EphemeralContainer       `json:"ephemeralContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,34,rep,name=ephemeralContainers"`
    RestartPolicy                 RestartPolicy              `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy"`
    TerminationGracePeriodSeconds *int64                     `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds"`
    ActiveDeadlineSeconds         *int64                     `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=activeDeadlineSeconds"`
    DNSPolicy                     DNSPolicy                  `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy"`
    NodeSelector                  map[string]string          `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`
    ServiceAccountName            string                     `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"`
    DeprecatedServiceAccount      string                     `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount"`
    AutomountServiceAccountToken  *bool                      `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken"`
    NodeName                      string                     `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
    HostNetwork                   bool                       `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork"`
    HostPID                       bool                       `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`
    HostIPC                       bool                       `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC"`
    ShareProcessNamespace         *bool                      `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace"`
    SecurityContext               *PodSecurityContext        `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
    ImagePullSecrets              []LocalObjectReference     `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
    Hostname                      string                     `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"`
    Subdomain                     string                     `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"`
    Affinity                      *Affinity                  `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
    SchedulerName                 string                     `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"`
    Tolerations                   []Toleration               `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
    HostAliases                   []HostAlias                `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
    PriorityClassName             string                     `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName"`
    Priority                      *int32                     `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority"`
    DNSConfig                     *PodDNSConfig              `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig"`
    ReadinessGates                []PodReadinessGate         `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates"`
    RuntimeClassName              *string                    `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
    EnableServiceLinks            *bool                      `json:"enableServiceLinks,omitempty" protobuf:"varint,30,opt,name=enableServiceLinks"`
    PreemptionPolicy              *PreemptionPolicy          `json:"preemptionPolicy,omitempty" protobuf:"bytes,31,opt,name=preemptionPolicy"`
    Overhead                      ResourceList               `json:"overhead,omitempty" protobuf:"bytes,32,opt,name=overhead"`
    TopologySpreadConstraints     []TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty" patchStrategy:"merge" patchMergeKey:"topologyKey" protobuf:"bytes,33,opt,name=topologySpreadConstraints"`
    SetHostnameAsFQDN             *bool                      `json:"setHostnameAsFQDN,omitempty" protobuf:"varint,35,opt,name=setHostnameAsFQDN"`

*/
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
	faker.FakeData(&j.Spec.Template.Spec.RestartPolicy)
	faker.FakeData(&j.Spec.Template.Spec.TopologySpreadConstraints)
	j.Spec.Template.Spec.RestartPolicy = "test"
	c := fakeContainer(t)
	j.Spec.Template.Spec.Containers = []v1.Container{c}
	j.Spec.Template.Spec.InitContainers = []v1.Container{c}
	ec := fakeEphemeralContainer(t)
	j.Spec.Template.Spec.EphemeralContainers = []v1.EphemeralContainer{ec}
	vm := fakeVolume(t)
	j.Spec.Template.Spec.Volumes = []v1.Volume{vm}
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
