package target

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type coreTarget struct {
	loc pluggable.Location
}

func (t *coreTarget) Where() pluggable.Location {
	return t.loc
}

func (t *coreTarget) What() pluggable.SymbolType {
	return pluggable.SymbolType("core.Target")
}

type CoreTargetVerb struct {
}

func (t *CoreTargetVerb) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	t1 := tokens[1].(pluggable.Identifier)
	target := &coreTarget{loc: t1.Loc()}
	repo.IntroduceSymbol(pluggable.SymbolName(t1.Id()), target)
	return TargetCommandInterpreter(repo)
}
