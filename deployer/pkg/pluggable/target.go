package pluggable

import "reflect"

const (
	DRYRUN_MODE int = iota
	EXECUTE_MODE
)

type RuntimeStorage interface {
	Bind(name SymbolName, value any)
	Errorf(loc Location, msg string, args ...any)
	SetMode(mode int)
	IsMode(mode int) bool
	ObtainDriver(forType reflect.Type) any
}

type InitMe interface {
	InitMe(storage RuntimeStorage)
}

type Executable interface {
	Prepare(runtime RuntimeStorage) any
	Execute(runtime RuntimeStorage)
}

// TODO: I feel this should "logically" be in coremod, because that's where the ensure logic is, but I'm not sure I know
// how to work that module magic yet.  So move it later.
type Ensurable interface {
	Ensure(runtime RuntimeStorage)
}
