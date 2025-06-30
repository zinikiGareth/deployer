package interpreters

import (
	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type collectMapInterpreter struct {
	tools  *driverbottom.CoreTools
	parent driverbottom.PropertyParent
	prop   driverbottom.Identifier
	pairs  []driverbottom.MapEntry
}

func (cmi *collectMapInterpreter) HaveTokens(tokens []driverbottom.Token) driverbottom.Interpreter {
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

	expr, ok := cmi.tools.Parser.Parse(tokens[2:])
	if !ok {
		return NewIgnoreInnerScope()
	}
	cmi.pairs = append(cmi.pairs, exprs.NewMapPair(prop, expr))

	return NewDisallowInnerScope(cmi.tools)
}

func (cli *collectMapInterpreter) Completed() {
	cli.parent.AddProperty(cli.prop, exprs.NewMapExpr(cli.pairs))
	cli.parent.Completed()
}
