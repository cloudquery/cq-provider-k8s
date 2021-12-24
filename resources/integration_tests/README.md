Environment variables for running integration tests:

* `INTEGRATION_TESTS=1`
* `TF_APPLY_RESOURCES=1`
* `TF_VAR_PREFIX=test`
* `TF_VAR_SUFFIX=test`
* `KUBECONFIG=~/.kube/config`
* `KUBE_CONFIG_PATH=~/.kube/config`

Command to run them:
`go test ./resources/integration_tests/...`