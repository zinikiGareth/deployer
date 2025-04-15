package pluggable

const (
	PREPARE_MODE int = iota
	EXECUTE_MODE
)

type TargetThing interface {
	Executable
	ExecuteAction
}

type RuntimeStorage interface {
	Bind(name SymbolName, value any)
	Errorf(loc Location, msg string, args ...any)
	SetMode(mode int)
	IsMode(mode int) bool
	ObtainDriver(forType string) any
	BindAction(a Executable, av ExecuteAction)
	RetrieveAction(a Executable) ExecuteAction
}

type InitMe interface {
	InitMe(storage RuntimeStorage) any
}

type Executable interface {
	Prepare(runtime RuntimeStorage) (ExecuteAction, any)
}

type ExecuteAction interface {
	Execute(runtime RuntimeStorage)
}

// TODO: I feel this should "logically" be in coremod, because that's where the ensure logic is, but I'm not sure I know
// how to work that module magic yet.  So move it later.
type Ensurable interface {
	Ensure(runtime RuntimeStorage) ExecuteAction
}
