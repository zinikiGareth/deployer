package pluggable

type RuntimeStorage interface {
	Bind(name SymbolName, value any)
	Errorf(loc Location, msg string, args ...any)
}

// There is a part of me that thinks this should be in coremod too,
// but given it is so central to what we do - command line arguments are targets - I don't think that's reasonable.
type Target interface {
	Execute(runtime RuntimeStorage)
}

// TODO: I feel this should "logically" be in coremod, because that's where the ensure logic is, but I'm not sure I know
// how to work that module magic yet.  So move it later.
type Ensurable interface {
	Ensure(runtime RuntimeStorage)
}
