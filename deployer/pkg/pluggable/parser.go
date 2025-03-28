package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}

type Interpreter interface {
	HaveTokens(reporter errors.ErrorRepI, tokens []Token) Interpreter
}

type Scoper interface {
	FindVerb(v Identifier) Action
}

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
