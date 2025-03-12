package parser

import (
	"ziniki.org/deployer/deployer/internal/repo"
)

// deffo need an error handler as well
func Parse(repo repo.Repository, fileName, file string) {
	globalScope := &ScopedHandlers{}
	globalInterpreter := NewInterpreter(repo, globalScope)
	lineLexicator := NewLineLexicator(globalInterpreter, fileName)
	blocker := NewBlocker(lineLexicator)
	provideLines(file, blocker)
}
