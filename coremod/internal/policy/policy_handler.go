package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyCommandHandler struct {
	tools *external.Tools
}

func (pch *policyCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) != 1 {
		pch.tools.Reporter.Report(tokens[0].Loc().Offset, "policy: this may in fact be valid, but if so I don't know how ")
		return interpreters.NewIgnoreInnerScope()
	}

	pa := &PolicyAction{tools: pch.tools, loc: tokens[0].Loc()}
	parent.Attach(pa)

	return interpreters.NewVerbCommandInterpreter(pch.tools.CoreTools, pa, "policy-statements", false)
}

func NewPolicyCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &policyCommandHandler{tools: tools}
}
