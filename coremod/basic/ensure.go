package basic

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureAction struct {
	loc      pluggable.Location
	what     pluggable.Identifier
	resolved pluggable.Noun
	named    string
	props    map[pluggable.Identifier]any
}

func (ea EnsureAction) Loc() pluggable.Location {
	return ea.loc
}

func (ea EnsureAction) Where() pluggable.Location {
	return ea.loc
}

func (ea EnsureAction) What() pluggable.SymbolType {
	return pluggable.SymbolType(ea.what.Id())
}

func (ea EnsureAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("EnsureCommand")
	w.AttrsWhere(ea)
	w.TextAttr("what", ea.what.Id())
	if ea.resolved == nil {
		w.TextAttr("not-resolved", ea.what.Id())
	} else {
		w.TextAttr("resolved", ea.resolved.ShortDescription())
	}
	w.TextAttr("named", ea.named)
	if len(ea.props) > 0 {
		w.Indent()
		for k, v := range ea.props {
			w.IndPrintf("%s <- %v\n", k, v)
		}
		w.UnIndent()
	}
	w.EndAttrs()
}

func (ea EnsureAction) ShortDescription() string {
	return fmt.Sprintf("Ensure[%s: %s]", ea.what.Id(), ea.named)
}

func (ea *EnsureAction) AddProperty(name pluggable.Identifier, value any) {
	ea.props[name] = value
}

func (ea *EnsureAction) Resolve(r pluggable.Resolver) {
	ea.resolved = r.Resolve(ea.what)
}

type EnsureCommandHandler struct{}

func (ensure *EnsureCommandHandler) Handle(reporter errors.ErrorRepI, repo pluggable.Repository, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	// TODO: allow 2 or 3
	// TODO: errors not panics
	if len(tokens) != 3 {
		panic("tokens are wrong")
	}
	if tokens[1].(pluggable.Identifier).Id() != "test.S3.Bucket" {
		panic("token[1] is wrong")
	}
	if tokens[2].(pluggable.String).Text() != "org.ziniki.launch_bucket" {
		panic("token[2] is wrong")
	}

	ea := &EnsureAction{loc: tokens[0].Loc(), what: tokens[1].(pluggable.Identifier), named: tokens[2].(pluggable.String).Text(), props: make(map[pluggable.Identifier]any)}
	parent.Add(ea)
	return interpreters.PropertiesInnerScope(ea)
}
