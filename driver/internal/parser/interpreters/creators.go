package interpreters

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func NewIgnoreInnerScope() driverbottom.Interpreter {
	return &ignoreInnerScope{}
}

func NewDisallowInnerScope(tools *driverbottom.CoreTools) driverbottom.Interpreter {
	return &disallowInnerScope{tools: tools}
}

func NewPropertiesInnerScope(tools *driverbottom.CoreTools, parent driverbottom.PropertyParent) driverbottom.Interpreter {
	return &propertiesInterpreter{tools: tools, parent: parent}
}

func NewCollectListInnerScope(loc *errorsink.Location, tools *driverbottom.CoreTools, parent driverbottom.PropertyParent, prop driverbottom.Identifier, addTo driverbottom.ValueParent) driverbottom.Interpreter {
	return &collectListInterpreter{loc: loc, tools: tools, parent: parent, prop: prop, addTo: addTo}
}

func NewCollectMapInnerScope(loc *errorsink.Location, tools *driverbottom.CoreTools, parent driverbottom.PropertyParent, prop driverbottom.Identifier, addTo driverbottom.ValueParent) driverbottom.Interpreter {
	return &collectMapInterpreter{loc: loc, tools: tools, parent: parent, prop: prop, addTo: addTo}
}
