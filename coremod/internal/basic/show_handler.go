package basic

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type showCommandHandler struct {
	tools *external.Tools
}

func (sch *showCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 {
		sch.tools.Reporter.Report(tokens[0].Loc().Offset, "show: expr...")
		return interpreters.NewIgnoreInnerScope()
	}
	// if assignTo != nil {
	// 	sch.tools.Reporter.Report(tokens[0].Loc().Offset, "show: cannot assign an output variable")
	// 	return interpreters.NewIgnoreInnerScope()
	// }

	exprs, ok := sch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	sa := &ShowAction{tools: sch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(sa)

	return interpreters.NewDisallowInnerScope(sch.tools.CoreTools)
}

func NewShowCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &showCommandHandler{tools: tools}
}
