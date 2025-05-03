package basic

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnsureAction struct {
	loc      *errors.Location
	what     pluggable.Identifier
	resolved pluggable.Noun
	named    pluggable.String
	props    map[pluggable.Identifier]pluggable.Expr
}

func (ea *EnsureAction) Loc() *errors.Location {
	return ea.loc
}

func (ea *EnsureAction) Where() *errors.Location {
	return ea.loc
}

func (ea *EnsureAction) What() pluggable.SymbolType {
	return pluggable.SymbolType(ea.what.Id())
}

func (ea *EnsureAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("EnsureCommand")
	w.AttrsWhere(ea)
	w.TextAttr("what", ea.what.Id())
	if ea.resolved == nil {
		w.TextAttr("not-resolved", ea.what.Id())
	} else {
		w.TextAttr("resolved", ea.resolved.ShortDescription())
	}
	if ea.named != nil {
		w.TextAttr("named", ea.named.Text())
	}
	if len(ea.props) > 0 {
		w.Indent()
		for k, v := range ea.props {
			w.IndPrintf("%s <- %s\n", k, v.String())
		}
		w.UnIndent()
	}
	w.EndAttrs()
}

func (ea *EnsureAction) ShortDescription() string {
	return fmt.Sprintf("Ensure[%s: %s]", ea.what.Id(), ea.named)
}

func (ea *EnsureAction) AddProperty(tools *pluggable.Tools, name pluggable.Identifier, value pluggable.Expr) {
	if name.Id() == "name" {
		if ea.named != nil {
			tools.Reporter.Report(name.Loc().Offset, "duplicate definition of name")
			return
		}
		str, ok := value.(pluggable.String)
		if !ok {
			tools.Reporter.Report(value.Loc().Offset, "name must be a string")
			return
		}
		ea.named = str
	} else {
		if ea.props[name] != nil {
			tools.Reporter.Reportf(name.Loc().Offset, "duplicate definition of %s", name.Id())
			return
		}
		ea.props[name] = value
	}
}

func (ea *EnsureAction) Completed(tools *pluggable.Tools) {
	if ea.named == nil {
		tools.Reporter.At(ea.loc.Line)
		tools.Reporter.Report(ea.loc.Offset, "ensure requires a name to be defined")
	}
}

func (ea *EnsureAction) Resolve(r pluggable.Resolver) {
	ea.resolved = r.Resolve(ea.what)
}

// TODO: we need to consider multiple phases here
// We need to store the thing we create first time for second time
// We need to have some way of associating this with the var
// Using a map from "action" (or other) to runtime value seems a reasonable way to go
func (ea *EnsureAction) Prepare(runtime pluggable.RuntimeStorage) (pluggable.ExecuteAction, any) {
	// So the logic for ensure is that we create an object "locally" that represents the thing we want to ensure
	// Then we call the "ensure" method on that
	// It is an error for the object created not to implement the Ensurable contract

	obj := ea.resolved.CreateWithName(ea.named.Text())
	ens, ok := obj.(pluggable.Ensurable)
	if !ok {
		runtime.Errorf(ea.loc, "the type "+ea.what.Id()+" is not ensurable")
		return nil, nil
	}
	exe := ens.Ensure(runtime)
	return exe, obj
}

type EnsureCommandHandler struct{}

func (ensure *EnsureCommandHandler) Handle(tools *pluggable.Tools, parent pluggable.ContainingContext, tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) < 2 || len(tokens) > 3 {
		tools.Reporter.Report(tokens[0].Loc().Offset, "ensure: <class-identifier> [class-name]")
		return interpreters.IgnoreInnerScope()
	}

	clz, ok := tokens[1].(pluggable.Identifier)
	if !ok {
		tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [class-name]")
		return interpreters.IgnoreInnerScope()
	}

	var name pluggable.String
	if len(tokens) == 3 {
		name, ok = tokens[2].(pluggable.String)
		if !ok {
			tools.Reporter.Report(tokens[1].Loc().Offset, "ensure: <class-identifier> [class-name]")
			return interpreters.IgnoreInnerScope()
		}
	}

	ea := &EnsureAction{loc: tokens[0].Loc(), what: clz, named: name, props: make(map[pluggable.Identifier]pluggable.Expr)}
	parent.Add(ea)
	return interpreters.PropertiesInnerScope(ea)
}
