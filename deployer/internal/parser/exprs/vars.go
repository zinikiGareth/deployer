package exprs

import (
	"log"

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
	log.Printf("Creating var reference for %s\n", id.Id())
	return &VarReference{id: id}
}

func IsVar(e pluggable.Expr, id pluggable.Identifier) bool {
	v, ok := e.(*VarReference)
	if !ok {
		return false
	}
	return v.id == id
}

type ActualVar struct {
}

// Binding implements pluggable.Var.
func (a *ActualVar) Binding() pluggable.Describable {
	panic("unimplemented")
}

// DumpTo implements pluggable.Var.
func (a *ActualVar) DumpTo(to pluggable.IndentWriter) {
	panic("unimplemented")
}

// Eval implements pluggable.Var.
func (a *ActualVar) Eval(s pluggable.RuntimeStorage) any {
	panic("unimplemented")
}

// Loc implements pluggable.Var.
func (a *ActualVar) Loc() *errors.Location {
	panic("unimplemented")
}

// Named implements pluggable.Var.
func (a *ActualVar) Named() pluggable.Identifier {
	panic("unimplemented")
}

// Resolve implements pluggable.Var.
func (a *ActualVar) Resolve(r pluggable.Resolver) {
	panic("unimplemented")
}

// ShortDescription implements pluggable.Var.
func (a *ActualVar) ShortDescription() string {
	panic("unimplemented")
}

// String implements pluggable.Var.
func (a *ActualVar) String() string {
	panic("unimplemented")
}

func SolidVar(id pluggable.Identifier) pluggable.Var {
	log.Printf("Creating actual var for %s\n", id.Id())
	return &ActualVar{}
}
