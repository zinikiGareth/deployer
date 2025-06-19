package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type policyDocument struct {
	loc *errorsink.Location
	// name  string
	items []external.PolicyEffect
}

// DumpTo implements driverbottom.Describable.
func (p *policyDocument) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("PolicyDocument[]")
	w.AttrsWhere(p)
	w.EndAttrs()
}

// Loc implements driverbottom.Describable.
func (p *policyDocument) Loc() *errorsink.Location {
	return p.loc
}

// ShortDescription implements driverbottom.Describable.
func (p *policyDocument) ShortDescription() string {
	return "PolicyDocument[]"
}

// func (p *policyDocument) Name() string {
// 	return p.name
// }

func (p *policyDocument) Item(effect string) external.PolicyEffect {
	ret := &policyItem{effect: effect, more: make(map[string][]any)}
	p.items = append(p.items, ret)
	return ret
}

func (p *policyDocument) Items() []external.PolicyEffect {
	return p.items
}

func NewPolicyDocument(loc *errorsink.Location) external.PolicyDocument {
	return &policyDocument{loc: loc}
}
