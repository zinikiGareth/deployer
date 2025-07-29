package drivertop

import (
	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/internal/lists"
	"ziniki.org/deployer/driver/internal/methods"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func RegisterBasicFunctions(tools *driverbottom.CoreTools) {
	// loc := &errorsink.Location{}
	tools.Register.ExtensionPoint("function-defn")

	tools.Register.Register("function-defn", "->", methods.MakeInvokeFunc(tools))
	tools.Register.Register("function-defn", "sum", lists.MakeSumFunc(tools))

	tools.Repository.TopScope().IntroduceSymbol(driverbottom.SymbolName("false"), lexicator.NewNumberToken(&errorsink.LineLoc{}, 0, 0))
	tools.Repository.TopScope().IntroduceSymbol(driverbottom.SymbolName("true"), lexicator.NewNumberToken(&errorsink.LineLoc{}, 0, 1))

	basicmath.RegisterAll(tools)
}

func MakeInvokeExpr(on driverbottom.Expr, call driverbottom.Identifier, args ...driverbottom.Expr) driverbottom.Expr {
	return methods.MakeInvokeExpr(on, call, args)
}
