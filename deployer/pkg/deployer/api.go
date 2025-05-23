package deployer

import (
	"io"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Deployer interface {
	AddSymbolListener(lsnr pluggable.SymbolListener)
	UseModule(mod string) error
	ReadScriptsFrom(indir string) error
	Deploy(targets ...string) error

	Traverse(lsnr pluggable.RepositoryTraverser)
	ObtainTools() *pluggable.Tools // for the benefit of plugins
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
