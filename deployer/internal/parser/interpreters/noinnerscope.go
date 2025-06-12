package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type DisallowInnerScope struct {
	tools *pluggable.Tools
}

func (b *DisallowInnerScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	b.tools.Reporter.Report(0, "nested content is not allowed here")
	return b
}

func (b *DisallowInnerScope) Completed() {
}
