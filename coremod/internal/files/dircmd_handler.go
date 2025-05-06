package files

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type dirCommandHandler struct {
	tools *pluggable.Tools
}

func (dch *dirCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token, assignTo pluggable.Identifier) pluggable.Interpreter {
	return interpreters.DisallowInnerScope(dch.tools)
}

func NewDirCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &dirCommandHandler{tools: tools}
}
