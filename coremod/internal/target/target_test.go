package target_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

func TestBasicTypes(t *testing.T) {
	var _ pluggable.TopLevelForm = &target.CoreTarget{}
}
