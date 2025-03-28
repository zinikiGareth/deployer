package target

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type CoreTarget struct {
}

func (t *CoreTarget) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, tokens []pluggable.Token) pluggable.Interpreter {
	t1 := tokens[1].(pluggable.Identifier)
	repo.IntroduceSymbol(t1.Loc(), pluggable.SymbolType("core.Target"), pluggable.SymbolName(t1.Id()))
	return TargetCommandInterpreter(repo)
}
