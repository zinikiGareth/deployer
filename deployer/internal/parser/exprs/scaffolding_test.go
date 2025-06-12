package exprs_test

import (
	"slices"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type myRecall struct {
	things map[string]any
}

type returnDataValue struct {
	value pluggable.Expr
}

func (rdv returnDataValue) ReduceExpr(me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	return rdv.value
}

type konstantFunc struct {
}

func (rdv konstantFunc) ReduceExpr(me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	return &exprs.Apply{Func: rdv, Args: slices.Concat(before, after)}
}

var recall myRecall
var idFunc pluggable.Function
var konstFunc pluggable.Function
var oneString pluggable.String
var lineloc *errorsink.LineLoc
var orb, crb pluggable.Punc

func init() {
	lineloc = &errorsink.LineLoc{Line: 1, Indent: 1, Text: "", File: &errorsink.FileLoc{File: "test"}}
	oneString = lexicator.NewStringToken(lineloc, 0, "string_1")
	idFunc = returnDataValue{value: oneString}
	konstFunc = konstantFunc{}
	orb = lexicator.NewPuncToken(lineloc, 0, '(')
	crb = lexicator.NewPuncToken(lineloc, 12, ')')
}

func (m myRecall) Find(_ string, noun string) any {
	return m.things[noun]
}

func (m myRecall) ObtainDriver(driver string) any {
	panic("unimplemented")
}

type Helpers struct {
	Sink *testhelpers.MockSink
}

func makeParser(t *testing.T) (pluggable.ExprParser, Helpers) {
	reporter, sink := testhelpers.MockReporter(t)
	recall = myRecall{things: make(map[string]any)}
	tools := &pluggable.Tools{Reporter: reporter, Recall: recall}
	reporter.At(lineloc)
	return exprs.NewExprParser(tools), Helpers{Sink: sink}
}
