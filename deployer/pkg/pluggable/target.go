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
}

type InitMe interface {
	InitMe(storage RuntimeStorage) any
}

// TODO: I feel this should "logically" be in coremod, because that's where the ensure logic is, but I'm not sure I know
// how to work that module magic yet.  So move it later.
type Ensurable interface {
	// TODO: do we want to just remove all of this?
	Ensure()
}
