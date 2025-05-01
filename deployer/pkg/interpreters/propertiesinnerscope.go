package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PropertyParent interface {
	AddProperty(tools *pluggable.Tools, name pluggable.Identifier, expr pluggable.Locatable) //TODO: should probably be "Expr"
	Completed(tools *pluggable.Tools)
}

type propertiesInnerScope struct {
	parent PropertyParent
}

func (pis *propertiesInnerScope) HaveTokens(tools *pluggable.Tools, tokens []pluggable.Token) pluggable.Interpreter {
	// TODO: the left hand side must be an identifier
	// then it must be "<-"
	// the rest of the tokens must be an expression
	if len(tokens) < 3 {
		tools.Reporter.Report(0, "<prop> <- <expr>")
		return DisallowInnerScope()
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

	expr, ok := parseExpr(tokens[2:])
	if !ok {
		return IgnoreInnerScope()
	}
	pis.parent.AddProperty(tools, prop, expr)
	return DisallowInnerScope()
}

// TODO: this should be elsewhere
func parseExpr(tokens []pluggable.Token) (pluggable.Locatable, bool) { // TODO: should be expr
	if len(tokens) == 1 {
		return tokens[0], true
	}
	panic("this needs to be implemented")
}

func (b *propertiesInnerScope) Completed(tools *pluggable.Tools) {
	b.parent.Completed(tools)
}

func PropertiesInnerScope(parent PropertyParent) pluggable.Interpreter {
	return &propertiesInnerScope{parent: parent}
}
