package basic

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type envCommandHandler struct {
	tools *external.Tools
}

func (ech *envCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "env: expr")
		return interpreters.NewIgnoreInnerScope()
	}

	expr, ok := ech.tools.Parser.Parse(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	ea := &EnvAction{tools: ech.tools, loc: tokens[0].Loc(), varname: expr}
	parent.Attach(ea)

	return interpreters.NewDisallowInnerScope(ech.tools.CoreTools)
}

func NewEnvCommandHandler(tools *external.Tools) driverbottom.VerbCommand {
	return &envCommandHandler{tools: tools}
}
