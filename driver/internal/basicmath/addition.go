package basicmath

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/utils"
)

type addExpr struct {
	binop
}

func (m *addExpr) Eval(s driverbottom.RuntimeStorage) any {
	l := m.lhs.Eval(s)
	ln, ok := utils.AsF64(l)
	if !ok {
		panic("lhs of add was not a number")
	}
	r := m.rhs.Eval(s)
	rn, ok := utils.AsF64(r)
	if !ok {
		panic("rhs of add was not a number")
	}
	return ln + rn
}

type addFunc struct {
	tools *driverbottom.CoreTools
}

func (i *addFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (i *addFunc) Precedence() int {
	return 5
}

func (i *addFunc) Associativity() bool {
	return true
}

func (i *addFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	if len(before) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "+ requires left operand")
		return nil
	}
	if len(after) != 1 {
		i.tools.Reporter.ReportAtf(me.Loc(), "+ requires right operand")
		return nil
	}
	return &addExpr{binop: binop{Locatable: me, opname: "add", lhs: before[0], rhs: after[0]}}
}

func MakeAddFunc(tools *driverbottom.CoreTools) *addFunc {
	return &addFunc{tools: tools}
}

var _ driverbottom.Function = &addFunc{}
var _ driverbottom.Expr = &addExpr{}
