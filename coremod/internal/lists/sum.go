package lists

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
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
	return nil
}

type SumListFunc struct {
	tools *corebottom.Tools
}

func (h *SumListFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	rep := h.tools.Reporter
	if len(before) != 0 || len(after) < 1 {
		rep.Report(me.Loc().Offset, "sum expr...")
		return nil
	}
	return &SumListExpr{Locatable: me, exprs: after}
}

func MakeSumFunc(tools *corebottom.Tools) any {
	return &SumListFunc{tools: tools}
}
