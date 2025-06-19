package files

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type copyCommandHandler struct {
	tools *corebottom.Tools
}

func (cch *copyCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		cch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: <from> <to>")
		return drivertop.NewIgnoreInnerScope()
	}

	exprs, ok := cch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	if len(exprs) != 2 {
		cch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: <from> <to>")
		return drivertop.NewIgnoreInnerScope()
	}

	ca := &copyAction{tools: cch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(ca)

	return drivertop.NewDisallowInnerScope(cch.tools.CoreTools) // for now, but we want to support it really
}

func NewCopyCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &copyCommandHandler{tools: tools}
}
