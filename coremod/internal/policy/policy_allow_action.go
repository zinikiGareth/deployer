package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyAllowAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	actions []pluggable.Action

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

func (paa *PolicyAllowAction) Add(entry pluggable.Action) {
	paa.actions = append(paa.actions, entry)
}

func (paa *PolicyAllowAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	return pluggable.MAY_BE_BOUND
}

func (paa *PolicyAllowAction) Prepare(pres pluggable.ValuePresenter) {
}
