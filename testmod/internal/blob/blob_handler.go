package blob

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

type blobCommandHandler struct {
	tools *external.Tools
}

func (bch *blobCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 1 {
		bch.tools.Reporter.Report(tokens[0].Loc().Offset, "blob <name>")
		return interpreters.NewIgnoreInnerScope()
	}

	expr, ok := bch.tools.Parser.Parse(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	parent.Attach(&createBlobAction{Locatable: tokens[0], expr: expr})

	return interpreters.NewDisallowInnerScope(bch.tools.CoreTools)
}

func NewBlobCommandHandler(tools *external.Tools) driverbottom.VerbCommand {
	return &blobCommandHandler{tools: tools}
}
