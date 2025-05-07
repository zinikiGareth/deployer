package basic

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type showCommandHandler struct {
	tools *pluggable.Tools
}

func (sch *showCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 {
		sch.tools.Reporter.Report(tokens[0].Loc().Offset, "show: expr...")
		return interpreters.IgnoreInnerScope()
	}
	// if assignTo != nil {
	// 	sch.tools.Reporter.Report(tokens[0].Loc().Offset, "show: cannot assign an output variable")
	// 	return interpreters.IgnoreInnerScope()
	// }

	exprs, ok := sch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	sa := &ShowAction{tools: sch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Add(sa)

	return interpreters.DisallowInnerScope(sch.tools)
}

func NewShowCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &showCommandHandler{tools: tools}
}
