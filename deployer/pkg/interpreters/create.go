package interpreters

import (
	"ziniki.org/deployer/deployer/internal/parser/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func NewIgnoreInnerScope() pluggable.Interpreter {
	return interpreters.NewIgnoreInnerScope()
}

func NewDisallowInnerScope(tools *pluggable.CoreTools) pluggable.Interpreter {
	return interpreters.NewDisallowInnerScope(tools)
}

func NewVerbCommandInterpreter(tools *pluggable.CoreTools, parent pluggable.AttachResult, forExtensionPoint string, allowAssign bool) pluggable.Interpreter {
	return interpreters.NewVerbCommandInterpreter(tools, parent, forExtensionPoint, allowAssign)
}

func NewPropertiesInnerScope(tools *pluggable.CoreTools, parent pluggable.PropertyParent) pluggable.Interpreter {
	return interpreters.NewPropertiesInnerScope(tools, parent)
}
