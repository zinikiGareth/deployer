package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type AttachPolicyCommandHandler struct {
	tools *corebottom.Tools
}

func (apch *AttachPolicyCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	exprs, ok := apch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	if len(exprs) != 2 {
		apch.tools.Reporter.Report(tokens[0].Loc().Offset, "attachPolicy: <to> <policy>")
		return drivertop.NewIgnoreInnerScope()
	}

	ea := &AttachPolicyAction{tools: apch.tools, loc: tokens[0].Loc(), to: exprs[0], policy: exprs[1]}
	parent.Attach(ea)

	return drivertop.NewPropertiesInnerScope(apch.tools.CoreTools, ea)
}

func NewAttachPolicyCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &AttachPolicyCommandHandler{tools: tools}
}
