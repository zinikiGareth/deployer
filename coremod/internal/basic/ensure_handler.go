package basic

import (
	"ziniki.org/deployer/coremod/internal/vars"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureCommandHandler struct {
	tools *pluggable.Tools
}

func (ech *EnsureCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	if len(tokens) < 2 || len(tokens) > 3 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return interpreters.IgnoreInnerScope()
	}

	clz, ok := tokens[1].(pluggable.Identifier)
	if !ok {
		ech.tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return interpreters.IgnoreInnerScope()
	}

	var name pluggable.String
	if len(tokens) == 3 {
		name, ok = tokens[2].(pluggable.String)
		if !ok {
			ech.tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
			return interpreters.IgnoreInnerScope()
		}
	}

	ea := &EnsureAction{tools: ech.tools, loc: tokens[0].Loc(), what: clz, named: name, props: make(map[pluggable.Identifier]pluggable.Expr)}
	parent.Add(ea)

	if assignTo != nil {
		ech.tools.Repository.IntroduceSymbol(pluggable.SymbolName(assignTo.Id()), ea)
		parent.Add(vars.BindVar(assignTo, ea))
	}
	return interpreters.PropertiesInnerScope(ech.tools, ea)
}

func NewEnsureCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &EnsureCommandHandler{tools: tools}
}
