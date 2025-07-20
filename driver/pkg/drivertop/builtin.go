package drivertop

import (
	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/internal/lists"
	"ziniki.org/deployer/driver/internal/methods"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func RegisterBasicFunctions(tools *driverbottom.CoreTools) {
	tools.Register.ExtensionPoint("function-defn")

	tools.Register.Register("function-defn", "->", methods.MakeInvokeFunc(tools))
	tools.Register.Register("function-defn", "sum", lists.MakeSumFunc(tools))

	basicmath.RegisterAll(tools)
}

func MakeInvokeExpr(on driverbottom.Expr, call driverbottom.Identifier, args ...driverbottom.Expr) driverbottom.Expr {
	return methods.MakeInvokeExpr(on, call, args)
}
