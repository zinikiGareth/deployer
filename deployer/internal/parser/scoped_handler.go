package parser

import (
	"ziniki.org/deployer/deployer/internal/repo"
	"ziniki.org/deployer/deployer/pkg/deployer"
)

type ScopedHandlers struct {
	repo repo.Repository
}

func (sh *ScopedHandlers) FindVerb(v *IdentifierToken) Action {
	if v.Id == "target" {
		return &CoreTarget{} // TODO: this should come from CoreMod
	}
	return nil
}

type CoreTarget struct {
}

func (t *CoreTarget) Handle(repo repo.Repository, tokens []Token) {
	t1 := tokens[1].(*IdentifierToken)
	repo.IntroduceSymbol(t1.Loc, deployer.SymbolType("core.Target"), deployer.SymbolName(t1.Id))
}

func NewScopeHandlers(repo repo.Repository) Scoper {
	return &ScopedHandlers{repo: repo}
}
