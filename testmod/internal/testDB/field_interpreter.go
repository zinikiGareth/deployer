package testDB

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type FieldInterpreter struct {
	tools  *driverbottom.CoreTools
	parent driverbottom.PropertyParent
	prop   driverbottom.Identifier

	model []driverbottom.Expr
}

func (f *FieldInterpreter) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 3 {
		f.tools.Reporter.Report(tokens[0].Loc().Offset, "<Key|Value> <name> <type>")
		return drivertop.NewIgnoreInnerScope()
	}
	id, ok := tokens[0].(driverbottom.Identifier)
	if !ok {
		f.tools.Reporter.Report(tokens[0].Loc().Offset, "field kind must be Key or Value")
		return drivertop.NewIgnoreInnerScope()
	}
	switch id.Id() {
	case "Key":
	case "Value":
	default:
		f.tools.Reporter.Report(tokens[0].Loc().Offset, "field kind must be Key or Value")
		return drivertop.NewIgnoreInnerScope()
	}
	name, ok := tokens[1].(driverbottom.Identifier)
	if !ok {
		f.tools.Reporter.Report(tokens[1].Loc().Offset, "field name must be an Identifier")
		return drivertop.NewIgnoreInnerScope()
	}
	fty, ok := tokens[2].(driverbottom.Identifier)
	if !ok {
		f.tools.Reporter.Report(tokens[2].Loc().Offset, "field type must be an Identifier")
		return drivertop.NewIgnoreInnerScope()
	}
	field := &FieldExpr{loc: id.Loc(), kind: id.Id(), name: name.Id(), ftype: fty.Id()}
	f.model = append(f.model, field)
	return drivertop.NewDisallowInnerScope(f.tools)
}

func (f *FieldInterpreter) Completed() {
	expr := drivertop.NewListExpr(f.prop.Loc(), f.model)
	f.parent.AddProperty(f.prop, expr)
}

func CreateFieldInterpreter(tools *driverbottom.CoreTools, scope driverbottom.Scope, parent driverbottom.PropertyParent, prop driverbottom.Identifier, tokens []driverbottom.Token) driverbottom.Interpreter {
	return &FieldInterpreter{tools: tools, parent: parent, prop: prop}
}

var _ driverbottom.Interpreter = &FieldInterpreter{}
