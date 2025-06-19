package files

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type dirCommandHandler struct {
	tools *external.Tools
}

func (dch *dirCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: expr...")
		return interpreters.NewIgnoreInnerScope()
	}
	// if assignTo == nil {
	// 	dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: must assign to an output variable")
	// 	return interpreters.NewIgnoreInnerScope()
	// }

	exprs, ok := dch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	da := &dirAction{tools: dch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(da)

	return interpreters.NewDisallowInnerScope(dch.tools.CoreTools)
}

func NewDirCommandHandler(tools *external.Tools) driverbottom.VerbCommand {
	return &dirCommandHandler{tools: tools}
}
