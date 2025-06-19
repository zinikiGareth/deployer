package testS3

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type assertBucketHandler struct {
	tools *corebottom.Tools
}

func (abh *assertBucketHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 1 {
		abh.tools.Reporter.Report(tokens[0].Loc().Offset, "test.assertBucketHas: <bucket>")
		return drivertop.NewIgnoreInnerScope()
	}
	// if assignTo != nil {
	// 	abh.tools.Reporter.Report(tokens[0].Loc().Offset, "test.assertBucketHas: cannot assign an output variable")
	// 	return drivertop.NewIgnoreInnerScope()
	// }

	expr, ok := abh.tools.Parser.Parse(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	ca := &assertBucketAction{tools: abh.tools, loc: tokens[0].Loc(), bucket: expr}
	parent.Attach(ca)

	return BucketContentsScope(abh.tools, ca)
}

func NewAssertBucketHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &assertBucketHandler{tools: tools}
}
