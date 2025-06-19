package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/interpreters"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type policyActionCommandHandler struct {
	tools *external.Tools
}

func (pah *policyActionCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	if len(exprs) < 1 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "action: <action> ...")
		return interpreters.NewIgnoreInnerScope()
	}

	pa := &policyAction{tools: pah.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(pa)

	return interpreters.NewDisallowInnerScope(pah.tools.CoreTools)
}

func NewPolicyActionCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &policyActionCommandHandler{tools: tools}
}
