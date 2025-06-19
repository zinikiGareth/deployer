package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyAllowCommandHandler struct {
	tools *external.Tools
}

func (pah *policyAllowCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	if len(exprs) > 2 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "allow: <action> <resource>")
		return interpreters.NewIgnoreInnerScope()
	}

	pa := &PolicyAllowAction{tools: pah.tools, loc: tokens[0].Loc(), allowActions: []pluggable.Expr{exprs[0]}, allowResources: []pluggable.Expr{exprs[1]}}
	parent.Attach(pa)

	return interpreters.NewVerbCommandInterpreter(pah.tools.CoreTools, pa, "policy-inner", false)
}

func NewPolicyAllowCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &policyAllowCommandHandler{tools: tools}
}
