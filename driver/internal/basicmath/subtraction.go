package basicmath

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/utils"
)

type subExpr struct {
	binop
}

func (m *subExpr) Eval(s driverbottom.RuntimeStorage) any {
	l := m.lhs.Eval(s)
	ln, ok := utils.AsF64(l)
	if !ok {
		panic("lhs of sub was not a number")
	}
	r := m.rhs.Eval(s)
	rn, ok := utils.AsF64(r)
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

func (i *subFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	if len(before) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "- requires left operand")
		return nil, false
	}
	if len(after) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "- requires right operand")
		return nil, false
	}
	return &subExpr{binop: binop{Locatable: me, opname: "sub", lhs: before[0], rhs: after[0]}}, true
}

func MakeSubFunc(tools *driverbottom.CoreTools) *subFunc {
	return &subFunc{tools: tools}
}

var _ driverbottom.Function = &subFunc{}
var _ driverbottom.Expr = &subExpr{}
