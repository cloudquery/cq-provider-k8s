package resources

import (
	"github.com/cloudquery/faker/v3"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

func fakeSkipFields(tst *testing.T, data interface{}, skipFields []string) {
	skipMap := make(map[string][]string)
	for _, s := range skipFields {
		parts := strings.Split(s, ".")
		if len(parts) == 1 {
			skipMap[s] = []string{}
		} else {
			skipMap[parts[0]] = append(skipMap[parts[0]], strings.Join(parts[1:], "."))
		}
	}
	var v reflect.Value
	switch d := data.(type) {
	case reflect.Value:
		v = d
	default:
		v = reflect.ValueOf(data)
	}
	ind := reflect.Indirect(v)
	s := ind.Type()

	for i := 0; i < s.NumField(); i++ {
		if fields, ok := skipMap[s.Field(i).Name]; !ok {
			ifc := ind.Field(i).Interface()
			if err := faker.FakeData(&ifc); err != nil {
				tst.Fatal(err)
			}
			ind.Field(i).Set(reflect.ValueOf(ifc))
		} else if len(fields) > 0 {
			fakeSkipFields(tst, ind.Field(i), fields)
		}
	}
}

func clearMultiplexers(f func() *provider.Provider) func() *provider.Provider {
	return func() *provider.Provider {
		p := f()
		for _, table := range p.ResourceMap {
			table.Multiplex = nil
		}
		return p
	}
}

func k8sTestHelper(t *testing.T, table *schema.Table, builder func(t *testing.T, ctrl *gomock.Controller) client.Services) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resource := providertest.ResourceTestData{
		Table:  table,
		Config: client.Config{},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			return &client.Client{
				Log: logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}),
				Services: builder(t, ctrl),
			}, nil
		},
	}
	providertest.TestResource(t, clearMultiplexers(Provider), resource)
}
