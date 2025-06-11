package blob

import (
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type blobCommandHandler struct {
	tools *pluggable.Tools
}

func (bch *blobCommandHandler) Handle(parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 1 {
		bch.tools.Reporter.Report(tokens[0].Loc().Offset, "blob <name>")
		return interpreters.IgnoreInnerScope()
	}

	expr, ok := bch.tools.Parser.Parse(tokens[1:])
	if !ok {
		return interpreters.IgnoreInnerScope()
	}

	parent.Add(&createBlobAction{Locatable: tokens[0], expr: expr})

	return interpreters.DisallowInnerScope(bch.tools)
}

func NewBlobCommandHandler(tools *pluggable.Tools) pluggable.TargetCommand {
	return &blobCommandHandler{tools: tools}
}
