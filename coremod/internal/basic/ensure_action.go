package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/ensurable"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type EnsureAction struct {
	tools    *pluggable.Tools
	loc      *errorsink.Location
	what     pluggable.Identifier
	resolved pluggable.Blank
	named    pluggable.String
	teardown pluggable.TearDown
	// TODO: this really isn't a map here, because what we want is to index by string name
	// Now, we possibly want the "identifier" to that we have the location of the symbol, but it can't be the key
	// because the same "symbol" at two different locations will be different, and we can't index by it.
	props map[pluggable.Identifier]pluggable.Expr
	ens   ensurable.Ensurable
}

func (ea *EnsureAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *EnsureAction) What() pluggable.SymbolType {
	return pluggable.SymbolType(ea.what.Id())
}

func (ea *EnsureAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("EnsureAction")
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
	return fmt.Sprintf("Ensure[%s: %s]", ea.what.Id(), ea.named.Text())
}

func (ea *EnsureAction) AddProperty(name pluggable.Identifier, value pluggable.Expr) {
	if name.Id() == "name" {
		if ea.named != nil {
			ea.tools.Reporter.Report(name.Loc().Offset, "duplicate definition of name")
			return
		}
		str, ok := value.(pluggable.String)
		if !ok {
			ea.tools.Reporter.Report(value.Loc().Offset, "name must be a string")
			return
		}
		ea.named = str
	} else {
		if ea.props[name] != nil {
			ea.tools.Reporter.Reportf(name.Loc().Offset, "duplicate definition of %s", name.Id())
			return
		}
		ea.props[name] = value
	}
}

func (ea *EnsureAction) AddAdverb(adv pluggable.Adverb, tokens []pluggable.Token) pluggable.Interpreter {
	if adv.Name() == "teardown" {
		if ea.teardown != nil {
			panic("duplicate teardown")
		}
		if len(tokens) != 1 {
			panic("invalid tokens")
		}
		ea.teardown = &MyTearDown{mode: tokens[0].(pluggable.Identifier).Id()}

	}
	return interpreters.DisallowInnerScope(ea.tools)
}

func (ea *EnsureAction) Completed() {
	if ea.tools.Reporter.HasErrors() {
		return
	}
	if ea.named == nil {
		ea.tools.Reporter.At(ea.loc.Line)
		ea.tools.Reporter.Report(ea.loc.Offset, "ensure requires a name to be defined")
	}
	if ea.teardown == nil {
		ea.tools.Reporter.At(ea.loc.Line)
		ea.tools.Reporter.Report(ea.loc.Offset, "ensure requires a teardown strategy to be declared")
	}
}

func (ea *EnsureAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	res, ok := r.Resolve(ea.what).(pluggable.Blank)
	if !ok {
		return pluggable.ERROR_OCCURRED
	}
	for _, y := range ea.props {
		y.Resolve(r)
	}
	ea.resolved = res
	obj := ea.resolved.Mint(ea.tools, ea.Loc(), ea.named.Text(), ea.props, ea.teardown)
	ens, ok := obj.(ensurable.Ensurable)
	if !ok {
		ea.tools.Storage.Errorf(ea.loc, "the type "+ea.what.Id()+" is not ensurable")
		return pluggable.ERROR_OCCURRED
	}
	ea.ens = ens
	// b.MayBind(ens)
	return pluggable.MAY_BE_BOUND
}

func (ea *EnsureAction) Prepare(pres pluggable.ValuePresenter) {
	ea.ens.Prepare(pres)
}

func (ea *EnsureAction) Execute() {
	ea.ens.Execute()
}

func (ea *EnsureAction) TearDown() {
	ea.ens.TearDown()
}

type MyTearDown struct {
	mode string
}

func (m *MyTearDown) Mode() string {
	return m.mode
}
