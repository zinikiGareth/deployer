package exprs

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type VarReference struct {
	id        pluggable.Identifier
	actualVar pluggable.Describable
}

func (a *VarReference) Resolve(r pluggable.Resolver) {
	v := r.Resolve(a.id)
	a.actualVar = v
}

func (v *VarReference) Eval(s pluggable.RuntimeStorage) any {
	return s.Get(v)
}

func (v *VarReference) Loc() *errors.Location {
	return v.id.Loc()
}

func (v *VarReference) ShortDescription() string {
	return v.Loc().String() + " Var[" + v.id.Id() + "]"
}

func (t *VarReference) DumpTo(iw pluggable.IndentWriter) {
	panic("not implemented")
}

func (v *VarReference) String() string {
	return "Var[" + v.id.Id() + "]"
}

func (a *VarReference) Named() pluggable.Identifier {
	panic("unimplemented")
}

func (a *VarReference) Binding() pluggable.Describable {
	return a.actualVar
}

func VarRefer(id pluggable.Identifier) pluggable.Var {
	return &VarReference{id: id}
}

func IsVar(e pluggable.Expr, id pluggable.Identifier) bool {
	v, ok := e.(*VarReference)
	if !ok {
		return false
	}
	return v.id == id
}
