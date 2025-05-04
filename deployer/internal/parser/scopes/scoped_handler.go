package scopes

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopedHandlers struct {
	tools *pluggable.Tools
}

func (sh *ScopedHandlers) FindTopCommand(v pluggable.Identifier) pluggable.TopCommand {
	return sh.tools.Recall.Find(reflect.TypeFor[pluggable.TopCommand](), v.Id()).(pluggable.TopCommand)
}

func NewScopedHandlers(tools *pluggable.Tools) pluggable.Scoper {
	ret := &ScopedHandlers{tools: tools}
	return ret
}
