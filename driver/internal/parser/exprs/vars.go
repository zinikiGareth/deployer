package exprs

import (
	"fmt"
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type VarReference struct {
	scope     driverbottom.Scope
	id        driverbottom.Identifier
	isValue   bool
	value     driverbottom.Describable
	actualVar driverbottom.Holder
}

func (v *VarReference) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND
	val := r.Resolve(v.scope, v.id)
	if val == nil {
		return driverbottom.ERROR_OCCURRED
	}
	h, ok := val.(driverbottom.Holder)
	if ok { // it is a "real var"
		v.actualVar = h
	} else {
		d, ok := val.(driverbottom.Describable)
		if ok {
			// it resolved to some other value, such as a constant number or string
			v.isValue = true
			v.value = d
		} else {
			r.ErrorAtf(v.Loc(), "resolution was not to a describable")
			return driverbottom.ERROR_OCCURRED
		}
	}
	return ret
}

func (v *VarReference) Eval(s driverbottom.RuntimeStorage) any {
	if v.isValue {
		// it didn't resolve to a variable but a value
		return v.value
	}
	if v.actualVar == nil {
		log.Printf("in processing of Resolve(), nobody remembered to ask for variable %s to be resolved, it seems (used at %s)\n", v.id.Id(), v.id.Loc().String())
		panic("resolve never called on " + v.id.Id())
	}
	// log.Printf("Eval(vr) %s %v => %T %v\n", v.id, v, s.Get(v.actualVar), s.Get(v.actualVar))
	out := s.Get(v.actualVar)
	return out
}

func (v *VarReference) Loc() *errorsink.Location {
	return v.id.Loc()
}

func (v *VarReference) ShortDescription() string {
	return v.Loc().String() + " Var[" + v.id.Id() + "]"
}

func (v *VarReference) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("Var %s", v.id.Id())
	iw.AttrsWhere(v)
	iw.TextAttr("isValue", fmt.Sprintf("%v", v.isValue))
	if v.isValue {
		iw.NestedAttr("value", v.value)
	}
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

func VarRefer(fromScope driverbottom.Scope, id driverbottom.Identifier) driverbottom.Var {
	return &VarReference{scope: fromScope, id: id}
}

func IsVar(e driverbottom.Expr, id driverbottom.Identifier) bool {
	v, ok := e.(*VarReference)
	if !ok {
		return false
	}
	return v.id == id
}
