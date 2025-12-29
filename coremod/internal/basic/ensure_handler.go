package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

// For each action verb, there is exactly one handler.  It is created up front and it does the job of parsing lines and creating individual actions.

type EnsureCommandHandler struct {
	tools *corebottom.Tools
}

func (ech *EnsureCommandHandler) Handle(parent driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	if len(tokens) < 2 {
		ech.tools.Reporter.Report(tokens[0].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return drivertop.NewIgnoreInnerScope()
	}

	clz, ok := tokens[1].(driverbottom.Identifier)
	if !ok {
		ech.tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [instance-name]")
		return drivertop.NewIgnoreInnerScope()
	}

	var name driverbottom.Expr
	if len(tokens) > 2 {
		name, ok = ech.tools.Parser.Parse(scope, tokens[2:])
		if !ok {
			ech.tools.Reporter.Report(tokens[2].Loc().Offset, "ensure: <class-identifier> [instance-name]")
			return drivertop.NewIgnoreInnerScope()
		}
	}

	ea := &EnsureAction{tools: ech.tools, scope: scope, loc: tokens[0].Loc(), what: clz, named: name, props: make(map[driverbottom.Identifier]driverbottom.Expr)}
	parent.Attach(ea)

	return drivertop.NewPropertiesInnerScope(ech.tools.CoreTools, ea)
}

func NewEnsureCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &EnsureCommandHandler{tools: tools}
}
