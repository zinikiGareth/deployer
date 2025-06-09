package basic_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestCasting(t *testing.T) {
	var a pluggable.Action = &basic.EnsureAction{}
	var b pluggable.Action = &basic.EnvAction{}
	if a == b {
		t.Fatalf("Huh?")
	}
}
