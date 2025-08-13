package impl

import (
	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/internal/lists"
	"ziniki.org/deployer/driver/internal/methods"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func RegisterBasicFunctions(tools *driverbottom.CoreTools) {
	// register the main execution entry point here.  Exactly one module needs to do this.
	tools.Register.ExtensionPoint("main-args")

	// register all top level forms here.  There must be at least one in order to do anything useful.
	tools.Register.ExtensionPoint("top-level")

	// register the top level attacher that will collect all the items together at the top level.
	tools.Register.ExtensionPoint("attacher")

	// functions are within our remit, so we provide this extension point.
	tools.Register.ExtensionPoint("function-defn")

	tools.Register.Register("function-defn", "->", methods.MakeInvokeFunc(tools))
	tools.Register.Register("function-defn", "sum", lists.MakeSumFunc(tools))

	tools.Repository.TopScope().IntroduceSymbol(driverbottom.SymbolName("false"), lexicator.NewNumberToken(&errorsink.LineLoc{}, 0, 0))
	tools.Repository.TopScope().IntroduceSymbol(driverbottom.SymbolName("true"), lexicator.NewNumberToken(&errorsink.LineLoc{}, 0, 1))

	basicmath.RegisterAll(tools)
}
