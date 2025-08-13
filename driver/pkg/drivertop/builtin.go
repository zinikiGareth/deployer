package drivertop

import (
	"ziniki.org/deployer/driver/internal/methods"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func MakeInvokeExpr(on driverbottom.Expr, call driverbottom.Identifier, args ...driverbottom.Expr) driverbottom.Expr {
	return methods.MakeInvokeExpr(on, call, args)
}
