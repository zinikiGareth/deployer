package files

type FileSource interface {
	// TODO: something that says "here's what I got" (TDA, of course)
	// All() FileSource
	// Match(pattern string) FileSource
	// One(name string) FileSource
	PourOut(name string, into FileDest)
}

type FileDest interface {
	PourInto(name string /* TODO: needs a second arg with the body in it */)
}

type ThingyHolder interface {
	ObtainDest() FileDest
}
