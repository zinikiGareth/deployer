package exprs

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type exprParser struct {
	tools *pluggable.Tools
}

func (p *exprParser) Parse(tokens []pluggable.Token) (pluggable.Expr, bool) {
	if len(tokens) == 0 {
		p.tools.Reporter.Reportf(0, "no expression found")
		return nil, false
	}
	fn, before, after := p.split(tokens)
	if fn != nil {
		return fn.Eval(p.tools, makeArgs(before), makeArgs(after)), true
	} else {
		if len(before) > 1 {
			p.tools.Reporter.Reportf(before[0].Loc().Offset, "no function found")
			return nil, false
		}
		return before[0], true
	}
}

func makeArgs(tokens []pluggable.Token) []pluggable.Expr {
	args := make([]pluggable.Expr, len(tokens))
	for k, tok := range tokens {
		args[k] = tok
	}
	return args
}

func (p *exprParser) split(tokens []pluggable.Token) (pluggable.Function, []pluggable.Token, []pluggable.Token) {
	for i, t := range tokens {
		if f := p.matchFunc(t); f != nil {
			return f, tokens[0:i], tokens[i+1:]
		}
	}
	return nil, tokens, nil
}

func (p *exprParser) matchFunc(tok pluggable.Token) pluggable.Function {
	id, isId := tok.(pluggable.Identifier)
	if isId {
		v := p.tools.Recall.FindFunc(id.Id())
		if v != nil {
			return v
		}
	}
	op, isOp := tok.(pluggable.Operator)
	if isOp {
		v := p.tools.Recall.FindFunc(op.Op())
		if v != nil {
			return v
		}
	}
	return nil
}

func NewExprParser(tools *pluggable.Tools) pluggable.ExprParser {
	return &exprParser{tools: tools}
}
