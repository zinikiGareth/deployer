package methods_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/methods"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestType(t *testing.T) {
	var _ driverbottom.Function = methods.MakeInvokeFunc(nil)
	var _ driverbottom.Expr = &methods.InvokeExpr{}
}
