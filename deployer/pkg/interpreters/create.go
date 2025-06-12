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

func NewVerbCommandInterpreter(tools *pluggable.Tools, parent pluggable.AttachResult, forExtensionPoint string, allowAssign bool) pluggable.Interpreter {
	return interpreters.NewVerbCommandInterpreter(tools, parent, forExtensionPoint, allowAssign)
}

func NewPropertiesInnerScope(tools *pluggable.Tools, parent pluggable.PropertyParent) pluggable.Interpreter {
	return interpreters.NewPropertiesInnerScope(tools, parent)
}
