package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type policyDocument struct {
	loc *errorsink.Location
	// name  string
	items []corebottom.PolicyEffect
}

func (p *policyDocument) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("PolicyDocument[]")
	w.AttrsWhere(p)
	w.EndAttrs()
}

func (p *policyDocument) Loc() *errorsink.Location {
	return p.loc
}

func (p *policyDocument) ShortDescription() string {
	return "PolicyDocument[]"
}

func (p *policyDocument) Item(effect string) corebottom.PolicyEffect {
	ret := &policyItem{effect: effect, more: make(map[string][]any)}
	p.items = append(p.items, ret)
	return ret
}

func (p *policyDocument) Items() []corebottom.PolicyEffect {
	return p.items
}

func NewPolicyDocument(loc *errorsink.Location) corebottom.PolicyDocument {
	return &policyDocument{loc: loc}
}
