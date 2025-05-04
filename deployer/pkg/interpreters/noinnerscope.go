package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type noInnerScope struct {
	tools *pluggable.Tools
}

func (b *noInnerScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	b.tools.Reporter.Report(0, "nested content is not allowed here")
	return b
}

func (b *noInnerScope) Completed() {
}

func DisallowInnerScope(tools *pluggable.Tools) pluggable.Interpreter {
	return &noInnerScope{tools: tools}
}
