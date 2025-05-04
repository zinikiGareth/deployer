package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PropertyParent interface {
	AddProperty(name pluggable.Identifier, expr pluggable.Expr)
	Completed()
}

type propertiesInnerScope struct {
	tools  *pluggable.Tools
	parent PropertyParent
}

func (pis *propertiesInnerScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 3 {
		pis.tools.Reporter.Report(0, "<prop> <- <expr>")
		return DisallowInnerScope(pis.tools)
	}

	prop, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		panic("nice error please")
	}

	op, ok := tokens[1].(pluggable.Operator)
	if !ok {
		panic("nice error please")
	} else if !op.Is("<-") {
		panic("not <-")
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
