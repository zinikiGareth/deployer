package exprs

import (
	"fmt"
	"reflect"
	"strings"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ParenReduction interface {
	ReduceParens(tokens []pluggable.Token) ([]pluggable.Token, bool)
}

type Bracketed struct {
	Tokens []pluggable.Token
}

func (b Bracketed) Loc() *errors.Location {
	return b.Tokens[0].Loc()
}

func (b Bracketed) String() string {
	strs := make([]string, len(b.Tokens))
	for i := 0; i < len(b.Tokens); i++ {
		strs[i] = b.Tokens[i].String()
	}
	return strings.Join(strs, " ")
}

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
		return AsExpr(before[0]), true
	}
}

func AsExpr(x pluggable.Token) pluggable.Expr {
	switch x.(type) {
	case pluggable.Expr:
		return x.(pluggable.Expr)
	case pluggable.Identifier:
		return VarRefer(x.(pluggable.Identifier))
	default:
		panic(fmt.Sprintf("cannot handle type %T", x))
	}
}

func (p *exprParser) ParseMultiple(tokens []pluggable.Token) ([]pluggable.Expr, bool) {
	if len(tokens) == 0 {
		return nil, true
	}
	blocks, ok := p.ReduceParens(tokens)
	if !ok {
		return nil, false
	}
	var ret []pluggable.Expr
	for _, b := range blocks {
		var bs []pluggable.Token
		brack, ok := b.(Bracketed)
		if ok {
			bs = brack.Tokens
		} else {
			bs = []pluggable.Token{b}
		}
		expr, ok := p.Parse(bs)
		if !ok {
			return nil, false
		}
		ret = append(ret, expr)
	}
	return ret, true
}

func (p *exprParser) ReduceParens(tokens []pluggable.Token) ([]pluggable.Token, bool) {
	i := 0
	var ret []pluggable.Token
	ret, i = p.ScanLoop(tokens, ret, i, ' ')
	if i != len(tokens) {
		return nil, false
	}
	return ret, true
}

func (p *exprParser) ScanFor(tokens []pluggable.Token, i int, end rune) ([]pluggable.Token, int) {
	ret, j := p.ScanLoop(tokens, []pluggable.Token{tokens[i]}, i+1, end)
	if len(ret) < 1 || !IsPuncChar(ret[len(ret)-1], end) {
		p.tools.Reporter.Reportf(tokens[i].Loc().Offset, "did not find matching %c", end)
		return nil, -1
	}
	return ret, j
}

func (p *exprParser) ScanLoop(tokens []pluggable.Token, ret []pluggable.Token, i int, end rune) ([]pluggable.Token, int) {
	for i < len(tokens) {
		t := tokens[i]
		if IsPunc(t) {
			if IsPuncChar(t, end) {
				ret = append(ret, t)
				i++
				return ret, i
			}
			if IsPuncChar(t, '(') {
				inner, j := p.ScanFor(tokens, i, ')')
				if j == -1 {
					return nil, -1
				}
				ret = append(ret, Bracketed{Tokens: inner})
				i = j
			} else {
				p.tools.Reporter.Reportf(tokens[i].Loc().Offset, "unexpected close paren: %c", t.(pluggable.Punc).Which())
				return nil, -1
			}
		} else {
			ret = append(ret, t)
			i++
		}
	}
	return ret, i
}

func makeArgs(tokens []pluggable.Token) []pluggable.Expr {
	args := make([]pluggable.Expr, len(tokens))
	for k, tok := range tokens {
		args[k] = AsExpr(tok)
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
