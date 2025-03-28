package basic

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureCommandHandler struct{}

func (ensure *EnsureCommandHandler) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, tokens []pluggable.Token) pluggable.Interpreter {
	return interpreters.DisallowInnerScope()
}
