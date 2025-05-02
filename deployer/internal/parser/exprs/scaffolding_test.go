package exprs_test

import (
	"fmt"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type myRecall struct {
	funcs map[string]pluggable.Function
}

type returnDataValue struct {
	value pluggable.Expr
}

func (rdv returnDataValue) Eval(tools *pluggable.Tools, tokens []pluggable.Expr) pluggable.Expr {
	return rdv.value
}

type konstantFunc struct {
}

func (rdv konstantFunc) Eval(tools *pluggable.Tools, tokens []pluggable.Expr) pluggable.Expr {
	return exprs.Apply{Func: rdv, Args: tokens}
}

var recall myRecall
var idFunc pluggable.Function
var konstFunc pluggable.Function
var oneString pluggable.String
var lineloc *errors.LineLoc

func init() {
	fmt.Println("init")
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

func makeParser() pluggable.ExprParser {
	recall = myRecall{funcs: make(map[string]pluggable.Function)}
	lineloc = &errors.LineLoc{}
	tools := &pluggable.Tools{Recall: recall}
	return exprs.NewExprParser(tools)
}
