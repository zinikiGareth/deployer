package files

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type dirCommandHandler struct {
	tools *external.Tools
}

func (dch *dirCommandHandler) Handle(parent pluggable.AttachResult, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 {
		dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: expr...")
		return interpreters.NewIgnoreInnerScope()
	}
	// if assignTo == nil {
	// 	dch.tools.Reporter.Report(tokens[0].Loc().Offset, "files.dir: must assign to an output variable")
	// 	return interpreters.NewIgnoreInnerScope()
	// }

	exprs, ok := dch.tools.Parser.ParseMultiple(tokens[1:])
	if !ok {
		return interpreters.NewIgnoreInnerScope()
	}

	da := &dirAction{tools: dch.tools, loc: tokens[0].Loc(), exprs: exprs}
	parent.Attach(da)

	return interpreters.NewDisallowInnerScope(dch.tools.CoreTools)
}

func NewDirCommandHandler(tools *external.Tools) pluggable.VerbCommand {
	return &dirCommandHandler{tools: tools}
}
