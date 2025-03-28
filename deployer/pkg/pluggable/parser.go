package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type Token interface {
	Loc() Location
}

type Identifier interface {
	Token
	Id() string
}

type String interface {
	Token
	Text() string
}

type Action interface {
	Handle(reporter errors.ErrorRepI, repo Repository, tokens []Token) Interpreter
}
