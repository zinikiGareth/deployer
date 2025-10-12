package corebottom

import "io"

type FileSource interface {
	// TODO: something that says "here's what I got" (TDA, of course)
	// All() FileSource
	// Match(pattern string) FileSource
	// One(name string) FileSource
	PourAll(into FileDest) error
	PourOut(name string, into FileDest) error
}

type FileDest interface {
	PourInto(name string, contents io.Reader) error
	Relative(s string) (FileDest, error)
}

type DestHolder interface {
	ObtainDest() FileDest
}
