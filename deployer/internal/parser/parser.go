package parser

import (
	"ziniki.org/deployer/deployer/internal/parser/blocker"
	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/interpreters"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/utils"
)

func Parse(tools *pluggable.Tools, fileName, file string) {
	globalInterpreter := interpreters.NewVerbCommandInterpreter(tools, NewTopLevelAttacher(tools), "top-level", false)
	lineLexicator := lexicator.NewLineLexicator(tools, fileName)
	tools.Parser = exprs.NewExprParser(tools)
	blocker := blocker.NewBlocker(tools, lineLexicator, globalInterpreter)
	provideLines(tools.Reporter, file, blocker)
}

func provideLines(reporter errorsink.ErrorRepI, fromFile string, to pluggable.ProvideLine) {
	lines, err := utils.FileAsLines(fromFile)
	if err != nil {
		reporter.Reportf(0, "could not open file %s: %v", fromFile, err)
	}
	to.BeginFile(fromFile)
	for n, l := range lines {
		// turn 0-(n-1) into 1-n by adding 1 to the index
		to.HaveLine(n+1, l)
	}
	to.EndFile()
}
