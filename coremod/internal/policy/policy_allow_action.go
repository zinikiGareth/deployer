package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// Is probably more general than this, but who knows?
type UpdatePolicyAllowAction interface {
	driverbottom.Describable
	Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement
	ApplyTo(doc external.PolicyEffect)
}

type PolicyAllowAction struct {
	tools   *external.Tools
	loc     *errorsink.Location
	actions []UpdatePolicyAllowAction

	allowActions   []driverbottom.Expr
	allowResources []driverbottom.Expr
}

func (paa *PolicyAllowAction) Loc() *errorsink.Location {
	return paa.loc
}

func (paa *PolicyAllowAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("PolicyAllowAction")
	w.AttrsWhere(paa)
	w.ListAttr("allowActions")
	for _, a := range paa.allowActions {
		a.DumpTo(w)
	}
	w.EndList()
	w.ListAttr("allowResources")
	for _, a := range paa.allowResources {
		a.DumpTo(w)
	}
	w.EndList()
	w.ListAttr("actions")
	for _, a := range paa.actions {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (paa *PolicyAllowAction) ShortDescription() string {
	return "PolicyAllow[]"
}

func (paa *PolicyAllowAction) Completed() {
}

func (paa *PolicyAllowAction) Attach(entry any) {
	paa.actions = append(paa.actions, entry.(UpdatePolicyAllowAction))
}

func (paa *PolicyAllowAction) ApplyTo(doc external.PolicyDocument) {
	item := doc.Item("Allow")
	for _, a := range paa.allowActions {
		item.Action(paa.tools.Storage.EvalAsString(a))
	}
	for _, r := range paa.allowResources {
		item.Resource(paa.tools.Storage.EvalAsString(r))
	}

	for _, aa := range paa.actions {
		aa.ApplyTo(item)
	}
}

func (paa *PolicyAllowAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, a := range paa.actions {
		a.Resolve(r)
	}
	for _, a := range paa.allowActions {
		// log.Printf("need to resolve %T %v\n", a, a)
		a.Resolve(r)
	}
	for _, ar := range paa.allowResources {
		// log.Printf("need to resolve %T %v\n", ar, ar)
		ar.Resolve(r)
	}
	return driverbottom.MAY_BE_BOUND
}

func (paa *PolicyAllowAction) BuildModel(pres driverbottom.ValuePresenter) {
}
