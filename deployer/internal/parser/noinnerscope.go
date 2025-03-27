package parser

type NoInnerScope struct {
}

func (b *NoInnerScope) BlockedLine(lineNo, lenIndent int, text string) ProvideBlockedLine {
	// TODO: should be an error
	panic("nested content is not allowed here")
}
