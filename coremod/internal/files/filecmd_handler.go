package files

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corepkg"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type fileCommandHandler struct {
	tools *corebottom.Tools
}

func (dch *fileCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.file: expr...")
		return drivertop.NewIgnoreInnerScope()
	}

	exprs, ok := dch.tools.Parser.ParseMultiple(scope, tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	da := corepkg.NewCoreAction(dch.tools, tokens[0].Loc(), "FileAction", &fileAction{exprs: exprs})
	parent.Attach(da)

	return drivertop.NewDisallowInnerScope(dch.tools.CoreTools)
}

func NewFileCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &fileCommandHandler{tools: tools}
}
