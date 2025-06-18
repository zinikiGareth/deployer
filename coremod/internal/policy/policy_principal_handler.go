package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyPrincipalCommandHandler struct {
	tools *external.Tools
}

func (pah *policyPrincipalCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	if len(exprs) != 2 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "principal: <type> <id>")
		return interpreters.NewIgnoreInnerScope()
	}

	pa := &policyPrincipalAction{tools: pah.tools, loc: tokens[0].Loc(), ofType: exprs[0], id: exprs[1]}
	parent.Attach(pa)

	return interpreters.NewDisallowInnerScope(&pah.tools.CoreTools)
}

func NewPolicyPrincipalCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &policyPrincipalCommandHandler{tools: tools}
}
