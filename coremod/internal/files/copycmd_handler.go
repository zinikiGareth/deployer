package files

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type copyCommandHandler struct {
	tools *pluggable.Tools
}

func (cch *copyCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	return interpreters.DisallowInnerScope(cch.tools)
}

func NewCopyCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &copyCommandHandler{tools: tools}
}
