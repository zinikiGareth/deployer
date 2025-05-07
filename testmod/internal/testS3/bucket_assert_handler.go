package testS3

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type assertBucketHandler struct {
	tools *pluggable.Tools
}

func (abh *assertBucketHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 1 {
		abh.tools.Reporter.Report(tokens[0].Loc().Offset, "test.assertBucketHas: <bucket>")
		return interpreters.IgnoreInnerScope()
	}
	// if assignTo != nil {
	// 	abh.tools.Reporter.Report(tokens[0].Loc().Offset, "test.assertBucketHas: cannot assign an output variable")
	// 	return interpreters.IgnoreInnerScope()
	// }

	expr, ok := abh.tools.Parser.Parse(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	ca := &assertBucketAction{tools: abh.tools, loc: tokens[0].Loc(), bucket: expr}
	parent.Add(ca)

	return BucketContentsScope(abh.tools, ca)
}

func NewAssertBucketHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &assertBucketHandler{tools: tools}
}
