package target

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type CoreTargetHandler struct {
	tools *external.Tools
}

func (t *CoreTargetHandler) Handle(attacher pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) != 2 {
		t.tools.Reporter.Reportf(0, "target: <name>")
		return interpreters.NewIgnoreInnerScope()
	}
	t1 := tokens[1].(pluggable.Identifier)
	name := pluggable.SymbolName(t1.Id())
	target := &CoreTarget{loc: t1.Loc(), name: name, actions: []pluggable.ModelBuilder{}}

	attacher.Attach(target)
	return interpreters.NewVerbCommandInterpreter(t.tools, target, "target", true)
}

func MakeCoreTargetVerb(tools *external.Tools) *CoreTargetHandler {
	return &CoreTargetHandler{tools: tools}
}
