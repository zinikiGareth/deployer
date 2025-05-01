package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type noInnerScope struct {
}

func (b *noInnerScope) HaveTokens(tools *pluggable.Tools, tokens []pluggable.Token) pluggable.Interpreter {
	tools.Reporter.Report(0, "nested content is not allowed here")
	return b
}

func (b *noInnerScope) Completed(tools *pluggable.Tools) {
}

func DisallowInnerScope() pluggable.Interpreter {
	return &noInnerScope{}
}
