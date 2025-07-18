package basicmath

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type binop struct {
	driverbottom.Locatable
	opname string
	lhs    driverbottom.Expr
	rhs    driverbottom.Expr
}

type multiplyExpr struct {
	binop
}

func (m *binop) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (m *binop) ShortDescription() string {
	return fmt.Sprintf("%s [%s,%s]", m.opname, m.lhs.ShortDescription(), m.rhs.ShortDescription())
}

func (m *binop) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (m *binop) String() string {
	panic("unimplemented")
}

func (m *binop) Resolve(r driverbottom.Resolver) {
	panic("unimplemented")
}

func (m *multiplyExpr) Eval(s driverbottom.RuntimeStorage) any {
	panic("unimplemented")
}

type multiplyFunc struct {
	tools *driverbottom.CoreTools
}

func (i *multiplyFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	if len(before) != 1 {
		panic("should be an error")
	}
	if len(after) != 1 {
		panic("should be an error")
	}
	return &multiplyExpr{binop: binop{Locatable: me, opname: "mult", lhs: before[0], rhs: after[0]}}
}

func MakeMultiplyFunc(tools *driverbottom.CoreTools) *multiplyFunc {
	return &multiplyFunc{tools: tools}
}

var _ driverbottom.Function = &multiplyFunc{}
var _ driverbottom.Expr = &multiplyExpr{}
