package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type policyCommandHandler struct {
	tools *external.Tools
}

func (pch *policyCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 1 {
		pch.tools.Reporter.Report(tokens[0].Loc().Offset, "policy: this may in fact be valid, but if so I don't know how ")
		return interpreters.NewIgnoreInnerScope()
	}

	pa := &PolicyAction{tools: pch.tools, loc: tokens[0].Loc()}
	parent.Attach(pa)

	return interpreters.NewVerbCommandInterpreter(pch.tools.CoreTools, pa, "policy-statements", false)
}

func NewPolicyCommandHandler(tools *external.Tools) driverbottom.VerbCommand {
	return &policyCommandHandler{tools: tools}
}
