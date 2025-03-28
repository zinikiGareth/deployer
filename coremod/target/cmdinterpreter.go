package target

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type commandScope struct {
	reporter errors.ErrorRepI
}

func (b *commandScope) HaveTokens(reporter errors.ErrorRepI, tokens []pluggable.Token) pluggable.Interpreter {
	// I am hacking this in first, and then I need to come back and do more on it
	log.Printf("hello %d %v?\n", len(tokens), tokens[1])

	if len(tokens) != 3 {
		panic("tokens are wrong")
	}
	if tokens[0].(pluggable.Identifier).Id() != "ensure" {
		panic("token[0] is wrong")
	}
	if tokens[1].(pluggable.Identifier).Id() != "test.S3.Bucket" {
		panic("token[1] is wrong")
	}
	if tokens[2].(pluggable.String).Text() != "org.ziniki.launch_bucket" {
		panic("token[2] is wrong")
	}

	log.Printf("needs work")
	return interpreters.DisallowInnerScope()
}

func TargetCommandInterpreter(reporter errors.ErrorRepI) pluggable.Interpreter {
	return &commandScope{reporter: reporter}
}
