package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type NgnoreScope struct {
}

func (b *NgnoreScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	// we are just ignoring this (presumably there was an outer error, which has already been reported)
	return b // ignore anything inside here too ...
}

func (b *NgnoreScope) Completed() {
}

func NewIgnoreInnerScope() pluggable.Interpreter {
	return &NgnoreScope{}
}

func NewDisallowInnerScope(tools *pluggable.Tools) pluggable.Interpreter {
	return &DisallowInnerScope{tools: tools}
}

func NewPropertiesInnerScope(tools *pluggable.Tools, parent pluggable.PropertyParent) pluggable.Interpreter {
	return &PropertiesInnerScope{tools: tools, parent: parent}
}
