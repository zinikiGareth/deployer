package target

import (
	"fmt"
	"io"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type coreTarget struct {
	loc  pluggable.Location
	name pluggable.SymbolName
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

func (t *coreTarget) DumpTo(w io.Writer) {
	fmt.Fprintf(w, "target %s {\n", t.name)
	fmt.Fprintf(w, "    _where_: %s\n", t.loc.String())
	fmt.Fprintf(w, "}\n")
}

type CoreTargetVerb struct {
}

func (t *CoreTargetVerb) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	t1 := tokens[1].(pluggable.Identifier)
	name := pluggable.SymbolName(t1.Id())
	target := &coreTarget{loc: t1.Loc(), name: name}
	repo.IntroduceSymbol(name, target)
	return TargetCommandInterpreter(repo)
}
