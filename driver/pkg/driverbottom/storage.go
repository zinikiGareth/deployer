package driverbottom

import (
	"fmt"
	"io"

	"ziniki.org/deployer/driver/pkg/errorsink"
)

type RuntimeStorage interface {
	Bind(name Describable, value any)
	Get(name Var) any
	Read(name SymbolName) any
	Errorf(loc *errorsink.Location, msg string, args ...any)
	SetMode(mode int)
	IsMode(mode int) bool
	Eval(e Expr) any
	EvalAsStringer(e Expr) (fmt.Stringer, bool)
	EvalAsNumber(e Expr) AsNumber
	DumpTo(w io.Writer)
}

type AsNumber interface {
	F64() float64
}
type InitMe interface {
	InitMe(storage RuntimeStorage) any
}
