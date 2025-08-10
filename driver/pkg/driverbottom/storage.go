package driverbottom

import (
	"fmt"
	"io"

	"ziniki.org/deployer/driver/pkg/errorsink"
)

type RuntimeStorage interface {
	Bind(name Holder, value any)
	Adopt(name Holder, found any)
	Get(name Holder) any
	IgnoreDuplicate(value any)
	HasCoin(coin Holder, mode int) bool
	GetCoin(coin Holder, mode int) any
	GetCoinFrom(coin Holder, modes []int) any
	Errorf(loc *errorsink.Location, msg string, args ...any)
	SetMode(mode int)
	IsMode(mode int) bool
	CurrentMode() int
	Eval(e Expr) any
	EvalAsStringer(e Expr) (fmt.Stringer, bool)
	EvalAsNumber(e Expr) AsNumber
	DumpTo(w io.Writer)
	SetStepName(s string)
	EnableSymbol(to Holder)
	ExportSymbolsTo(iw IndentWriter)
	NewObjId(loc *errorsink.Location) ResolvableHolder
	PendingObjId(loc *errorsink.Location) ResolvableHolder
}

type AsNumber interface {
	F64() float64
}

type InitMe interface {
	InitMe(storage RuntimeStorage) any
}
