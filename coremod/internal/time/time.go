package time

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/external"
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

func (t *TimeOf) Resolve(r driverbottom.Resolver) {
}

func (t *TimeOf) Eval(s driverbottom.RuntimeStorage) any {
	return t
}

func (t *TimeOf) ShortDescription() string {
	panic("not implemented")
}

func (t *TimeOf) DumpTo(iw driverbottom.IndentWriter) {
	panic("not implemented")
}

func (t TimeOf) String() string {
	return fmt.Sprintf("%s Time[%d,%s]", t.Locatable.Loc(), t.Number, t.Unit)
}

type HoursFunc struct {
	tools *external.Tools
}

func (h *HoursFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	rep := h.tools.Reporter
	if len(before) != 1 || len(after) != 0 {
		rep.Report(me.Loc().Offset, "<nn> hours")
		return nil
	}
	value := before[0]
	konst, ok := value.(driverbottom.Number)
	if !ok {
		panic("not implemented: not-const hours")
	}
	return &TimeOf{Locatable: konst, Number: int(konst.Value()), Unit: HOURS}
}

func MakeHoursFunc(tools *external.Tools) *HoursFunc {
	return &HoursFunc{tools: tools}
}
