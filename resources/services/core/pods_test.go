//go:build mock
// +build mock

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationPods(t *testing.T) {
	client.K8sTestHelper(t, Pods(), "./snapshots")
}
