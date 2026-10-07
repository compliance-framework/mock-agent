# mock-agent

Mock repo for developing CCF release automation. Not a product.

It mirrors [compliance-framework/agent](https://github.com/compliance-framework/agent):

- a Go module that requires `mock-api` and `mock-gooci` (as agent requires `api` and
  `gooci`) and `github.com/open-policy-agent/opa` at the same version as agent;
- an exported `runner` package (`Plugin` interface) that the mock plugins import;
- three images, like agent's: `Dockerfile`, `Dockerfile-ci` and `Dockerfile-custodian`.
  Each one ships the static binary at `/app/mock-agent`, which mock-agent-action copies.

```sh
make build          # dist/mock-agent
make test
make docker-build   # all three images
```
