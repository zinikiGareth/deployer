package parser

import (
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopedHandlers struct {
	repo   pluggable.Repository
	recall registry.Recall
}

func (sh *ScopedHandlers) FindVerb(v pluggable.Identifier) pluggable.Action {
	return sh.recall.FindVerb(v.Id())
}

func NewScopedHandlers(registry registry.Recall, repo pluggable.Repository) Scoper {
	ret := &ScopedHandlers{repo: repo, recall: registry}
	return ret
}
