package basicmath

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type subExpr struct {
	binop
}

func (m *subExpr) Eval(s driverbottom.RuntimeStorage) any {
	l := m.lhs.Eval(s)
	ln, ok := l.(float64)
	if !ok {
		panic("lhs of sub was not a number")
	}
	r := m.rhs.Eval(s)
	rn, ok := r.(float64)
	if !ok {
		panic("rhs of sub was not a number")
	}
	return ln - rn
}

type subFunc struct {
	tools *driverbottom.CoreTools
}

func (i *subFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (i *subFunc) Precedence() int {
	return 5
}

func (i *subFunc) Associativity() bool {
	return true
}

func (i *subFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	if len(before) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "- requires left operand")
		return nil
	}
	if len(after) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "- requires right operand")
		return nil
	}
	return &subExpr{binop: binop{Locatable: me, opname: "sub", lhs: before[0], rhs: after[0]}}
}

func MakeSubFunc(tools *driverbottom.CoreTools) *subFunc {
	return &subFunc{tools: tools}
}

var _ driverbottom.Function = &subFunc{}
var _ driverbottom.Expr = &subExpr{}
