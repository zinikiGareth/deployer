package parser

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}

type Interpreter interface {
	HaveTokens(tokens []Token)
}

type Token interface {
}

type Scoper interface {
}
