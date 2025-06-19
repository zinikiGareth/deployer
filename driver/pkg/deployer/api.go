package deployer

import (
	"io"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type Driver interface {
	AddSymbolListener(lsnr driverbottom.SymbolListener)
	UseModule(mod string) error
	ReadScriptsFrom(indir string) error
	FindAndReadEnvs(dirs []string, file string) bool
	Traverse(lsnr driverbottom.RepositoryTraverser)

	DoStuff() error
	ObtainCoreTools() *driverbottom.CoreTools

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
