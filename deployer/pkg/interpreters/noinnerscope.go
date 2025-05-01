package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type noInnerScope struct {
}

func (b *noInnerScope) HaveTokens(reporter errors.ErrorRepI, tokens []pluggable.Token) pluggable.Interpreter {
	reporter.Report(0, "nested content is not allowed here")
	return b
}

func (b *noInnerScope) Completed(reporter errors.ErrorRepI) {
}

func DisallowInnerScope() pluggable.Interpreter {
	return &noInnerScope{}
}
