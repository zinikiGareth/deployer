package basic

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type envCommandHandler struct {
	tools *pluggable.Tools
}

func (ech *envCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	if len(tokens) < 2 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "env: expr")
		return interpreters.IgnoreInnerScope()
	}

	expr, ok := ech.tools.Parser.Parse(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	sa := &EnvAction{tools: ech.tools, loc: tokens[0].Loc(), expr: expr, assignTo: assignTo}
	parent.Add(sa)

	return interpreters.DisallowInnerScope(ech.tools)
}

func NewEnvCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &envCommandHandler{tools: tools}
}
