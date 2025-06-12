package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/findable"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type FindAction struct {
	tools    *pluggable.Tools
	loc      *errorsink.Location
	what     pluggable.Identifier
	resolved pluggable.Blank
	named    pluggable.String
	props    map[pluggable.Identifier]pluggable.Expr
	ens      findable.Findable
}

func (ea *FindAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *FindAction) What() pluggable.SymbolType {
	return pluggable.SymbolType(ea.what.Id())
}

func (ea *FindAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("FindAction")
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

func (ea *FindAction) ShortDescription() string {
	return fmt.Sprintf("Find[%s: %s]", ea.what.Id(), ea.named.Text())
}

func (ea *FindAction) AddProperty(name pluggable.Identifier, value pluggable.Expr) {
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

func (ea *FindAction) AddAdverb(adv pluggable.Adverb, tokens []pluggable.Token) pluggable.Interpreter {
	ea.tools.Reporter.Reportf(adv.Loc().Offset, "find cannot handle %s\n", adv.Name())
	return interpreters.DisallowInnerScope(ea.tools)
}

func (ea *FindAction) Completed() {
	if ea.tools.Reporter.HasErrors() {
		return
	}
	if ea.named == nil {
		ea.tools.Reporter.At(ea.loc.Line)
		ea.tools.Reporter.Report(ea.loc.Offset, "Find requires a name to be defined")
	}
}

func (ea *FindAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	res, ok := r.Resolve(ea.what).(pluggable.Blank)
	if !ok {
		return pluggable.ERROR_OCCURRED
	}
	ea.resolved = res
	obj := ea.resolved.Find(ea.tools, ea.Loc(), ea.named.Text())
	ens, ok := obj.(findable.Findable)
	if !ok {
		ea.tools.Storage.Errorf(ea.loc, "the type "+ea.what.Id()+" is not findable")
		return pluggable.ERROR_OCCURRED
	}
	ea.ens = ens
	// b.MayBind(ens)
	return pluggable.MAY_BE_BOUND
}

func (ea *FindAction) Prepare(pres pluggable.ValuePresenter) {
	ea.ens.Prepare(pres)
}

func (ea *FindAction) Execute() {
	ea.ens.Execute()
}

func (ea *FindAction) TearDown() {
}
