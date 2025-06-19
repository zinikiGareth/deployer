package driverbottom

import (
	"io"
)

type Driver interface {
	AddSymbolListener(lsnr SymbolListener)
	UseModule(mod string) error
	ReadScriptsFrom(indir string) error
	FindAndReadEnvs(dirs []string, file string) bool
	Traverse(lsnr RepositoryTraverser)

	DoStuff() error
	ObtainCoreTools() *CoreTools

	UserError(msg string)
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
