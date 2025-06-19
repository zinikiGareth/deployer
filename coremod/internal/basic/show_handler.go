package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type showCommandHandler struct {
	tools *corebottom.Tools
}

func (sch *showCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		sch.tools.Reporter.Report(tokens[0].Loc().Offset, "show: expr...")
		return drivertop.NewIgnoreInnerScope()
	}
	// if assignTo != nil {
	// 	sch.tools.Reporter.Report(tokens[0].Loc().Offset, "show: cannot assign an output variable")
	// 	return drivertop.NewIgnoreInnerScope()
	// }

	exprs, ok := sch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	sa := &ShowAction{tools: sch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(sa)

	return drivertop.NewDisallowInnerScope(sch.tools.CoreTools)
}

func NewShowCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &showCommandHandler{tools: tools}
}
