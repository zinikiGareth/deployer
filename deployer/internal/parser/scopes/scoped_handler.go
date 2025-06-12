package scopes

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopedHandlers struct {
	tools *pluggable.Tools
}

func (sh *ScopedHandlers) FindVerbCommand(v pluggable.Identifier) pluggable.VerbCommand {
	cmd := sh.tools.Recall.Find(reflect.TypeFor[pluggable.VerbCommand](), v.Id())
	if cmd == nil {
		sh.tools.Reporter.Reportf(v.Loc().Offset, "there is no top-level command %s", v.Id())
		return nil
	}
	tc, ok := cmd.(pluggable.VerbCommand)
	if !ok {
		sh.tools.Reporter.Reportf(v.Loc().Offset, "%s: not a top-level command", v.Id())
		return nil
	}
	return tc
}

func NewScopedHandlers(tools *pluggable.Tools) pluggable.Scoper {
	ret := &ScopedHandlers{tools: tools}
	return ret
}
