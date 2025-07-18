package basicmath

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type multiplyExpr struct {
	binop
}

func (m *multiplyExpr) Eval(s driverbottom.RuntimeStorage) any {
	l := m.lhs.Eval(s)
	ln, ok := l.(float64)
	if !ok {
		panic("lhs of multiply was not a number")
	}
	r := m.rhs.Eval(s)
	rn, ok := r.(float64)
	if !ok {
		panic("rhs of multiply was not a number")
	}
	return ln * rn
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
		i.tools.Reporter.ReportAtf(me.Loc(), "* requires right operand")
		return nil
	}
	return &multiplyExpr{binop: binop{Locatable: me, opname: "mult", lhs: before[0], rhs: after[0]}}
}

func MakeMultiplyFunc(tools *driverbottom.CoreTools) *multiplyFunc {
	return &multiplyFunc{tools: tools}
}

var _ driverbottom.Function = &multiplyFunc{}
var _ driverbottom.Expr = &multiplyExpr{}
