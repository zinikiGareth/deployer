package target

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type commandScope struct {
	reporter *errors.ErrorReporter
}

func (b *commandScope) BlockedLine(lineNo, lenIndent int, text string) pluggable.ProvideBlockedLine {
	log.Printf("needs work")
	return interpreters.DisallowInnerScope(b.reporter.Sink())
}

func TargetCommandInterpreter(reporter *errors.ErrorReporter) pluggable.ProvideBlockedLine {
	return &commandScope{reporter: reporter}
}
