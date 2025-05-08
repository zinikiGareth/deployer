package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

const (
	PREPARE_MODE int = iota
	EXECUTE_MODE
)

type RuntimeStorage interface {
	Bind(name SymbolName, value any)
	Get(name SymbolName) any
	Errorf(loc *errors.Location, msg string, args ...any)
	SetMode(mode int)
	IsMode(mode int) bool
	Eval(e Expr) any
	EvalAsString(e Expr) string
}

type InitMe interface {
	InitMe(storage RuntimeStorage) any
}
