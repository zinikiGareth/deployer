package deployer

import (
	"io"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Deployer interface {
	AddSymbolListener(lsnr pluggable.SymbolListener)
	ReadScriptsFrom(indir string) error
	Deploy(targets ...string) error

	Traverse(lsnr pluggable.RepositoryTraverser)
	ObtainRegister() pluggable.Register // for the benefit of plugins
	ObtainStorage() pluggable.RuntimeStorage
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
