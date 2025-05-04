package time

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

const (
	HOURS = "HOURS"
)

type TimeOf struct {
	pluggable.Locatable
	Number int
	Unit   string
}

func (t TimeOf) String() string {
	return fmt.Sprintf("%s Time[%d,%s]", t.Locatable.Loc(), t.Number, t.Unit)
}

type HoursFunc struct {
	tools *pluggable.Tools
}

func (h *HoursFunc) Eval(me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	rep := h.tools.Reporter
	if len(before) != 1 || len(after) != 0 {
		rep.Report(me.Loc().Offset, "<nn> hours")
		return nil
	}
	value := before[0]
	konst, ok := value.(pluggable.Number)
	if !ok {
		panic("not implemented: not-const hours")
	}
	return TimeOf{Locatable: konst, Number: int(konst.Value()), Unit: HOURS}
}

func MakeHoursFunc(tools *pluggable.Tools) *HoursFunc {
	return &HoursFunc{tools: tools}
}
