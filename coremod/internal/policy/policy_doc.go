package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyDocument struct {
	loc *errorsink.Location
}

// DumpTo implements pluggable.Describable.
func (p *PolicyDocument) DumpTo(w pluggable.IndentWriter) {
	w.Intro("PolicyDocument[]")
	w.AttrsWhere(p)
	w.EndAttrs()
}

// Loc implements pluggable.Describable.
func (p *PolicyDocument) Loc() *errorsink.Location {
	return p.loc
}

// ShortDescription implements pluggable.Describable.
func (p *PolicyDocument) ShortDescription() string {
	return "PolicyDocument[]"
}
