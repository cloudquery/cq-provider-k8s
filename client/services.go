package client

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Services struct {
	Client     *kubernetes.Clientset
	Nodes      NodesClient
	Pods       PodsClient
	Services   ServicesClient
	Namespaces NamespacesClient
}

//go:generate mockgen -package=mocks -destination=./mocks/namespaces.go . NamespacesClient
type NamespacesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/nodes.go . NodesClient
type NodesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NodeList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/pods.go . PodsClient
type PodsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/services.go . ServicesClient
type ServicesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceList, error)
}
