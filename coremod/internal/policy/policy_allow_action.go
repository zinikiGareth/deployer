package policy

import (
	"ziniki.org/deployer/coremod/internal/vars"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type PolicyAllowAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location

	allowActions   []driverbottom.Expr
	allowResources []driverbottom.Expr
	updates        []corebottom.UpdatePolicyAllowAction
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
	w.ListAttr("updates")
	for _, a := range paa.updates {
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

func (paa *PolicyAllowAction) Attach(entry any) error {
	paa.updates = append(paa.updates, entry.(corebottom.UpdatePolicyAllowAction))
	return nil
}

func (paa *PolicyAllowAction) ApplyTo(doc corebottom.PolicyDocument) {
	item := doc.Item("Allow")
	for _, a := range paa.allowActions {
		a1, ok := paa.tools.Storage.EvalAsStringer(a)
		if !ok {
			panic("not a stringer")
		}
		// log.Printf("have action %T %p: %s\n", a1, a1, a1.String())
		item.Action(a1.String())
	}
	for _, r := range paa.allowResources {
		r1, ok := paa.tools.Storage.EvalAsStringer(r)
		if !ok {
			panic("not a stringer")
		}
		// log.Printf("have resource %T %p: %s\n", r1, r1, r1.String())
		item.Resource(r1.String())
	}

	for _, aa := range paa.updates {
		// log.Printf("have update %T %p\n", aa, aa)
		aa.ApplyTo(item)
	}
}

func (paa *PolicyAllowAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND
	for _, a := range paa.updates {
		if a.Resolve(r) == driverbottom.ERROR_OCCURRED {
			ret = driverbottom.ERROR_OCCURRED
		}
	}
	for _, a := range paa.allowActions {
		// log.Printf("need to resolve %T %v\n", a, a)
		if a.Resolve(r) == driverbottom.ERROR_OCCURRED {
			ret = driverbottom.ERROR_OCCURRED
		}
	}
	for _, ar := range paa.allowResources {
		// log.Printf("need to resolve %T %v\n", ar, ar)
		if ar.Resolve(r) == driverbottom.ERROR_OCCURRED {
			ret = driverbottom.ERROR_OCCURRED
		}
	}
	return ret
}

func NewPolicyAllowAction(tools *corebottom.Tools, loc *errorsink.Location, actions []driverbottom.Expr, resources []driverbottom.Expr, updates []corebottom.UpdatePolicyAllowAction) corebottom.PolicyRuleAction {
	return &PolicyAllowAction{tools: tools, loc: loc, allowActions: actions, allowResources: resources, updates: updates}
}

var _ corebottom.PolicyRuleAction = &PolicyAllowAction{}
