package parser

import (
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func Parse(registry registry.Recall, repo pluggable.Repository, errorSink errors.ErrorSink, fileName, file string) {
	reporter := errors.NewErrorReporter(errorSink)
	globalScope := NewScopedHandlers(registry, repo)
	globalInterpreter := NewInterpreter(repo, globalScope)
	lineLexicator := NewLineLexicator(reporter, globalInterpreter, fileName)
	blocker := NewBlocker(reporter, lineLexicator)
	provideLines(file, blocker)
}
