package parser

import (
	"fmt"

	"ziniki.org/deployer/deployer/internal/repo"
)

// deffo need an error handler as well
func Parse(repo repo.Repository, file string) {
	provideLines(file, NewBlocker(&DisplayIt{}))
}

type DisplayIt struct {
}

func (d *DisplayIt) HaveLine(n int, s string) {
	fmt.Printf("%d: %s\n", n, s)
}
