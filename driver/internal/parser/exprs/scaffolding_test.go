package exprs_test

import (
	"fmt"
	"slices"
	"testing"

	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
	"ziniki.org/deployer/driver/pkg/utils"
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

// Fixity implements driverbottom.Function.
func (rdv returnDataValue) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

// Precedence implements driverbottom.Function.
func (rdv returnDataValue) Precedence() int {
	return 10
}

func (rdv returnDataValue) Associativity() bool {
	return true
}

func (rdv returnDataValue) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	return rdv.value
}

type konstExpr struct {
	driverbottom.Locatable
	args []driverbottom.Expr
}

func (ke *konstExpr) ShortDescription() string {
	ret := "<konst>"
	ret += "("
	for k, e := range ke.args {
		if k > 0 {
			ret += ","
		}
		ret += e.ShortDescription()
	}
	ret += ")"
	return ret
}

func (ke *konstExpr) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (ke *konstExpr) Resolve(r driverbottom.Resolver) {
	panic("unimplemented")
}

func (ke *konstExpr) String() string {
	return fmt.Sprintf("konst[%d]", len(ke.args))
}

func (ke *konstExpr) Eval(s driverbottom.RuntimeStorage) any {
	panic("unimplemented")
}

type konstantFunc struct {
}

func (pff *konstantFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (kf *konstantFunc) Precedence() int {
	return 8
}

func (kf *konstantFunc) Associativity() bool {
	return true
}

func (kf *konstantFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	return &konstExpr{args: slices.Concat(before, after)}
}

type facExpr struct {
	driverbottom.Locatable
	arg driverbottom.Expr
}

func (fe *facExpr) ShortDescription() string {
	return fmt.Sprintf("<fac>(%s)", fe.arg.ShortDescription())
}

func (fe *facExpr) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (fe *facExpr) Resolve(r driverbottom.Resolver) {
	panic("unimplemented")
}

func (fe *facExpr) String() string {
	return fmt.Sprintf("fac[%s]", fe.arg.String())
}

func (fe *facExpr) Eval(s driverbottom.RuntimeStorage) any {
	arg := fe.arg.Eval(s)
	i32, ok := utils.AsI32(arg)
	if !ok {
		panic("could not convert to int")
	}
	if i32 <= 0 {
		return 0
	}
	ret := 1
	for i32 > 1 {
		ret *= int(i32)
		i32--
	}
	return ret
}

type postFacFunc struct {
}

func (pff *postFacFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_POSTFIX
}

func (pff *postFacFunc) Precedence() int {
	return 7
}

func (pff *postFacFunc) Associativity() bool {
	return true
}

func (pff *postFacFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	if len(before) != 1 {
		panic("need prefix arg")
	}
	if len(after) != 0 {
		panic("have postfix arg")
	}
	return &facExpr{arg: before[0]}
}

var recall myRecall
var idFunc driverbottom.Function
var postFac driverbottom.Function
var konstFunc driverbottom.Function
var oneString driverbottom.String
var lineloc *errorsink.LineLoc
var orb, crb driverbottom.Punc
var osb, csb driverbottom.Punc
var comma driverbottom.Punc

func init() {
	lineloc = &errorsink.LineLoc{Line: 1, Indent: 1, Text: "", File: &errorsink.FileLoc{File: "test"}}
	oneString = lexicator.NewStringToken(lineloc, 0, "string_1")
	postFac = &postFacFunc{}
	idFunc = returnDataValue{value: oneString}
	konstFunc = &konstantFunc{}
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
