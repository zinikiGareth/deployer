package deployer

import (
	"io"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Driver interface {
	AddSymbolListener(lsnr pluggable.SymbolListener)
	UseModule(mod string) error
	ReadScriptsFrom(indir string) error
	FindAndReadEnvs(dirs []string, file string) bool
	Traverse(lsnr pluggable.RepositoryTraverser)
}

type TestRunner interface {
	ErrorHandlerFor(purpose string) ErrorHandler
}

type ErrorHandler interface {
	io.Writer
	WriteMsg(msg string)
	Writef(fmt string, args ...any)
	Fail()
	Close()
}
