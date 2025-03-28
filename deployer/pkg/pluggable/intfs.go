package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}

type ProvideBlockedLine interface {
	BlockedLine(lineNo, indent int, text string) ProvideBlockedLine
}

type Interpreter interface {
	HaveTokens(reporter *errors.ErrorReporter, tokens []Token) ProvideBlockedLine
}

type Scoper interface {
	FindVerb(v Identifier) Action
}
