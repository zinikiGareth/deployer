package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	actions []PolicyRuleAction
}

func (pa *PolicyAction) Attach(entry any) {
	a, ok := entry.(PolicyRuleAction)
	if !ok {
		panic("throw an error")
	}
	pa.actions = append(pa.actions, a)
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
	return pluggable.MUST_BE_BOUND
}

func (pa *PolicyAction) BuildModel(pres pluggable.ValuePresenter) {
	doc := NewPolicyDocument(pa.loc)

	for _, a := range pa.actions {
		a.ApplyTo(doc)
	}

	pres.Present(doc)
}
