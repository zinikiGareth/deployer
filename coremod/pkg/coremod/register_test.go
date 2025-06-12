package coremod_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestTypes(t *testing.T) {
	var _ pluggable.VerbCommand = &target.CoreTargetHandler{}
}
