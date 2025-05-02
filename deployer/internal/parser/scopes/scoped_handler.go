package scopes

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopedHandlers struct {
	repo   pluggable.Repository
	recall pluggable.Recall
}

func (sh *ScopedHandlers) FindAction(v pluggable.Identifier) pluggable.Action {
	return sh.recall.FindAction(v.Id())
}

func NewScopedHandlers(registry pluggable.Recall, repo pluggable.Repository) pluggable.Scoper {
	ret := &ScopedHandlers{repo: repo, recall: registry}
	return ret
}
