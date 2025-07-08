package testDB

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type FieldInterpreter struct {
	tools  *driverbottom.CoreTools
	parent driverbottom.PropertyParent
	prop   driverbottom.Identifier

	model driverbottom.Expr
}

func (f *FieldInterpreter) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	return drivertop.NewDisallowInnerScope(f.tools)
}

func (f *FieldInterpreter) Completed() {
	f.model = drivertop.NewAnyExpr(f.prop.Loc(), true)
	f.parent.AddProperty(f.prop, f.model)
}

func CreateFieldInterpreter(tools *driverbottom.CoreTools, parent driverbottom.PropertyParent, prop driverbottom.Identifier) driverbottom.Interpreter {
	return &FieldInterpreter{tools: tools, parent: parent, prop: prop}
}

var _ driverbottom.Interpreter = &FieldInterpreter{}
