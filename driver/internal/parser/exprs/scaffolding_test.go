package exprs_test

import (
	"slices"
	"testing"

	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type myRecall struct {
	things map[string]any
}

func (m *myRecall) ExtensionPoint(name string) {
	panic("unimplemented")
}

func (m *myRecall) ProvideDriver(s string, env any) {
	panic("unimplemented")
}

func (m *myRecall) Register(_ string, called string, item any) {
	m.things[called] = item
}

type returnDataValue struct {
	value driverbottom.Expr
}

// Precedence implements driverbottom.Function.
func (rdv returnDataValue) Precedence() int {
	return 10
}

func (rdv returnDataValue) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	return rdv.value
}

type konstantFunc struct {
}

// Precedence implements driverbottom.Function.
func (rdv konstantFunc) Precedence() int {
	return 8
}

func (rdv konstantFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	return &exprs.Apply{Func: rdv, Args: slices.Concat(before, after)}
}

var recall myRecall
var idFunc driverbottom.Function
var konstFunc driverbottom.Function
var oneString driverbottom.String
var lineloc *errorsink.LineLoc
var orb, crb driverbottom.Punc
var osb, csb driverbottom.Punc
var comma driverbottom.Punc

func init() {
	lineloc = &errorsink.LineLoc{Line: 1, Indent: 1, Text: "", File: &errorsink.FileLoc{File: "test"}}
	oneString = lexicator.NewStringToken(lineloc, 0, "string_1")
	idFunc = returnDataValue{value: oneString}
	konstFunc = konstantFunc{}
	orb = lexicator.NewPuncToken(lineloc, 0, '(')
	crb = lexicator.NewPuncToken(lineloc, 12, ')')
	osb = lexicator.NewPuncToken(lineloc, 2, '[')
	csb = lexicator.NewPuncToken(lineloc, 10, ']')
	comma = lexicator.NewPuncToken(lineloc, 10, ',')
}

func (m *myRecall) Find(_ string, noun string) any {
	return m.things[noun]
}

func (m *myRecall) ObtainDriver(driver string) any {
	panic("unimplemented")
}

var _ driverbottom.Register = &myRecall{}

type Helpers struct {
	Tools *driverbottom.CoreTools
	Sink  *testhelpers.MockSink
	Lex   lexicator.Lexicator
}

func makeParser(t *testing.T) (driverbottom.ExprParser, Helpers) {
	reporter, sink := testhelpers.MockReporter(t)
	recall = myRecall{things: make(map[string]any)}
	tools := &driverbottom.CoreTools{Reporter: reporter, Register: &recall, Recall: &recall}
	ll := lexicator.NewLineLexicator(tools, "test")
	reporter.At(lineloc)
	return exprs.NewExprParser(tools), Helpers{Sink: sink, Lex: ll, Tools: tools}
}
