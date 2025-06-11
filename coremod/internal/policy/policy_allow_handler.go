package policy

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyAllowCommandHandler struct {
	tools *pluggable.Tools
}

func (pah *policyAllowCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}
	log.Printf("pah has %d exprs\n", len(exprs))

	if len(exprs) > 3 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "allow: <action> <resource> <principal>")
		return interpreters.IgnoreInnerScope()
	}

	pa := &PolicyAllowAction{tools: pah.tools, loc: tokens[0].Loc(), allowActions: []pluggable.Expr{exprs[0]}, allowResources: []pluggable.Expr{exprs[1]}, allowPrincipals: []pluggable.Expr{exprs[2]}}
	parent.Add(pa)

	// TODO: not true; should allow inner things like "actions", "resources" and "principals"
	return interpreters.DisallowInnerScope(pah.tools)
}

func NewPolicyAllowCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &policyAllowCommandHandler{tools: tools}
}
