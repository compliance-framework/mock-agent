// Package policy parses Rego modules with OPA. It exists so that mock-agent
// genuinely depends on github.com/open-policy-agent/opa (and go mod tidy keeps
// the requirement), pinned at the same version as the real agent.
package policy

import "github.com/open-policy-agent/opa/v1/ast"

// Validate parses a Rego module and returns its package path, e.g.
// "data.mock". It returns the parse error when the module is invalid.
func Validate(filename, module string) (string, error) {
	m, err := ast.ParseModule(filename, module)
	if err != nil {
		return "", err
	}
	return m.Package.Path.String(), nil
}
