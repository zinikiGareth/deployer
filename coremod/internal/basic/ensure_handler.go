package basic

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// For each action verb, there is exactly one handler.  It is created up front and it does the job of parsing lines and creating individual actions.

type EnsureCommandHandler struct {
	tools *external.Tools
}

func (ech *EnsureCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 || len(tokens) > 3 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return interpreters.NewIgnoreInnerScope()
	}

	clz, ok := tokens[1].(pluggable.Identifier)
	if !ok {
		ech.tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return interpreters.NewIgnoreInnerScope()
	}

	var name pluggable.String
	if len(tokens) == 3 {
		name, ok = tokens[2].(pluggable.String)
		if !ok {
			ech.tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
			return interpreters.NewIgnoreInnerScope()
		}
	}

	ea := &EnsureAction{tools: ech.tools, loc: tokens[0].Loc(), what: clz, named: name, props: make(map[pluggable.Identifier]pluggable.Expr)}
	parent.Attach(ea)

	return interpreters.NewPropertiesInnerScope(&ech.tools.CoreTools, ea)
}

func NewEnsureCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &EnsureCommandHandler{tools: tools}
}
