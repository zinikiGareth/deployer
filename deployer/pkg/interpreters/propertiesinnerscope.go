package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PropertyParent interface {
	AddProperty(name pluggable.Identifier, expr any) //TODO: should probably be "Expr"
}

type propertiesInnerScope struct {
	parent PropertyParent
}

func (pis *propertiesInnerScope) HaveTokens(reporter errors.ErrorRepI, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) != 3 {
		reporter.Report(0, "invalid property yada yada")
		return DisallowInnerScope()
	}
	pis.parent.AddProperty(tokens[0].(pluggable.Identifier), tokens[2])
	return DisallowInnerScope()
}

func PropertiesInnerScope(parent PropertyParent) pluggable.Interpreter {
	return &propertiesInnerScope{parent: parent}
}
