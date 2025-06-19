package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyConditionCommandHandler struct {
	tools *external.Tools
}

func (pah *policyConditionCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	// TODO: I don't think it's as simple as this and depends on the "test"
	if len(exprs) != 3 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "condition: <test> <left> <right>")
		return interpreters.NewIgnoreInnerScope()
	}

	pa := &PolicyCondAction{tools: pah.tools, loc: tokens[0].Loc(), test: exprs[0], left: exprs[1], right: exprs[2]}
	parent.Attach(pa)

	return interpreters.NewDisallowInnerScope(pah.tools.CoreTools)
}

func NewPolicyConditionCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &policyConditionCommandHandler{tools: tools}
}
