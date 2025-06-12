package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	actions []pluggable.Describable
	doc     *PolicyDocument
}

// This feels weird to me and I'm not sure what to do about it.
// It's the "action" bit that seems wrong, when I expect to be assembling an object
// But I think the thing is that we need the idea of "actions" to make the whole "pineal" model work
func (pa *PolicyAction) Attach(entry any) {
	pa.actions = append(pa.actions, entry.(pluggable.Describable))
}

func (pa *PolicyAction) Loc() *errorsink.Location {
	return pa.loc
}

func (pa *PolicyAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("PolicyAction")
	w.AttrsWhere(pa)
	w.ListAttr("actions")
	for _, a := range pa.actions {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (pa *PolicyAction) ShortDescription() string {
	return "Policy[]"
}

func (pa *PolicyAction) Completed() {
}

func (pa *PolicyAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	pa.doc = &PolicyDocument{loc: pa.loc}
	// b.MustBind(pa.doc)
	return pluggable.MUST_BE_BOUND
}

func (pa *PolicyAction) Prepare(pres pluggable.ValuePresenter) {
}
