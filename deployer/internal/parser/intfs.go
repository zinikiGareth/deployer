package parser

import "ziniki.org/deployer/deployer/pkg/pluggable"

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}

type ProvideBlockedLine interface {
	BlockedLine(lineNo, indent int, text string)
}

type Interpreter interface {
	HaveTokens(tokens []pluggable.Token)
}

type Scoper interface {
	FindVerb(v pluggable.Identifier) pluggable.Action
}
