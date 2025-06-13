package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyRuleAction interface {
	pluggable.Describable
	Resolve(r pluggable.Resolver) pluggable.BindingRequirement
	ApplyTo(doc external.PolicyDocument)
}

type PolicyAllowAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	actions []pluggable.Describable

	allowActions    []pluggable.Expr
	allowResources  []pluggable.Expr
	allowPrincipals []pluggable.Expr
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
	w.ListAttr("allowPrincipals")
	for _, a := range paa.allowPrincipals {
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
	paa.actions = append(paa.actions, entry.(pluggable.Describable))
}

func (paa *PolicyAllowAction) ApplyTo(doc external.PolicyDocument) {
	pd, ok := doc.(*policyDocument)
	if !ok {
		panic("internal error")
	}
	actions := []string{}
	resources := []string{}
	principals := []string{}
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
	for _, p := range paa.allowPrincipals {
		// v := paa.tools.Storage.Eval(r)
		// log.Printf("r = %T %v %T %v\n", r, r, v, v)
		principals = append(principals, paa.tools.Storage.EvalAsString(p))
	}
	pd.items = append(pd.items, &policyItem{effect: "Allow", actions: actions, resources: resources, principals: principals})
}

func (paa *PolicyAllowAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	for _, a := range paa.allowActions {
		// log.Printf("need to resolve %T %v\n", a, a)
		a.Resolve(r)
	}
	for _, ar := range paa.allowResources {
		// log.Printf("need to resolve %T %v\n", ar, ar)
		ar.Resolve(r)
	}
	for _, ap := range paa.allowPrincipals {
		// log.Printf("need to resolve %T %v\n", ap, ap)
		ap.Resolve(r)
	}
	return pluggable.MAY_BE_BOUND
}

func (paa *PolicyAllowAction) BuildModel(pres pluggable.ValuePresenter) {
}
