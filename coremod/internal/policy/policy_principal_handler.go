package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type policyPrincipalCommandHandler struct {
	tools *corebottom.Tools
}

func (pah *policyPrincipalCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(scope, tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	if len(exprs) != 2 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "principal: <type> <id>")
		return drivertop.NewIgnoreInnerScope()
	}

	pa := &policyPrincipalAction{tools: pah.tools, loc: tokens[0].Loc(), ofType: exprs[0], id: exprs[1]}
	parent.Attach(pa)

	return drivertop.NewDisallowInnerScope(pah.tools.CoreTools)
}

func NewPolicyPrincipalCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &policyPrincipalCommandHandler{tools: tools}
}
