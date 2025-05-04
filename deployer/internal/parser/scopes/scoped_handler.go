package scopes

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopedHandlers struct {
	repo   pluggable.Repository
	recall pluggable.Recall
}

func (sh *ScopedHandlers) FindAction(v pluggable.Identifier) pluggable.TargetCommand {
	return sh.recall.Find(reflect.TypeFor[pluggable.TargetCommand](), v.Id()).(pluggable.TargetCommand)
}

func NewScopedHandlers(registry pluggable.Recall, repo pluggable.Repository) pluggable.Scoper {
	ret := &ScopedHandlers{repo: repo, recall: registry}
	return ret
}
