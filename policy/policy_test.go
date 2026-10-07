package policy

import "testing"

func TestPackagePath(t *testing.T) {
	got, err := PackagePath("mock.rego", "package mock\n\nallow if true\n")
	if err != nil {
		t.Fatalf("PackagePath: unexpected error: %v", err)
	}
	if want := "data.mock"; got != want {
		t.Fatalf("PackagePath() = %q, want %q", got, want)
	}
}

func TestPackagePathRejectsInvalidRego(t *testing.T) {
	if _, err := PackagePath("bad.rego", "package\n"); err == nil {
		t.Fatal("PackagePath: expected an error for an invalid module, got nil")
	}
}
