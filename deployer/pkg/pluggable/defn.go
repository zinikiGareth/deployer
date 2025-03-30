package pluggable

import "io"

type Definition interface {
	Where() Location
	What() SymbolType

	ShortDescription() string
	DumpTo(to io.Writer)
}
