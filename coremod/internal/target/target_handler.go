package target

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type CoreTargetHandler struct {
	tools *external.Tools
}

func (t *CoreTargetHandler) Handle(attacher driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 2 {
		t.tools.Reporter.Reportf(0, "target: <name>")
		return interpreters.NewIgnoreInnerScope()
	}
	t1 := tokens[1].(driverbottom.Identifier)
	name := driverbottom.SymbolName(t1.Id())
	target := &CoreTarget{loc: t1.Loc(), name: name, actions: []driverbottom.ModelBuilder{}}

	attacher.Attach(target)
	return interpreters.NewVerbCommandInterpreter(t.tools.CoreTools, target, "target", true)
}

func MakeCoreTargetVerb(tools *external.Tools) *CoreTargetHandler {
	return &CoreTargetHandler{tools: tools}
}

var _ driverbottom.VerbCommand = &CoreTargetHandler{}
