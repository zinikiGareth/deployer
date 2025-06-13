package exprs

import (
	"fmt"
	"log"

	"ziniki.org/deployer/deployer/pkg/errorsink"
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
	// log.Printf("Eval(vr) %s %v => %T %v\n", v.id, v, s.Get(v), s.Get(v))
	out := s.Get(v)
	if out != nil {
		return out
	}
	out = s.Read(pluggable.SymbolName(v.id.Id()))
	if out != nil {
		return out
	}
	panic(fmt.Sprintf("cannot find %v\n", v))
}

func (v *VarReference) Loc() *errorsink.Location {
	return v.id.Loc()
}

func (v *VarReference) ShortDescription() string {
	return v.Loc().String() + " Var[" + v.id.Id() + "]"
}

func (v *VarReference) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("Var %s", v.id)
	iw.AttrsWhere(v)
	if v.actualVar != nil {
		iw.NestedAttr("actualVar", v.actualVar)
	}
	iw.EndAttrs()
}

func (v *VarReference) String() string {
	return "Var[" + v.id.Id() + "]"
}

func (v *VarReference) Named() pluggable.Identifier {
	return v.id
}

func (a *VarReference) Binding() pluggable.Describable {
	if a.actualVar == nil {
		// panic("help!")
		log.Fatalf("var was not resolved: %s %s\n", a.id.Id(), a.id.Loc().String())
	}
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
