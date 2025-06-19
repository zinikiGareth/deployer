package exprs

import (
	"fmt"
	"strings"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type ParenReduction interface {
	ReduceParens(tokens []driverbottom.Token) ([]driverbottom.Token, bool)
}

type Bracketed struct {
	Tokens []driverbottom.Token
}

func (b Bracketed) Loc() *errorsink.Location {
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
	tools *driverbottom.CoreTools
}

func (p *exprParser) Parse(tokens []driverbottom.Token) (driverbottom.Expr, bool) {
	if len(tokens) == 0 {
		p.tools.Reporter.Reportf(0, "no expression found")
		return nil, false
	}
	tok, fn, before, after := p.split(tokens)
	if fn != nil {
		return fn.ReduceExpr(tok, makeArgs(before), makeArgs(after)), true
	} else {
		if len(before) > 1 {
			p.tools.Reporter.Reportf(before[0].Loc().Offset, "no function symbol found in this expression")
			return nil, false
		}
		return AsExpr(before[0]), true
	}
}

func AsExpr(x driverbottom.Token) driverbottom.Expr {
	switch x := x.(type) {
	case driverbottom.Expr:
		return x
	case driverbottom.Identifier:
		return VarRefer(x)
	default:
		panic(fmt.Sprintf("cannot handle type %T", x))
	}
}

func (p *exprParser) ParseMultiple(tokens []driverbottom.Token) ([]driverbottom.Expr, bool) {
	if len(tokens) == 0 {
		return nil, true
	}
	blocks, ok := p.ReduceParens(tokens)
	if !ok {
		return nil, false
	}
	var ret []driverbottom.Expr
	for _, b := range blocks {
		var bs []driverbottom.Token
		brack, ok := b.(Bracketed)
		if ok {
			bs = brack.Tokens
		} else {
			bs = []driverbottom.Token{b}
		}
		expr, ok := p.Parse(bs)
		if !ok {
			return nil, false
		}
		ret = append(ret, expr)
	}
	return ret, true
}

func (p *exprParser) ReduceParens(tokens []driverbottom.Token) ([]driverbottom.Token, bool) {
	i := 0
	var ret []driverbottom.Token
	ret, i = p.ScanLoop(tokens, ret, i, ' ')
	if i != len(tokens) {
		return nil, false
	}
	return ret, true
}

func (p *exprParser) ScanFor(tokens []driverbottom.Token, i int, end rune) ([]driverbottom.Token, int) {
	ret, j := p.ScanLoop(tokens, []driverbottom.Token{tokens[i]}, i+1, end)
	if len(ret) < 1 || !IsPuncChar(ret[len(ret)-1], end) {
		p.tools.Reporter.Reportf(tokens[i].Loc().Offset, "did not find matching %c", end)
		return nil, -1
	}
	return ret, j
}

func (p *exprParser) ScanLoop(tokens []driverbottom.Token, ret []driverbottom.Token, i int, end rune) ([]driverbottom.Token, int) {
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
				p.tools.Reporter.Reportf(tokens[i].Loc().Offset, "unexpected close paren: %c", t.(driverbottom.Punc).Which())
				return nil, -1
			}
		} else {
			ret = append(ret, t)
			i++
		}
	}
	return ret, i
}

func makeArgs(tokens []driverbottom.Token) []driverbottom.Expr {
	args := make([]driverbottom.Expr, len(tokens))
	for k, tok := range tokens {
		args[k] = AsExpr(tok)
	}
	return args
}

func (p *exprParser) split(tokens []driverbottom.Token) (driverbottom.Token, driverbottom.Function, []driverbottom.Token, []driverbottom.Token) {
	for i, t := range tokens {
		if f := p.matchFunc(t); f != nil {
			k := 0
			p1, ok1 := tokens[0].(driverbottom.Punc)
			p2, ok2 := tokens[len(tokens)-1].(driverbottom.Punc)
			if ok1 && ok2 && p1.Is('(') && p2.Is(')') {
				k = 1
			}
			return t, f, tokens[k:i], tokens[i+1 : len(tokens)-k]
		}
	}
	return nil, nil, tokens, nil
}

func (p *exprParser) matchFunc(tok driverbottom.Token) driverbottom.Function {
	id, isId := tok.(driverbottom.Identifier)
	if isId {
		v, ok := p.tools.Recall.Find("function-defn", id.Id()).(driverbottom.Function)
		if ok && v != nil {
			return v
		}
	}
	op, isOp := tok.(driverbottom.Operator)
	if isOp {
		v, ok := p.tools.Recall.Find("function-defn", op.Op()).(driverbottom.Function)
		if ok && v != nil {
			return v
		}
	}
	return nil
}

func IsPunc(tok driverbottom.Token) bool {
	_, ok := tok.(driverbottom.Punc)
	return ok
}

func IsPuncChar(tok driverbottom.Token, pc rune) bool {
	punc, ok := tok.(driverbottom.Punc)
	if !ok {
		return false
	}
	return punc.Is(pc)
}

func NewExprParser(tools *driverbottom.CoreTools) driverbottom.ExprParser {
	return &exprParser{tools: tools}
}
