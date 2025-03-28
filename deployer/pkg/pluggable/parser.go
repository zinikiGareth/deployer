package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type Token interface {
	Loc() Location
}

type Identifier interface {
	Token
	Id() string
}

type Action interface {
	Handle(reporter *errors.ErrorReporter, repo Repository, tokens []Token) Interpreter
}
