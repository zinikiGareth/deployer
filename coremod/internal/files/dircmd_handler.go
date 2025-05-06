package files

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type dirCommandHandler struct {
	tools *pluggable.Tools
}

func (dch *dirCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	if len(tokens) < 2 {
		dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: expr...")
		return interpreters.IgnoreInnerScope()
	}
	if assignTo == nil {
		dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: must assign to an output variable")
		return interpreters.IgnoreInnerScope()
	}

	exprs, ok := dch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	da := &dirAction{tools: dch.tools, loc: tokens[0].Loc(), exprs: exprs, assignTo: assignTo}
	parent.Add(da)

	return interpreters.DisallowInnerScope(dch.tools)
}

func NewDirCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &dirCommandHandler{tools: tools}
}
