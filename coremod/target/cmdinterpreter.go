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

func (b *commandScope) HaveTokens(reporter *errors.ErrorReporter, tokens []pluggable.Token) pluggable.Interpreter {
	log.Printf("needs work")
	return interpreters.DisallowInnerScope()
}

func TargetCommandInterpreter(reporter *errors.ErrorReporter) pluggable.Interpreter {
	return &commandScope{reporter: reporter}
}
