package interpreters

import (
	"log"

	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type propertiesInterpreter struct {
	tools  *driverbottom.CoreTools
	parent driverbottom.PropertyParent
}

func (pis *propertiesInterpreter) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) >= 1 {
		if adv, ok := tokens[0].(driverbottom.Adverb); ok {
			return pis.parent.AddAdverb(adv, tokens[1:])
		}
	}

	if len(tokens) < 3 {
		pis.tools.Reporter.Report(0, "line must be of form <prop> <- <expr>, <prop> <= <interpreter> or @adverb ...")
		return NewIgnoreInnerScope()
	}

	prop, ok := tokens[0].(driverbottom.Identifier)
	if !ok {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property must be an identifier")
		return NewIgnoreInnerScope()
	}

	op, ok := tokens[1].(driverbottom.Operator)
	if !ok {
		pis.tools.Reporter.Report(tokens[1].Loc().Offset, "line must be of form <prop> <- <expr>, <prop> <= <interpreter> or @adverb ...")
		return NewIgnoreInnerScope()
	} else if op.Is("<=") {
		if len(tokens) != 3 {
			pis.tools.Reporter.Report(tokens[1].Loc().Offset, "line must be of form <prop> <- <expr>, <prop> <= <interpreter> or @adverb ...")
			return NewIgnoreInnerScope()
		}
		intname, ok := tokens[2].(driverbottom.Identifier)
		if !ok {
			pis.tools.Reporter.Report(tokens[2].Loc().Offset, "property interpreter must be an identifier")
			return NewIgnoreInnerScope()
		}
		useInner := pis.tools.Recall.Find("prop-interpreter", intname.Id())
		if useInner == nil {
			pis.tools.Reporter.Reportf(tokens[2].Loc().Offset, "there is no property interpreter %s", intname.Id())
			return NewIgnoreInnerScope()
		}
		createInt, ok := useInner.(driverbottom.CreateInterpreter)
		if !ok {
			pis.tools.Reporter.Reportf(tokens[2].Loc().Offset, "%s is not a property interpreter creator but %T", intname.Id(), useInner)
			return NewIgnoreInnerScope()
		}
		log.Printf("have %T", createInt)
		return createInt(pis.tools, pis.parent, prop)
	} else if !op.Is("<-") {
		pis.tools.Reporter.Report(tokens[1].Loc().Offset, "line must be of form <prop> <- <expr>, <prop> <= <interpreter> or @adverb ...")
		return NewIgnoreInnerScope()
	}

	expr, ok := pis.tools.Parser.Parse(scope, tokens[2:])
	if !ok {
		return NewIgnoreInnerScope()
	}

	switch expr := expr.(type) {
	case *exprs.ListExpr:
		if expr.IsEmpty() {
			return NewCollectListInnerScope(pis.tools, pis.parent, prop, nil)
		}
	case *exprs.MapExpr:
		if expr.IsEmpty() {
			return NewCollectMapInnerScope(pis.tools, pis.parent, prop, nil)
		}
	}
	pis.parent.AddProperty(prop, expr)
	return NewDisallowInnerScope(pis.tools)
}

func (pis *propertiesInterpreter) Completed() {
	pis.parent.Completed()
}
