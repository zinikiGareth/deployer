package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type propertiesInterpreter struct {
	tools  *pluggable.Tools
	parent pluggable.PropertyParent
}

func (pis *propertiesInterpreter) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) >= 1 {
		if adv, ok := tokens[0].(pluggable.Adverb); ok {
			return pis.parent.AddAdverb(adv, tokens[1:])
		}
	}

	if len(tokens) < 3 {
		pis.tools.Reporter.Report(0, "<prop> <- <expr>")
		return NewDisallowInnerScope(pis.tools)
	}

	prop, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property must be an identifier")
		return NewIgnoreInnerScope()
	}

	op, ok := tokens[1].(pluggable.Operator)
	if !ok {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property <- expr")
		return NewIgnoreInnerScope()
	} else if !op.Is("<-") {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property <- expr")
		return NewIgnoreInnerScope()
	}

	expr, ok := pis.tools.Parser.Parse(tokens[2:])
	if !ok {
		return NewIgnoreInnerScope()
	}
	pis.parent.AddProperty(prop, expr)
	return NewDisallowInnerScope(pis.tools)
}

func (pis *propertiesInterpreter) Completed() {
	pis.parent.Completed()
}
