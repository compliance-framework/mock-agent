package runner

import "testing"

type fakePlugin struct{}

func (fakePlugin) Name() string { return "fake" }

func TestPluginInterface(t *testing.T) {
	var p Plugin = fakePlugin{}
	if got := p.Name(); got != "fake" {
		t.Fatalf("Name() = %q, want %q", got, "fake")
	}
}
