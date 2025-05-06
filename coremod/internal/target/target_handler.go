package target

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type CoreTargetHandler struct {
	tools *pluggable.Tools
}

func (t *CoreTargetHandler) Handle(tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	if len(tokens) != 2 {
		t.tools.Reporter.Reportf(0, "target: <name>")
		return interpreters.IgnoreInnerScope()
	}
	t1 := tokens[1].(pluggable.Identifier)
	name := pluggable.SymbolName(t1.Id())
	target := &coreTarget{loc: t1.Loc(), name: name, actions: []action{}}
	t.tools.Repository.IntroduceSymbol(name, target)
	return TargetCommandInterpreter(t.tools, target)
}

func MakeCoreTargetVerb(tools *pluggable.Tools) *CoreTargetHandler {
	return &CoreTargetHandler{tools: tools}
}
