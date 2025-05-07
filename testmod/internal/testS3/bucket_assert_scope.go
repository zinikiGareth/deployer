package testS3

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type bucketContentsScope struct {
	tools *pluggable.Tools
	aba   *assertBucketAction
}

func (b *bucketContentsScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) != 1 {
		b.tools.Reporter.Reportf(0, "may only have one file per line")
		return interpreters.IgnoreInnerScope()
	}
	str, ok := tokens[0].(pluggable.String)
	if !ok {
		b.tools.Reporter.Reportf(0, "file name must be a string literal")
		return interpreters.IgnoreInnerScope()
	}
	b.aba.files = append(b.aba.files, str.Text())
	return interpreters.DisallowInnerScope(b.tools)
}

func (b *bucketContentsScope) Completed() {
}

func BucketContentsScope(tools *pluggable.Tools, aba *assertBucketAction) pluggable.Interpreter {
	return &bucketContentsScope{tools: tools, aba: aba}
}
