package files

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type copyCommandHandler struct {
	tools *pluggable.Tools
}

func (cch *copyCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	if len(tokens) < 2 {
		cch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: <from> <to>")
		return interpreters.IgnoreInnerScope()
	}

	exprs, ok := cch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	if len(exprs) != 2 {
		cch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: <from> <to>")
		return interpreters.IgnoreInnerScope()
	}

	ca := &copyAction{tools: cch.tools, loc: tokens[0].Loc(), exprs: exprs, assignTo: assignTo}
	parent.Add(ca)

	return interpreters.DisallowInnerScope(cch.tools) // for now, but we want to support it really
}

func NewCopyCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &copyCommandHandler{tools: tools}
}
