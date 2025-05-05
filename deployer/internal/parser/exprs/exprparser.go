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
	if len(tokens) == 0 {
		return nil, true
	}
	blocks, ok := p.reduceParens(tokens)
	if !ok {
		return nil, false
	}
	var ret []pluggable.Expr
	for _, b := range blocks {
		expr, ok := p.Parse(b)
		if !ok {
			return nil, false
		}
		ret = append(ret, expr)
	}
	return ret, true
}

func (p *exprParser) reduceParens(tokens []pluggable.Token) ([][]pluggable.Token, bool) {
	i := 0
	var ret [][]pluggable.Token
	for i < len(tokens) {
		t := tokens[i]
		if IsPunc(t) {
			if IsPuncChar(t, '(') {
				panic("unimplemented")
			} else {
				panic("unhandled punctuation char: " + t.String())
			}
		} else {
			ret = append(ret, []pluggable.Token{t})
			i++
		}
	}
	return ret, true
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

func IsPunc(tok pluggable.Token) bool {
	_, ok := tok.(pluggable.Punc)
	return ok
}

func IsPuncChar(tok pluggable.Token, pc rune) bool {
	punc, ok := tok.(pluggable.Punc)
	if !ok {
		return false
	}
	return punc.Is(pc)
}

func NewExprParser(tools *pluggable.Tools) pluggable.ExprParser {
	return &exprParser{tools: tools}
}
