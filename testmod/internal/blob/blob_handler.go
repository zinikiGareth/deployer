package blob

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type blobCommandHandler struct {
	tools *corebottom.Tools
}

func (bch *blobCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 1 {
		bch.tools.Reporter.Report(tokens[0].Loc().Offset, "blob <name>")
		return drivertop.NewIgnoreInnerScope()
	}

	expr, ok := bch.tools.Parser.Parse(tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}

	parent.Attach(&createBlobAction{Locatable: tokens[0], expr: expr})

	return drivertop.NewDisallowInnerScope(bch.tools.CoreTools)
}

func NewBlobCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &blobCommandHandler{tools: tools}
}
