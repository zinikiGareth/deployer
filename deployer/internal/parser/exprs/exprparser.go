package exprs

import (
	"reflect"

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
	tok, fn, before, after := p.split(tokens)
	if fn != nil {
		return fn.Eval(tok, makeArgs(before), makeArgs(after)), true
	} else {
		if len(before) > 1 {
			p.tools.Reporter.Reportf(before[0].Loc().Offset, "no function found")
			return nil, false
		}
		return before[0], true
	}
}

func (p *exprParser) ParseMultiple(tokens []pluggable.Token) ([]pluggable.Expr, bool) {
	return nil, true
}

func makeArgs(tokens []pluggable.Token) []pluggable.Expr {
	args := make([]pluggable.Expr, len(tokens))
	for k, tok := range tokens {
		args[k] = tok
	}
	return args
}

func (p *exprParser) split(tokens []pluggable.Token) (pluggable.Token, pluggable.Function, []pluggable.Token, []pluggable.Token) {
	for i, t := range tokens {
		if f := p.matchFunc(t); f != nil {
			return t, f, tokens[0:i], tokens[i+1:]
		}
	}
	return nil, nil, tokens, nil
}

func (p *exprParser) matchFunc(tok pluggable.Token) pluggable.Function {
	id, isId := tok.(pluggable.Identifier)
	if isId {
		v, ok := p.tools.Recall.Find(reflect.TypeFor[pluggable.Function](), id.Id()).(pluggable.Function)
		if ok && v != nil {
			return v
		}
	}
	op, isOp := tok.(pluggable.Operator)
	if isOp {
		v, ok := p.tools.Recall.Find(reflect.TypeFor[pluggable.Function](), op.Op()).(pluggable.Function)
		if ok && v != nil {
			return v
		}
	}
	return nil
}

func NewExprParser(tools *pluggable.Tools) pluggable.ExprParser {
	return &exprParser{tools: tools}
}
