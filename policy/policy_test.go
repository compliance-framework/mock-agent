package policy

import "testing"

func TestValidate(t *testing.T) {
	got, err := Validate("mock.rego", "package mock\n\nallow if true\n")
	if err != nil {
		t.Fatalf("Validate: unexpected error: %v", err)
	}
	if want := "data.mock"; got != want {
		t.Fatalf("Validate() = %q, want %q", got, want)
	}
}

func TestValidateRejectsInvalidRego(t *testing.T) {
	if _, err := Validate("bad.rego", "package\n"); err == nil {
		t.Fatal("Validate: expected an error for an invalid module, got nil")
	}
}
