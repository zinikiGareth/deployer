package policy

import (
	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type PolicyRuleAction interface {
	driverbottom.Describable
	Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement
	ApplyTo(doc corebottom.PolicyDocument)
}

type PolicyAction struct {
	tools   *corebottom.Tools
	loc     *errorsink.Location
	actions []PolicyRuleAction
}

func (pa *PolicyAction) MakeAssign(holder driverbottom.Describable, assignTo driverbottom.Identifier, action driverbottom.ModelBuilder) any {
	ret := target.MakeDoAssign(pa.tools, holder, assignTo, action)
	return ret
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

func (pa *PolicyAction) DumpTo(w driverbottom.IndentWriter) {
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

func (pa *PolicyAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, a := range pa.actions {
		a.Resolve(r)
	}
	return driverbottom.MUST_BE_BOUND
}

func (pa *PolicyAction) BuildModel(pres driverbottom.ValuePresenter) {
	doc := NewPolicyDocument(pa.loc)

	for _, a := range pa.actions {
		a.ApplyTo(doc)
	}

	pres.Present(doc)
}
