package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type policyConditionCommandHandler struct {
	tools *corebottom.Tools
}

func (pah *policyConditionCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(scope, tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	// TODO: I don't think it's as simple as this and depends on the "test"
	if len(exprs) != 3 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "condition: <test> <left> <right>")
		return drivertop.NewIgnoreInnerScope()
	}

	pa := &PolicyCondAction{tools: pah.tools, loc: tokens[0].Loc(), test: exprs[0], left: exprs[1], right: exprs[2]}
	parent.Attach(pa)

	return drivertop.NewDisallowInnerScope(pah.tools.CoreTools)
}

func NewPolicyConditionCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &policyConditionCommandHandler{tools: tools}
}
