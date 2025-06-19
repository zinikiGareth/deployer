package target_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestBasicTypes(t *testing.T) {
	var _ driverbottom.TopLevelForm = &target.CoreTarget{}
}
