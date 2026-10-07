package main

import (
	"strings"
	"testing"

	mockapi "github.com/compliance-framework/mock-api/pkg/version"
)

func TestRun(t *testing.T) {
	var out strings.Builder
	if err := run(&out, "v1.2.3"); err != nil {
		t.Fatalf("run: %v", err)
	}
	want := mockapi.Hello() + "\nmock-agent v1.2.3 (image ghcr.io/compliance-framework/mock-agent:v1.2.3)\n"
	if got := out.String(); got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}
