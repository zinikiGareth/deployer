package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ignoreScope struct {
}

func (b *ignoreScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	// we are just ignoring this (presumably there was an outer error, which has already been reported)
	return b // ignore anything inside here too ...
}

func (b *ignoreScope) Completed() {
}

func IgnoreInnerScope() pluggable.Interpreter {
	return &ignoreScope{}
}
