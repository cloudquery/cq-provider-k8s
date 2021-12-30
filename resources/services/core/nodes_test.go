//go:build mock
// +build mock

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationNodes(t *testing.T) {
	client.K8sTestHelper(t, Nodes(), "./snapshots")
}
