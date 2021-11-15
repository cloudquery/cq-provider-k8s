package integration_tests

import (
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationCoreNodes(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.CoreNodes(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_core_nodes",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = 'minikube'")
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": "minikube",
				}},
			},
		}
	})
}
