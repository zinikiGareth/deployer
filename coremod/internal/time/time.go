package time

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

const (
	HOURS = "HOURS"
)

type TimeOf struct {
	driverbottom.Locatable
	Number int
	Unit   string
}

func (t *TimeOf) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	return driverbottom.MAY_BE_BOUND
}

func (t *TimeOf) Eval(s driverbottom.RuntimeStorage) any {
	return t
}

func (t *TimeOf) ShortDescription() string {
	return fmt.Sprintf("TimeOf[%d %s]", t.Number, t.Unit)
}

func (t *TimeOf) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("TimeOf")
	iw.AttrsWhere(t)
	iw.TextAttr("number", fmt.Sprintf("%d", t.Number))
	iw.TextAttr("unit", t.Unit)
	iw.EndAttrs()
}

func (t TimeOf) String() string {
	return fmt.Sprintf("%s Time[%d,%s]", t.Locatable.Loc(), t.Number, t.Unit)
}

type HoursFunc struct {
	tools *corebottom.Tools
}

func (i *HoursFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_POSTFIX
}

func (h *HoursFunc) Associativity() bool {
	return false
}

func (h *HoursFunc) Precedence() int {
	return 1 // we want everything else done first
}

func (h *HoursFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	rep := h.tools.Reporter
	if len(before) != 1 || len(after) != 0 {
		rep.Report(me.Loc().Offset, "<nn> hours")
		return nil, false
	}
	value := before[0]
	konst, ok := value.(driverbottom.Number)
	if !ok {
		panic("not implemented: not-const hours")
	}
	return &TimeOf{Locatable: konst, Number: int(konst.F64()), Unit: HOURS}, true
}

func MakeHoursFunc(tools *corebottom.Tools) *HoursFunc {
	return &HoursFunc{tools: tools}
}

var _ driverbottom.Function = &HoursFunc{}
