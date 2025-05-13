package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PropertyParent interface {
	AddProperty(name pluggable.Identifier, expr pluggable.Expr)
	AddAdverb(adverb pluggable.Adverb, args []pluggable.Token) pluggable.Interpreter
	Completed()
}

type propertiesInnerScope struct {
	tools  *pluggable.Tools
	parent PropertyParent
}

func (pis *propertiesInnerScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) >= 1 {
		if adv, ok := tokens[0].(pluggable.Adverb); ok {
			return pis.parent.AddAdverb(adv, tokens[1:])
		}
	}

	if len(tokens) < 3 {
		pis.tools.Reporter.Report(0, "<prop> <- <expr>")
		return DisallowInnerScope(pis.tools)
	}

	prop, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property must be an identifier")
		return IgnoreInnerScope()
	}

	op, ok := tokens[1].(pluggable.Operator)
	if !ok {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property <- expr")
		return IgnoreInnerScope()
	} else if !op.Is("<-") {
		pis.tools.Reporter.Reportf(tokens[0].Loc().Offset, "property <- expr")
		return IgnoreInnerScope()
	}

	expr, ok := pis.tools.Parser.Parse(tokens[2:])
	if !ok {
		return IgnoreInnerScope()
	}
	pis.parent.AddProperty(prop, expr)
	return DisallowInnerScope(pis.tools)
}

func (pis *propertiesInnerScope) Completed() {
	pis.parent.Completed()
}

func PropertiesInnerScope(tools *pluggable.Tools, parent PropertyParent) pluggable.Interpreter {
	return &propertiesInnerScope{tools: tools, parent: parent}
}
