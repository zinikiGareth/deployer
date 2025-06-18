package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ignoreInnerScope struct {
}

func (b *ignoreInnerScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	// we are just ignoring this (presumably there was an outer error, which has already been reported)
	return b // ignore anything inside here too ...
}

func (b *ignoreInnerScope) Completed() {
}

func NewIgnoreInnerScope() pluggable.Interpreter {
	return &ignoreInnerScope{}
}

func NewDisallowInnerScope(tools *pluggable.CoreTools) pluggable.Interpreter {
	return &disallowInnerScope{tools: tools}
}

func NewPropertiesInnerScope(tools *pluggable.CoreTools, parent pluggable.PropertyParent) pluggable.Interpreter {
	return &propertiesInterpreter{tools: tools, parent: parent}
}
