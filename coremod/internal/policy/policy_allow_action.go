package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// Is probably more general than this, but who knows?
type UpdatePolicyAllowAction interface {
	pluggable.Describable
	Resolve(r pluggable.Resolver) pluggable.BindingRequirement
	ApplyTo(doc *policyItem)
}

type PolicyAllowAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	actions []UpdatePolicyAllowAction

	allowActions   []pluggable.Expr
	allowResources []pluggable.Expr
}

func (paa *PolicyAllowAction) Loc() *errorsink.Location {
	return paa.loc
}

func (paa *PolicyAllowAction) DumpTo(w pluggable.IndentWriter) {
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
	pd, ok := doc.(*policyDocument)
	if !ok {
		panic("internal error")
	}
	actions := []string{}
	resources := []string{}
	for _, a := range paa.allowActions {
		// v := paa.tools.Storage.Eval(a)
		// log.Printf("a = %T %v %T %v\n", a, a, v, v)
		actions = append(actions, paa.tools.Storage.EvalAsString(a))
	}
	for _, r := range paa.allowResources {
		// v := paa.tools.Storage.Eval(r)
		// log.Printf("r = %T %v %T %v\n", r, r, v, v)
		resources = append(resources, paa.tools.Storage.EvalAsString(r))
	}
	item := &policyItem{effect: "Allow", actions: actions, resources: resources}
	pd.items = append(pd.items, item)

	for _, aa := range paa.actions {
		aa.ApplyTo(item)
	}
}

func (paa *PolicyAllowAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
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
	return pluggable.MAY_BE_BOUND
}

func (paa *PolicyAllowAction) BuildModel(pres pluggable.ValuePresenter) {
}
