package testS3

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type assertBucketHandler struct {
	tools *external.Tools
}

func (abh *assertBucketHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 1 {
		abh.tools.Reporter.Report(tokens[0].Loc().Offset, "test.assertBucketHas: <bucket>")
		return interpreters.NewIgnoreInnerScope()
	}
	// if assignTo != nil {
	// 	abh.tools.Reporter.Report(tokens[0].Loc().Offset, "test.assertBucketHas: cannot assign an output variable")
	// 	return interpreters.NewIgnoreInnerScope()
	// }

	expr, ok := abh.tools.Parser.Parse(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	ca := &assertBucketAction{tools: abh.tools, loc: tokens[0].Loc(), bucket: expr}
	parent.Attach(ca)

	return BucketContentsScope(abh.tools, ca)
}

func NewAssertBucketHandler(tools *external.Tools) pluggable.VerbCommand {
	return &assertBucketHandler{tools: tools}
}
