package parser

import (
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func Parse(registry registry.Recall, repo pluggable.Repository, errorSink errors.ErrorSink, fileName, file string) {
	globalScope := NewScopedHandlers(registry, repo)
	globalInterpreter := NewInterpreter(repo, errorSink, globalScope)
	lineLexicator := NewLineLexicator(globalInterpreter, fileName)
	blocker := NewBlocker(errorSink, lineLexicator)
	provideLines(file, blocker)
}
