package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "k8s",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"pod": Pod(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
