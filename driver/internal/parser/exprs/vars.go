package exprs

import (
	"fmt"
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type VarReference struct {
	id        driverbottom.Identifier
	actualVar driverbottom.Describable
}

func (a *VarReference) Resolve(r driverbottom.Resolver) {
	v := r.Resolve(a.id)
	a.actualVar = v
}

func (v *VarReference) Eval(s driverbottom.RuntimeStorage) any {
	// log.Printf("Eval(vr) %s %v => %T %v\n", v.id, v, s.Get(v), s.Get(v))
	out := s.Get(v)
	if out != nil {
		return out
	}
	out = s.Read(driverbottom.SymbolName(v.id.Id()))
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

func (v *VarReference) DumpTo(iw driverbottom.IndentWriter) {
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

func (v *VarReference) Named() driverbottom.Identifier {
	return v.id
}

func (a *VarReference) Binding() driverbottom.Describable {
	if a.actualVar == nil {
		// panic("help!")
		log.Fatalf("var was not resolved: %s %s\n", a.id.Id(), a.id.Loc().String())
	}
	return a.actualVar
}

func VarRefer(id driverbottom.Identifier) driverbottom.Var {
	return &VarReference{id: id}
}

func IsVar(e driverbottom.Expr, id driverbottom.Identifier) bool {
	v, ok := e.(*VarReference)
	if !ok {
		return false
	}
	return v.id == id
}
