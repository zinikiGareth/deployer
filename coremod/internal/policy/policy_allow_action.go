package policy

import (
	"ziniki.org/deployer/coremod/internal/vars"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// Is probably more general than this, but who knows?
type UpdatePolicyAllowAction interface {
	driverbottom.Describable
	Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement
	ApplyTo(doc corebottom.PolicyEffect)
}

type PolicyAllowAction struct {
	tools   *corebottom.Tools
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

func (paa *PolicyAllowAction) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	ret := vars.MakeDoAssign(paa.tools, holder, assignTo, action)
	return ret
}

func (paa *PolicyAllowAction) Attach(entry any) {
	paa.actions = append(paa.actions, entry.(UpdatePolicyAllowAction))
}

func (paa *PolicyAllowAction) ApplyTo(doc corebottom.PolicyDocument) {
	item := doc.Item("Allow")
	for _, a := range paa.allowActions {
		a1, ok := paa.tools.Storage.EvalAsStringer(a)
		if !ok {
			panic("not a stringer")
		}
		item.Action(a1.String())
	}
	for _, r := range paa.allowResources {
		r1, ok := paa.tools.Storage.EvalAsStringer(r)
		if !ok {
			panic("not a stringer")
		}
		item.Resource(r1.String())
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

func (paa *PolicyAllowAction) DetermineDesiredState(pres corebottom.ValuePresenter) {
}
