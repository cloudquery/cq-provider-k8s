package main

import (
	"github.com/cloudquery/cq-provider-k8s/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		JSONFormat: true,
	})

	serve.Serve(&serve.Options{
		Name:                "k8s",
		Provider:            resources.Provider(),
		Logger:              logger,
		NoLogOutputOverride: false,
	})
}
