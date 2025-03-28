package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type noInnerScope struct {
	sink errors.ErrorSink
}

func (b *noInnerScope) BlockedLine(lineNo, lenIndent int, text string) pluggable.ProvideBlockedLine {
	b.sink.Report(lineNo, 0, text, "nested content is not allowed here")
	return b
}

func DisallowInnerScope(sink errors.ErrorSink) pluggable.ProvideBlockedLine {
	return &noInnerScope{sink: sink}
}
