package files

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type copyCommandHandler struct {
	tools *external.Tools
}

func (cch *copyCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
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

	return interpreters.NewDisallowInnerScope(cch.tools.CoreTools) // for now, but we want to support it really
}

func NewCopyCommandHandler(tools *external.Tools) driverbottom.VerbCommand {
	return &copyCommandHandler{tools: tools}
}
