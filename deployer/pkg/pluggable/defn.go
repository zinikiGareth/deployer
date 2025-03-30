package pluggable

type Definition interface {
	Where() Location
	What() SymbolType
}
