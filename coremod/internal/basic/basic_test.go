package basic_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestCasting(t *testing.T) {
	var a driverbottom.ModelBuilder = &basic.EnsureAction{}
	var b driverbottom.ModelBuilder = &basic.EnvAction{}
	if a == b {
		t.Fatalf("Huh?")
	}
}
