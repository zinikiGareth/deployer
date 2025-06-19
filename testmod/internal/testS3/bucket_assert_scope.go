package testS3

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type bucketContentsScope struct {
	tools *external.Tools
	aba   *assertBucketAction
}

func (b *bucketContentsScope) HaveTokens(tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 1 {
		b.tools.Reporter.Reportf(0, "may only have one file per line")
		return interpreters.NewIgnoreInnerScope()
	}
	str, ok := tokens[0].(driverbottom.String)
	if !ok {
		b.tools.Reporter.Reportf(0, "file name must be a string literal")
		return interpreters.NewIgnoreInnerScope()
	}
	b.aba.files = append(b.aba.files, str.Text())
	return interpreters.NewDisallowInnerScope(b.tools.CoreTools)
}

func (b *bucketContentsScope) Completed() {
}

func BucketContentsScope(tools *external.Tools, aba *assertBucketAction) driverbottom.Interpreter {
	return &bucketContentsScope{tools: tools, aba: aba}
}
