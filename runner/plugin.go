// Package runner is the exported plugin contract of mock-agent. Mock plugin
// repos import it, mirroring how real CCF plugins import agent's runner
// package, so the release train has a library to pin and bump.
package runner

// Plugin is the minimal interface a mock plugin implements.
type Plugin interface {
	// Name returns the plugin's name.
	Name() string
}
