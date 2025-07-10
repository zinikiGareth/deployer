package interpreters

import (
	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type collectMapInterpreter struct {
	tools  *driverbottom.CoreTools
	loc    *errorsink.Location
	parent driverbottom.PropertyParent
	prop   driverbottom.Identifier
	addTo  driverbottom.ValueParent
	pairs  []driverbottom.MapEntry
}

// AddAdverb implements driverbottom.PropertyParent.
func (cmi *collectMapInterpreter) AddAdverb(adverb driverbottom.Adverb, args []driverbottom.Token) driverbottom.Interpreter {
	panic("unimplemented - and should not be called")
}

// AddProperty implements driverbottom.PropertyParent.
func (cmi *collectMapInterpreter) AddProperty(name driverbottom.Identifier, expr driverbottom.Expr) {
	cmi.pairs = append(cmi.pairs, exprs.NewMapPair(name.Loc(), name, expr))
}

func (cmi *collectMapInterpreter) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 3 {
		cmi.tools.Reporter.Report(0, "line must be of form <key> <- <expr>")
		return NewDisallowInnerScope(cmi.tools)
	}

	var prop driverbottom.Identifier
	switch t := tokens[0].(type) {
	case driverbottom.Identifier:
		prop = t
	case driverbottom.String:
		prop = lexicator.NewIdentifierToken(t.Loc().Line, t.Loc().Offset, t.Text())
	default:
		cmi.tools.Reporter.Report(t.Loc().Offset, "identifier or string expected")
	}

	op, ok := tokens[1].(driverbottom.Operator)
	if !ok {
		cmi.tools.Reporter.Reportf(tokens[1].Loc().Offset, "property <- expr")
		return NewIgnoreInnerScope()
	} else if !op.Is("<-") {
		cmi.tools.Reporter.Reportf(tokens[1].Loc().Offset, "property <- expr")
		return NewIgnoreInnerScope()
	}

	expr, ok := cmi.tools.Parser.Parse(scope, tokens[2:])
	if !ok {
		return NewIgnoreInnerScope()
	}
	switch expr := expr.(type) {
	case *exprs.ListExpr:
		if expr.IsEmpty() {
			return NewCollectListInnerScope(prop.Loc(), cmi.tools, cmi, prop, nil)
		}
	case *exprs.MapExpr:
		if expr.IsEmpty() {
			return NewCollectMapInnerScope(prop.Loc(), cmi.tools, cmi, prop, nil)
		}
	}

	cmi.pairs = append(cmi.pairs, exprs.NewMapPair(prop.Loc(), prop, expr))
	return NewDisallowInnerScope(cmi.tools)
}

func (cmi *collectMapInterpreter) Completed() {
	val := exprs.NewMapExpr(cmi.loc, cmi.pairs)
	if cmi.addTo != nil {
		cmi.addTo.Add(val)
		// cli.addTo.Completed()
	} else {
		cmi.parent.AddProperty(cmi.prop, val)
		// cli.parent.Completed()
	}
}
