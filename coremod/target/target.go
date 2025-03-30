package target

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type coreTarget struct {
	loc     pluggable.Location
	name    pluggable.SymbolName
	actions *commandList
}

func (t *coreTarget) Loc() pluggable.Location {
	return t.loc
}

func (t *coreTarget) Where() pluggable.Location {
	return t.loc
}

func (t *coreTarget) What() pluggable.SymbolType {
	return pluggable.SymbolType("core.Target")
}

func (t *coreTarget) ShortDescription() string {
	return "Target[" + string(t.name) + "]"
}

func (t *coreTarget) DumpTo(w pluggable.IndentWriter) {
	w.Intro("target %s", t.name)
	w.AttrsWhere(t)
	// w.Printf("target %s {\n", t.name)
	// w.Printf("    _where_: %s\n", t.loc.String())
	w.ListAttr("actions")
	for _, a := range t.actions.commands {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

type CoreTargetVerb struct {
}

func (t *CoreTargetVerb) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	t1 := tokens[1].(pluggable.Identifier)
	name := pluggable.SymbolName(t1.Id())
	actions := &commandList{}
	target := &coreTarget{loc: t1.Loc(), name: name, actions: actions}
	repo.IntroduceSymbol(name, target)
	return TargetCommandInterpreter(repo, actions)
}
