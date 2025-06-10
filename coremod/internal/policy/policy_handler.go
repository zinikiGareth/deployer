package policy

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyCommandHandler struct {
	tools *pluggable.Tools
}

func (ech *policyCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) != 1 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "policy: this may in fact be valid, but if so I don't know how ")
		return interpreters.IgnoreInnerScope()
	}

	ea := &PolicyAction{tools: ech.tools, loc: tokens[0].Loc()}
	parent.Add(ea)

	return interpreters.DisallowInnerScope(ech.tools)
}

func NewPolicyCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &policyCommandHandler{tools: tools}
}
