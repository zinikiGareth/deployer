package target

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type CoreTargetHandler struct {
	tools *corebottom.Tools
}

func (t *CoreTargetHandler) Handle(attacher driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 2 {
		t.tools.Reporter.Reportf(0, "target: <name>")
		return drivertop.NewIgnoreInnerScope()
	}
	t1 := tokens[1].(driverbottom.Identifier)
	name := driverbottom.SymbolName(t1.Id())
	target := &CoreTarget{tools: t.tools, loc: t1.Loc(), name: name, actions: []driverbottom.ModelBuilder{}}

	attacher.Attach(target)
	return drivertop.NewVerbCommandInterpreter(t.tools.CoreTools, target, "target", true)
}

func MakeCoreTargetVerb(tools *corebottom.Tools) *CoreTargetHandler {
	return &CoreTargetHandler{tools: tools}
}

var _ driverbottom.VerbCommand = &CoreTargetHandler{}
