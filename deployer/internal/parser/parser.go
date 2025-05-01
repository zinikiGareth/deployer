package parser

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func Parse(registry pluggable.Recall, repo pluggable.Repository, errorSink errors.ErrorSink, fileName, file string) {
	reporter := errors.NewErrorReporter(errorSink)
	globalScope := NewScopedHandlers(registry, repo)
	globalInterpreter := NewInterpreter(repo, globalScope)
	lineLexicator := NewLineLexicator(reporter, fileName)
	tools := pluggable.NewTools(reporter, repo)
	tools.Parser = NewExprParser(tools)
	blocker := NewBlocker(tools, lineLexicator, globalInterpreter)
	provideLines(file, blocker)
}
