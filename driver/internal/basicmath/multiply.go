package basicmath

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type multiplyExpr struct {
	binop
}

func (m *multiplyExpr) Eval(s driverbottom.RuntimeStorage) any {
	panic("unimplemented")
}

type multiplyFunc struct {
	tools *driverbottom.CoreTools
}

func (i *multiplyFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	if len(before) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "* requires left operand")
		return nil
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
