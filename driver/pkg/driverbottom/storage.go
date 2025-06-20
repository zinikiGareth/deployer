package driverbottom

import (
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
	EvalAsString(e Expr) string
	DumpTo(w io.Writer)
}

type InitMe interface {
	InitMe(storage RuntimeStorage) any
}
