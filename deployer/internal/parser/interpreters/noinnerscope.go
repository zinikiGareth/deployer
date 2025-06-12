package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type disallowInnerScope struct {
	tools *pluggable.Tools
}

func (b *disallowInnerScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	b.tools.Reporter.Report(0, "nested content is not allowed here")
	return b
}

func (b *disallowInnerScope) Completed() {
}
