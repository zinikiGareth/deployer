package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type policyActionCommandHandler struct {
	tools *corebottom.Tools
}

func (pah *policyActionCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	exprs, ok := pah.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	if len(exprs) < 1 {
		pah.tools.Reporter.Report(tokens[0].Loc().Offset, "action: <action> ...")
		return drivertop.NewIgnoreInnerScope()
	}

	pa := &policyAction{tools: pah.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(pa)

	return drivertop.NewDisallowInnerScope(pah.tools.CoreTools)
}

func NewPolicyActionCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &policyActionCommandHandler{tools: tools}
}
