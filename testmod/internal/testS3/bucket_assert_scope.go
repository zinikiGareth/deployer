package testS3

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type bucketContentsScope struct {
	tools *corebottom.Tools
	aba   *assertBucketAction
}

func (b *bucketContentsScope) HaveTokens(tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) != 1 {
		b.tools.Reporter.Reportf(0, "may only have one file per line")
		return drivertop.NewIgnoreInnerScope()
	}
	str, ok := tokens[0].(driverbottom.String)
	if !ok {
		b.tools.Reporter.Reportf(0, "file name must be a string literal")
		return drivertop.NewIgnoreInnerScope()
	}
	b.aba.files = append(b.aba.files, str.Text())
	return drivertop.NewDisallowInnerScope(b.tools.CoreTools)
}

func (b *bucketContentsScope) Completed() {
}

func BucketContentsScope(tools *corebottom.Tools, aba *assertBucketAction) driverbottom.Interpreter {
	return &bucketContentsScope{tools: tools, aba: aba}
}
