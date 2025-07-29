package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type policyActionList struct {
	loc   *errorsink.Location
	items []corebottom.PolicyRuleAction
}

func (p *policyActionList) Add(r corebottom.PolicyRuleAction) {
	p.items = append(p.items, r)
}

func (p *policyActionList) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("PolicyActionList[]")
	w.AttrsWhere(p)
	w.EndAttrs()
}

func (p *policyActionList) Loc() *errorsink.Location {
	return p.loc
}

func (p *policyActionList) ShortDescription() string {
	return "PolicyActionList[]"
}

func (p *policyActionList) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, i := range p.items {
		i.Resolve(r)
	}
	return driverbottom.NO_VALUE
}

func (p *policyActionList) ApplyTo(doc corebottom.PolicyDocument) {
	for _, i := range p.items {
		i.ApplyTo(doc)
	}
}

func NewPolicyActionList(loc *errorsink.Location) corebottom.PolicyActionList {
	return &policyActionList{loc: loc}
}

var _ corebottom.PolicyActionList = &policyActionList{}
