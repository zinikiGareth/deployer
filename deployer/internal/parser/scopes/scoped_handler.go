package scopes

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopedHandlers struct {
	tools *pluggable.Tools
}

func (sh *ScopedHandlers) FindAction(v pluggable.Identifier) pluggable.TargetCommand {
	return sh.tools.Recall.Find(reflect.TypeFor[pluggable.TargetCommand](), v.Id()).(pluggable.TargetCommand)
}

func NewScopedHandlers(tools *pluggable.Tools) pluggable.Scoper {
	ret := &ScopedHandlers{tools: tools}
	return ret
}
