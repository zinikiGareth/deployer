package exprs_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type myRecall struct {
	things map[string]any
}

type returnDataValue struct {
	value pluggable.Expr
}

func (rdv returnDataValue) Eval(me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	return rdv.value
}

type konstantFunc struct {
}

func (rdv konstantFunc) Eval(me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
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

func (m myRecall) Find(ty reflect.Type, noun string) any {
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
