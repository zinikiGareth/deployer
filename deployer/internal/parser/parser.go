package parser

import (
	"ziniki.org/deployer/deployer/internal/repo"
)

// deffo need an error handler as well
func Parse(repo repo.Repository, file string) {
	provideLines(file, NewBlocker())
}
