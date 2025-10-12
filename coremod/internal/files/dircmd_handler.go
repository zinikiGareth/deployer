package files

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corepkg"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type dirCommandHandler struct {
	tools *corebottom.Tools
}

func (dch *dirCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: expr...")
		return drivertop.NewIgnoreInnerScope()
	}
	// if assignTo == nil {
	// 	dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: must assign to an output variable")
	// 	return drivertop.NewIgnoreInnerScope()
	// }

	exprs, ok := dch.tools.Parser.ParseMultiple(scope, tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	da := corepkg.NewCoreAction(dch.tools, tokens[0].Loc(), "DirAction", &dirAction{exprs: exprs})
	parent.Attach(da)

	return drivertop.NewDisallowInnerScope(dch.tools.CoreTools)
}

func NewDirCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &dirCommandHandler{tools: tools}
}
