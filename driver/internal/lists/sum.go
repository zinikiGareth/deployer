package lists

import (
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type SumListExpr struct {
	driverbottom.Locatable
	exprs []driverbottom.Expr
}

func (sle *SumListExpr) ShortDescription() string {
	return ""
}

func (sle *SumListExpr) DumpTo(iw driverbottom.IndentWriter) {
}

func (sle *SumListExpr) String() string {
	return ""
}

func (sle *SumListExpr) Resolve(r driverbottom.Resolver) {

}

func (sle *SumListExpr) Eval(storage driverbottom.RuntimeStorage) any {
	var sum float64 = 0
	for _, e := range sle.exprs {
		q := storage.Eval(e)
		sum += sle.SumIt(q)
	}
	return sum
}

func (sle *SumListExpr) SumIt(q any) float64 {
	var sum float64 = 0
	switch q := q.(type) {
	case []any:
		for _, x := range q {
			sum += sle.SumIt(x)
		}
	case float64:
		sum += q
	default:
		log.Fatalf("%T", q)
	}
	return sum
}

type SumListFunc struct {
	tools *driverbottom.CoreTools
}

func (i *SumListFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_PREFIX
}

func (h *SumListFunc) Associativity() bool {
	return false
}

func (h *SumListFunc) Precedence() int {
	return 1 // we want everything else done first
}

func (h *SumListFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	rep := h.tools.Reporter
	if len(before) != 0 {
		panic("bug in fixity")
	}
	if len(after) < 1 {
		rep.Report(me.Loc().Offset, "sum expr...")
		return nil
	}
	return &SumListExpr{Locatable: me, exprs: after}
}

func MakeSumFunc(tools *driverbottom.CoreTools) any {
	return &SumListFunc{tools: tools}
}

var _ driverbottom.Function = &SumListFunc{}
