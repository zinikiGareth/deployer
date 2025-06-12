package basic_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestCasting(t *testing.T) {
	var a pluggable.ModelBuilder = &basic.EnsureAction{}
	var b pluggable.ModelBuilder = &basic.EnvAction{}
	if a == b {
		t.Fatalf("Huh?")
	}
}
