package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

// For each action verb, there is exactly one handler.  It is created up front and it does the job of parsing lines and creating individual actions.

type FindCommandHandler struct {
	tools *corebottom.Tools
}

func (ech *FindCommandHandler) Handle(parent driverbottom.AttachResult, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 || len(tokens) > 3 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "Find: <class-identifier> [instance-name]")
		return drivertop.NewIgnoreInnerScope()
	}

	clz, ok := tokens[1].(driverbottom.Identifier)
	if !ok {
		ech.tools.Reporter.Report(tokens[1].Loc().Offset, "Find: <class-identifier> [instance-name]")
		return drivertop.NewIgnoreInnerScope()
	}

	var name driverbottom.String
	if len(tokens) == 3 {
		name, ok = tokens[2].(driverbottom.String)
		if !ok {
			ech.tools.Reporter.Report(tokens[1].Loc().Offset, "Find: <class-identifier> [instance-name]")
			return drivertop.NewIgnoreInnerScope()
		}
	}

	ea := &FindAction{tools: ech.tools, loc: tokens[0].Loc(), what: clz, named: name, props: make(map[driverbottom.Identifier]driverbottom.Expr)}
	parent.Attach(ea)

	return drivertop.NewPropertiesInnerScope(ech.tools.CoreTools, ea)
}

func NewFindCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &FindCommandHandler{tools: tools}
}
