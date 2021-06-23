package resources

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/thoas/go-funk"
)

func objectCommonColumns() []schema.Column {
	return []schema.Column{
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
	}
}

func PathResolver(path string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
		res <- funk.Get(parent.Item, path, funk.WithAllowZero())
		return nil
	}
}

func PathToJSONResolver(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		json, err := json.Marshal(funk.Get(resource.Item, path, funk.WithAllowZero()))
		if err != nil {
			return err
		}
		stringified := string(json)
		if stringified == "null" {
			return resource.Set(c.Name, nil)
		}
		return resource.Set(c.Name, stringified)
	}
}
