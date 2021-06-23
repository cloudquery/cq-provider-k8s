package client

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	// Usually you store here your 3rd party clients and use them in the fetcher
	KClient *kubernetes.Clientset
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	_ = providerConfig
	flagSet := flag.NewFlagSet("", flag.PanicOnError)
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flagSet.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flagSet.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flagSet.Parse([]string{})
	// use the current context in kubeconfig
	clientConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	client := Client{
		logger:  logger,
		KClient: clientset,
	}

	return &client, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
