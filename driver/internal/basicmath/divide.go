package basicmath

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/utils"
)

type divExpr struct {
	binop
}

func (m *divExpr) Eval(s driverbottom.RuntimeStorage) any {
	l := m.lhs.Eval(s)
	ln, ok := utils.AsF64(l)
	if !ok {
		panic("lhs of divide was not a number")
	}
	r := m.rhs.Eval(s)
	rn, ok := utils.AsF64(r)
	if !ok {
		panic("rhs of divide was not a number")
	}
	return ln / rn
}

type divFunc struct {
	tools *driverbottom.CoreTools
}

func (i *divFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (i *divFunc) Precedence() int {
	return 6
}

func (i *divFunc) Associativity() bool {
	return true
}

func (i *divFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	if len(before) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "/ requires left operand")
		return nil, false
	}
	if len(after) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "/ requires right operand")
		return nil, false
	}
	return &divExpr{binop: binop{Locatable: me, opname: "div", lhs: before[0], rhs: after[0]}}, true
}

func MakeDivFunc(tools *driverbottom.CoreTools) *divFunc {
	return &divFunc{tools: tools}
}

var _ driverbottom.Function = &divFunc{}
var _ driverbottom.Expr = &divExpr{}
