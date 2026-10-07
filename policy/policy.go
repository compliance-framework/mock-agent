// Package policy parses Rego modules with OPA. It exists so that mock-agent
// genuinely depends on github.com/open-policy-agent/opa (and go mod tidy keeps
// the requirement), pinned at the same version as the real agent.
package policy

import (
	"fmt"

	"github.com/open-policy-agent/opa/v1/ast"
)

// PackagePath parses a Rego module and returns its package path, e.g.
// "data.mock". It returns an error when the module does not parse.
func PackagePath(filename, module string) (string, error) {
	m, err := ast.ParseModule(filename, module)
	if err != nil {
		return "", fmt.Errorf("parse rego module: %w", err)
	}
	return m.Package.Path.String(), nil
}
