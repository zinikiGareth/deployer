package interpreters

import (
	"ziniki.org/deployer/driver/internal/parser/interpreters"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func NewIgnoreInnerScope() driverbottom.Interpreter {
	return interpreters.NewIgnoreInnerScope()
}

func NewDisallowInnerScope(tools *driverbottom.CoreTools) driverbottom.Interpreter {
	return interpreters.NewDisallowInnerScope(tools)
}

func NewVerbCommandInterpreter(tools *driverbottom.CoreTools, parent driverbottom.AttachResult, forExtensionPoint string, allowAssign bool) driverbottom.Interpreter {
	return interpreters.NewVerbCommandInterpreter(tools, parent, forExtensionPoint, allowAssign)
}

func NewPropertiesInnerScope(tools *driverbottom.CoreTools, parent driverbottom.PropertyParent) driverbottom.Interpreter {
	return interpreters.NewPropertiesInnerScope(tools, parent)
}
