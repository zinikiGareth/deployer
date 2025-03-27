package parser

import "ziniki.org/deployer/deployer/pkg/errors"

type noInnerScope struct {
	sink errors.ErrorSink
}

func (b *noInnerScope) BlockedLine(lineNo, lenIndent int, text string) ProvideBlockedLine {
	b.sink.Report(lineNo, 0, text, "nested content is not allowed here")
	return b
}

func DisallowInnerScope(sink errors.ErrorSink) ProvideBlockedLine {
	return &noInnerScope{sink: sink}
}
