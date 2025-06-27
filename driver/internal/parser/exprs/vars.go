package exprs

import (
	"fmt"
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type VarReference struct {
	id        driverbottom.Identifier
	value     driverbottom.Describable
	actualVar driverbottom.Holder
}

func (v *VarReference) Resolve(r driverbottom.Resolver) {
	val := r.Resolve(v.id)
	h, ok := val.(driverbottom.Holder)
	if ok { // it is a "real var"
		v.actualVar = h
	} else {
		// it resolved to some other value, such as a constant number or string
		v.value = h
	}
}

func (v *VarReference) Eval(s driverbottom.RuntimeStorage) any {
	if v.actualVar == nil {
		// it didn't resolve to a variable but a value
		return v.value
	}
	// log.Printf("Eval(vr) %s %v => %T %v\n", v.id, v, s.Get(v), s.Get(v))
	out := s.Get(v.actualVar)
	if out != nil {
		return out
	}
	out = s.Read(driverbottom.SymbolName(v.id.Id()))
	if out != nil {
		e, isExpr := out.(driverbottom.Expr)
		if isExpr {
			return e.Eval(s)
		} else {
			return out
		}
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

func (v *VarReference) Binding() driverbottom.Holder {
	if v.actualVar == nil {
		// panic("help!")
		log.Fatalf("var was not resolved: %s %s\n", v.id.Id(), v.id.Loc().String())
	}
	return v.actualVar
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
