// Command mock-agent mirrors the shape of the CCF agent in miniature, for
// developing shared CI and release workflows. It depends on mock-api and
// mock-gooci the way the real agent depends on api and gooci.
package main

import (
	"fmt"
	"io"
	"os"

	mockapi "github.com/compliance-framework/mock-api/pkg/version"
	"github.com/compliance-framework/mock-gooci/pkg/oci"
)

// imageName is the repository mock-agent's images are published to.
const imageName = "ghcr.io/compliance-framework/mock-agent"

// buildVersion is the version of this binary. The Makefile (and so every
// Dockerfile) sets it with -ldflags "-X main.buildVersion=<VERSION>".
var buildVersion = "dev"

func run(w io.Writer, version string) error {
	_, err := fmt.Fprintf(w, "%s\nmock-agent %s (image %s)\n", mockapi.Hello(), version, oci.Ref(imageName, version))
	return err
}

func main() {
	if err := run(os.Stdout, buildVersion); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
