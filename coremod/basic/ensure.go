package basic

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureAction struct {
	loc   pluggable.Location
	what  pluggable.SymbolType
	named string
}

func (ea EnsureAction) Loc() pluggable.Location {
	return ea.loc
}

func (ea EnsureAction) Where() pluggable.Location {
	return ea.loc
}

func (ea EnsureAction) What() pluggable.SymbolType {
	return ea.what
}

func (ea EnsureAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("EnsureCommand")
	w.AttrsWhere(ea)
	w.TextAttr("what", string(ea.what))
	w.TextAttr("named", ea.named)
	w.EndAttrs()
}

func (ea EnsureAction) ShortDescription() string {
	return fmt.Sprintf("Ensure[%s: %s]", ea.what, ea.named)
}

type EnsureCommandHandler struct{}

func (ensure *EnsureCommandHandler) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	ea := &EnsureAction{loc: tokens[0].Loc(), what: pluggable.SymbolType(tokens[1].(pluggable.Identifier).Id()), named: tokens[2].(pluggable.String).Text()}
	parent.Add(ea)
	return interpreters.DisallowInnerScope()
}
