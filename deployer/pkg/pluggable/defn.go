package pluggable

type Definition interface {
	Where() Location
	What() SymbolType

	ShortDescription() string
	DumpTo(to IndentWriter)
}
