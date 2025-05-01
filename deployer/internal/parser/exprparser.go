package parser

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type exprParser struct {
	tools *pluggable.Tools
}

func (p *exprParser) Parse(tokens []pluggable.Token) pluggable.Expr {
	t := tokens[0]
	id, isId := t.(pluggable.Identifier)
	if isId {
		v := p.tools.Recall.FindFunc(id.Id())
		if v != nil {
			return v.Eval(p.tools, tokens)
		}
	}
	return tokens[0]
}

func NewExprParser(tools *pluggable.Tools) pluggable.ExprParser {
	return &exprParser{tools: tools}
}
