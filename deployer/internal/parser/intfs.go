package parser

import "ziniki.org/deployer/deployer/internal/repo"

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}

type ProvideBlockedLine interface {
	BlockedLine(lineNo, indent int, text string)
}

type Interpreter interface {
	HaveTokens(tokens []Token)
}

type Token interface {
}

type Action interface {
	Handle(repo repo.Repository, tokens []Token)
}

type Scoper interface {
	FindVerb(v *IdentifierToken) Action
}
