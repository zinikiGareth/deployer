package exprs_test

import (
	"fmt"
	"slices"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type myRecall struct {
	funcs map[string]pluggable.Function
}

type returnDataValue struct {
	value pluggable.Expr
}

func (rdv returnDataValue) Eval(tools *pluggable.Tools, me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	return rdv.value
}

type konstantFunc struct {
}

func (rdv konstantFunc) Eval(tools *pluggable.Tools, me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	return exprs.Apply{Func: rdv, Args: slices.Concat(before, after)}
}

var recall myRecall
var idFunc pluggable.Function
var konstFunc pluggable.Function
var oneString pluggable.String
var lineloc *errors.LineLoc

func init() {
	fmt.Println("init")
	lineloc = &errors.LineLoc{Line: 1, Indent: 1}
	oneString = lexicator.NewStringToken(lineloc, 0, "string_1")
	idFunc = returnDataValue{value: oneString}
	konstFunc = konstantFunc{}
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

type Helpers struct {
	Sink *testhelpers.MockSink
}

func makeParser(t *testing.T) (pluggable.ExprParser, Helpers) {
	reporter, sink := testhelpers.MockReporter(t)
	recall = myRecall{funcs: make(map[string]pluggable.Function)}
	tools := &pluggable.Tools{Reporter: reporter, Recall: recall}
	reporter.At(lineloc)
	return exprs.NewExprParser(tools), Helpers{Sink: sink}
}
