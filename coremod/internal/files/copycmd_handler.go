package files

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type copyCommandHandler struct {
	tools *pluggable.Tools
}

func (cch *copyCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 {
		cch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: <from> <to>")
		return interpreters.NewIgnoreInnerScope()
	}

	exprs, ok := cch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	if len(exprs) != 2 {
		cch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: <from> <to>")
		return interpreters.NewIgnoreInnerScope()
	}

	ca := &copyAction{tools: cch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(ca)

	return interpreters.NewDisallowInnerScope(cch.tools) // for now, but we want to support it really
}

func NewCopyCommandHandler(tools *pluggable.Tools) pluggable.VerbCommand {
	return &copyCommandHandler{tools: tools}
}
