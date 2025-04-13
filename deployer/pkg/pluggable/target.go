package pluggable

type RuntimeStorage interface {
	Bind(name SymbolName, value any)
}

type Target interface {
	Execute(runtime RuntimeStorage)
}
