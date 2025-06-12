package policy

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyConditionCommandHandler struct {
	tools *pluggable.Tools
}

func (pah *policyConditionCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	// TODO: I don't think it's as simple as this and depends on the "test"
	if len(exprs) != 3 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "condition: <test> <left> <right>")
		return interpreters.IgnoreInnerScope()
	}

	pa := &PolicyCondAction{tools: pah.tools, loc: tokens[0].Loc(), test: exprs[0], left: exprs[1], right: exprs[2]}
	parent.Add(pa)

	return interpreters.DisallowInnerScope(pah.tools)
}

func NewPolicyConditionCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &policyConditionCommandHandler{tools: tools}
}
