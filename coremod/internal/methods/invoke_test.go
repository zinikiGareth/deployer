package methods_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/methods"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestType(t *testing.T) {
	var _ pluggable.Function = methods.MakeInvokeFunc(nil)
	var _ pluggable.Expr = &methods.InvokeExpr{}
}
