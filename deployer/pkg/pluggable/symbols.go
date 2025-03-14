package pluggable

import "fmt"

type SymbolListener interface {
	ReadingFile(file string)
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

func (loc *Location) String() string {
	return fmt.Sprintf("%d.%d", loc.Line, loc.Offset)
}

func NewLocation(file string, line, offset int) Location {
	return Location{File: file, Line: line, Offset: offset}
}
