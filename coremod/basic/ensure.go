package basic

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureCommandHandler struct{}

func (ensure *EnsureCommandHandler) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	parent.Add(tokens[0])
	return interpreters.DisallowInnerScope()
}
