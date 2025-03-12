package parser

import (
	"ziniki.org/deployer/deployer/pkg/utils"
)

// deffo need an error handler as well
func provideLines(fromFile string, to ProvideLine) {
	lines, err := utils.FileAsLines(fromFile)
	if err != nil {
		panic("need an error handler")
	}
	for n, l := range lines {
		// turn 0-(n-1) into 1-n by adding 1 to the index
		to.HaveLine(n+1, l)
	}
}
