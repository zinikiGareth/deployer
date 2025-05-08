package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/ensurable"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type EnsureAction struct {
	tools    *pluggable.Tools
	loc      *errors.Location
	what     pluggable.Identifier
	resolved pluggable.Blank
	named    pluggable.String
	props    map[pluggable.Identifier]pluggable.Expr
	ens      ensurable.Ensurable
}

func (ea *EnsureAction) Loc() *errors.Location {
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

func (ea *EnsureAction) Completed() {
	if ea.named == nil {
		ea.tools.Reporter.At(ea.loc.Line)
		ea.tools.Reporter.Report(ea.loc.Offset, "ensure requires a name to be defined")
	}
}

func (ea *EnsureAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	ea.resolved = r.Resolve(ea.what)
	obj := ea.resolved.Mint(ea.tools, ea.Loc(), ea.named.Text())
	ens, ok := obj.(ensurable.Ensurable)
	if !ok {
		ea.tools.Storage.Errorf(ea.loc, "the type "+ea.what.Id()+" is not ensurable")
		return
	}
	ea.ens = ens
	b.MayBind(ens)
}

func (ea *EnsureAction) Prepare(pres pluggable.ValuePresenter) {
	ea.ens.Prepare()
}

func (ea *EnsureAction) Execute() {
	ea.ens.Execute()
}
