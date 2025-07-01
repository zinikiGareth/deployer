package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type policyCommandHandler struct {
	tools *corebottom.Tools
}

func (pch *policyCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 1 {
		pch.tools.Reporter.Report(tokens[0].Loc().Offset, "policy: this may in fact be valid, but if so I don't know how ")
		return drivertop.NewIgnoreInnerScope()
	}

	pa := &PolicyAction{tools: pch.tools, loc: tokens[0].Loc()}
	parent.Attach(pa)

	return drivertop.NewVerbCommandInterpreter(pch.tools.CoreTools, pa, "policy-statements", false)
}

func NewPolicyCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &policyCommandHandler{tools: tools}
}
