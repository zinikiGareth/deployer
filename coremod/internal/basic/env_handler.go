package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type envCommandHandler struct {
	tools *corebottom.Tools
}

func (ech *envCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "env: expr")
		return drivertop.NewIgnoreInnerScope()
	}

	expr, ok := ech.tools.Parser.Parse(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	ea := &EnvAction{tools: ech.tools, loc: tokens[0].Loc(), varname: expr}
	parent.Attach(ea)

	return drivertop.NewDisallowInnerScope(ech.tools.CoreTools)
}

func NewEnvCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &envCommandHandler{tools: tools}
}
