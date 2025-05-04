package basic

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureCommandHandler struct{}

func (ensure *EnsureCommandHandler) Handle(tools *pluggable.Tools, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 || len(tokens) > 3 {
		tools.Reporter.Report(tokens[0].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return interpreters.IgnoreInnerScope()
	}

	clz, ok := tokens[1].(pluggable.Identifier)
	if !ok {
		tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return interpreters.IgnoreInnerScope()
	}

	var name pluggable.String
	if len(tokens) == 3 {
		name, ok = tokens[2].(pluggable.String)
		if !ok {
			tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
			return interpreters.IgnoreInnerScope()
		}
	}

	ea := &EnsureAction{loc: tokens[0].Loc(), what: clz, named: name, props: make(map[pluggable.Identifier]pluggable.Expr)}
	parent.Add(ea)
	return interpreters.PropertiesInnerScope(ea)
}
