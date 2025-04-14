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
	Locatable
}

type Identifier interface {
	Token
	Id() string
}

type Number interface {
	Token
	Value() float64
}

type String interface {
	Token
	Text() string
}

type Operator interface {
	Token
	Is(op string) bool
	Op() string
}

type Action interface {
	Handle(reporter errors.ErrorRepI, repo Repository, parent ContainingContext, tokens []Token) Interpreter
}

type Noun interface {
	ShortDescription() string
	CreateWithName(named string) any
}

