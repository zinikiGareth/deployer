package pluggable

type SymbolListener interface {
	ReadingFile(file string)
	Symbol(who SymbolName, is Action)
}

// These may want to change in the fullness of time
type SymbolType string
type SymbolName string
