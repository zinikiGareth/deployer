package deployer

import "io"

type Deployer interface {
	AddSymbolListener(lsnr SymbolListener)
	ReadScriptsFrom(indir string) error
	Deploy() error
}

type TestRunner interface {
	ErrorHandlerFor(purpose string) ErrorHandler
}

// TODO: the deployer cmd will want a version of this that writes to stdout
type ErrorHandler interface {
	io.Writer
	WriteMsg(msg string)
	Writef(fmt string, args ...any)
	Fail()
	Close()
}
