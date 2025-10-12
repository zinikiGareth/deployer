package shells

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type BashCommandHandler struct {
	tools *corebottom.Tools
}

func (bch *BashCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	es, ok := bch.tools.Parser.ParseMultiple(scope, tokens[1:])
	if !ok {
		return drivertop.NewIgnoreInnerScope()
	}
	if len(es) != 1 {
		bch.tools.Reporter.Report(tokens[0].Loc().Offset, "bash: <script>")
		return drivertop.NewIgnoreInnerScope()
	}

	log.Printf("have bash script %s\n", es[0])
	return drivertop.NewDisallowInnerScope(bch.tools.CoreTools)
}

func NewBashCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &BashCommandHandler{tools: tools}
}
