package parser

import (
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// deffo need an error handler as well
func Parse(registry registry.Recall, repo pluggable.Repository, fileName, file string) {
	globalScope := NewScopedHandlers(registry, repo)
	globalInterpreter := NewInterpreter(repo, globalScope)
	lineLexicator := NewLineLexicator(globalInterpreter, fileName)
	blocker := NewBlocker(lineLexicator)
	provideLines(file, blocker)
}
