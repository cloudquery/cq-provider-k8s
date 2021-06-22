package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Pod() *schema.Table {
	return &schema.Table{
		Name:     "k8s_pods",
		Resolver: fetchPods,
		Columns: []schema.Column{
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "namespace",
				Type: schema.TypeString,
			},
			{
				Name: "resource_version",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchPods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	pods, err := c.ThirdPartyClient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	res <- pods.Items
	return nil
}
