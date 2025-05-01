package parser_test

import (
	"fmt"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type myRecall struct {
	funcs map[string]pluggable.Function
}

type returnDataValue struct {
	value pluggable.Expr
}

func (rdv returnDataValue) Eval(tools *pluggable.Tools, tokens []pluggable.Token) pluggable.Expr {
	return rdv.value
}

var recall myRecall
var makeData pluggable.Function
var oneString pluggable.String
var lineloc *errors.LineLoc

func init() {
	fmt.Println("init")
	oneString = parser.NewStringToken(lineloc, 0, "string_1")
	makeData = returnDataValue{value: oneString}
}

func (m myRecall) FindFunc(verb string) pluggable.Function {
	return m.funcs[verb]
}

func (m myRecall) FindAction(noun string) pluggable.Action {
	panic("unimplemented")
}

func (m myRecall) FindNoun(noun string) pluggable.Noun {
	panic("unimplemented")
}

func (m myRecall) ObtainDriver(driver string) any {
	panic("unimplemented")
}

func makeParser() pluggable.ExprParser {
	recall = myRecall{funcs: make(map[string]pluggable.Function)}
	lineloc = &errors.LineLoc{}
	tools := &pluggable.Tools{Recall: recall}
	return parser.NewExprParser(tools)
}

func TestAStringIsAnExpr(t *testing.T) {
	p := makeParser()
	hello := parser.NewStringToken(lineloc, 0, "hello")
	expr := p.Parse([]pluggable.Token{hello})
	if expr != hello {
		t.Fatalf("returned expr was not hello")
	}
}

func TestANumberIsAnExpr(t *testing.T) {
	p := makeParser()
	nbr := parser.NewNumberToken(lineloc, 0, 46)
	expr := p.Parse([]pluggable.Token{nbr})
	if expr != nbr {
		t.Fatalf("returned expr was not 46")
	}
}

func TestAnUnboundIDIsAnExpr(t *testing.T) {
	p := makeParser()
	id := parser.NewIdentifierToken(lineloc, 0, "x")
	expr := p.Parse([]pluggable.Token{id})
	if expr != id {
		t.Fatalf("returned expr was not x")
	}
}

func TestAnIDBoundToAVerbProducesAnExpr(t *testing.T) {
	p := makeParser()
	recall.funcs["hello"] = makeData
	id := parser.NewIdentifierToken(lineloc, 0, "hello")
	expr := p.Parse([]pluggable.Token{id})
	if expr == id {
		t.Fatalf("returned expr was the verb")
	}
	if expr != oneString {
		t.Fatalf("returned expr was not the string")
	}
}
