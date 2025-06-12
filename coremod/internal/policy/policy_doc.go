package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyDocument struct {
	loc *errorsink.Location
}

// DumpTo implements pluggable.Describable.
func (p *policyDocument) DumpTo(w pluggable.IndentWriter) {
	w.Intro("PolicyDocument[]")
	w.AttrsWhere(p)
	w.EndAttrs()
}

// Loc implements pluggable.Describable.
func (p *policyDocument) Loc() *errorsink.Location {
	return p.loc
}

// ShortDescription implements pluggable.Describable.
func (p *policyDocument) ShortDescription() string {
	return "PolicyDocument[]"
}
