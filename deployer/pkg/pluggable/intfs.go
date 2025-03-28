package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}

type Interpreter interface {
	HaveTokens(reporter *errors.ErrorReporter, tokens []Token) Interpreter
}

type Scoper interface {
	FindVerb(v Identifier) Action
}
