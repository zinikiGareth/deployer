package parser

import (
	"ziniki.org/deployer/deployer/internal/parser/blocker"
	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/interpreters"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/internal/parser/scopes"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/utils"
)

func Parse(tools *pluggable.Tools, fileName, file string) {
	globalScope := scopes.NewScopedHandlers(tools)
	globalInterpreter := interpreters.NewInterpreter(tools, globalScope)
	lineLexicator := lexicator.NewLineLexicator(tools, fileName)
	tools.Parser = exprs.NewExprParser(tools)
	blocker := blocker.NewBlocker(tools, lineLexicator, globalInterpreter)
	provideLines(file, blocker)
}

func provideLines(fromFile string, to pluggable.ProvideLine) {
	lines, err := utils.FileAsLines(fromFile)
	if err != nil {
		panic("need an error handler")
	}
	to.BeginFile(fromFile)
	for n, l := range lines {
		// turn 0-(n-1) into 1-n by adding 1 to the index
		to.HaveLine(n+1, l)
	}
	to.EndFile()
}
