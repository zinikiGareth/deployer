package interpreters

import "ziniki.org/deployer/driver/pkg/driverbottom"

type ignoreInnerScope struct {
}

func (b *ignoreInnerScope) HaveTokens(tokens []driverbottom.Token) driverbottom.Interpreter {
	// we are just ignoring this (presumably there was an outer error, which has already been reported)
	return b // ignore anything inside here too ...
}

func (b *ignoreInnerScope) Completed() {
}
