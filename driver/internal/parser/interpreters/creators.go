package interpreters

import "ziniki.org/deployer/driver/pkg/driverbottom"

func NewIgnoreInnerScope() driverbottom.Interpreter {
	return &ignoreInnerScope{}
}

func NewDisallowInnerScope(tools *driverbottom.CoreTools) driverbottom.Interpreter {
	return &disallowInnerScope{tools: tools}
}

func NewPropertiesInnerScope(tools *driverbottom.CoreTools, parent driverbottom.PropertyParent) driverbottom.Interpreter {
	return &propertiesInterpreter{tools: tools, parent: parent}
}

func NewCollectListInnerScope(tools *driverbottom.CoreTools, parent driverbottom.PropertyParent, prop driverbottom.Identifier) driverbottom.Interpreter {
	return &collectListInterpreter{tools: tools, parent: parent, prop: prop}
}

func NewCollectMapInnerScope(tools *driverbottom.CoreTools, parent driverbottom.PropertyParent, prop driverbottom.Identifier) driverbottom.Interpreter {
	return &collectMapInterpreter{tools: tools, parent: parent, prop: prop}
}
