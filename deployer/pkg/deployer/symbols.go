package deployer

type SymbolListener interface {
	Symbol(where Location, what SymbolType, who SymbolName)
}

// These may want to change in the fullness of time
type SymbolType string
type SymbolName string

type Location struct {
	File   string
	Line   int
	Offset int
}

func NewLocation(file string, line, offset int) Location {
	return Location{File: file, Line: line, Offset: offset}
}
