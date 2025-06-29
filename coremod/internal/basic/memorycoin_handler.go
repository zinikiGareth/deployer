package basic

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

// For each action verb, there is exactly one handler.  It is created up front and it does the job of parsing lines and creating individual actions.

type MemoryCoinCommandHandler struct {
	tools *corebottom.Tools
}

func (mcch *MemoryCoinCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 || len(tokens) > 3 {
		mcch.tools.Reporter.Report(tokens[0].Loc().Offset, "coin: <class-identifier> [instance-name]")
		return drivertop.NewIgnoreInnerScope()
	}

	clz, ok := tokens[1].(driverbottom.Identifier)
	if !ok {
		mcch.tools.Reporter.Report(tokens[1].Loc().Offset, "coin: <class-identifier> [instance-name]")
		return drivertop.NewIgnoreInnerScope()
	}

	var name driverbottom.String
	if len(tokens) == 3 {
		name, ok = tokens[2].(driverbottom.String)
		if !ok {
			mcch.tools.Reporter.Report(tokens[1].Loc().Offset, "coin: <class-identifier> [instance-name]")
			return drivertop.NewIgnoreInnerScope()
		}
	}

	mca := &MemoryCoinAction{tools: mcch.tools, loc: tokens[0].Loc(), what: clz, named: name, props: make(map[driverbottom.Identifier]driverbottom.Expr)}
	log.Printf("mcch handle attaching %p to %p", mca, parent)
	parent.Attach(mca)

	return drivertop.NewPropertiesInnerScope(mcch.tools.CoreTools, mca)
}

func NewMemoryCoinCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &MemoryCoinCommandHandler{tools: tools}
}
