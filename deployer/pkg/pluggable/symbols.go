package pluggable

import "fmt"

type SymbolListener interface {
	ReadingFile(file string)
	Symbol(who SymbolName, is Definition)
}

// These may want to change in the fullness of time
type SymbolType string
type SymbolName string

type Location struct {
	File   string
	Line   int
	Offset int
}

func (loc Location) InFile() string {
	return fmt.Sprintf("%d.%d", loc.Line, loc.Offset)
}

func (loc Location) String() string {
	return fmt.Sprintf("%s:%d.%d", loc.File, loc.Line, loc.Offset)
}

func NewLocation(file string, line, offset int) Location {
	return Location{File: file, Line: line, Offset: offset}
}
