package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type policyAllowCommandHandler struct {
	tools *external.Tools
}

func (pah *policyAllowCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	if len(exprs) > 2 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "allow: <action> <resource>")
		return drivertop.NewIgnoreInnerScope()
	}

	pa := &PolicyAllowAction{tools: pah.tools, loc: tokens[0].Loc(), allowActions: []driverbottom.Expr{exprs[0]}, allowResources: []driverbottom.Expr{exprs[1]}}
	parent.Attach(pa)

	return drivertop.NewVerbCommandInterpreter(pah.tools.CoreTools, pa, "policy-inner", false)
}

func NewPolicyAllowCommandHandler(tools *external.Tools) driverbottom.VerbCommand {
	return &policyAllowCommandHandler{tools: tools}
}
