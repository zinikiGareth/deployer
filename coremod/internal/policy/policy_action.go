package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyAction struct {
	tools *pluggable.Tools
	loc   *errorsink.Location
	doc   *PolicyDocument
}

func (pa *PolicyAction) Loc() *errorsink.Location {
	return pa.loc
}

func (pa *PolicyAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("PolicyAction")
	w.AttrsWhere(pa)
	// w.TextAttr("varname", pa.varname.String())
	w.EndAttrs()
}

func (pa *PolicyAction) ShortDescription() string {
	return "Policy[]"
}

func (pa *PolicyAction) Completed() {
}

func (pa *PolicyAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	pa.doc = &PolicyDocument{loc: pa.loc}
	b.MustBind(pa.doc)
}

func (pa *PolicyAction) Prepare(pres pluggable.ValuePresenter) {
}
