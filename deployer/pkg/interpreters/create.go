package interpreters

import (
	"ziniki.org/deployer/deployer/internal/parser/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func NewIgnoreInnerScope() pluggable.Interpreter {
	return interpreters.NewIgnoreInnerScope()
}

func NewDisallowInnerScope(tools *pluggable.Tools) pluggable.Interpreter {
	return interpreters.NewDisallowInnerScope(tools)
}

func NewPropertiesInnerScope(tools *pluggable.Tools, parent pluggable.PropertyParent) pluggable.Interpreter {
	return interpreters.NewPropertiesInnerScope(tools, parent)
}
