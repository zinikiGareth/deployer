package vars

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func MakeDoAssign(tools *corebottom.Tools, holder driverbottom.Holder, assignTo driverbottom.Identifier, asAny any) *DoAssign {
	action, ok := asAny.(driverbottom.Describable)
	if !ok {
		log.Printf("is not a describable: %T", asAny)
		panic("not a describable")
	}
	ret := DoAssign{tools: tools, assignTo: assignTo, holder: holder, action: action}
	return &ret
}

type DoAssign struct {
	tools    *corebottom.Tools
	assignTo driverbottom.Identifier
	holder   driverbottom.Holder
	action   driverbottom.Describable
	andThen  corebottom.ValuePresenter
}

func (da *DoAssign) Loc() *errorsink.Location {
	return da.assignTo.Loc()
}

func (da *DoAssign) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("AssignTo")
	w.AttrsWhere(da.assignTo)
	w.TextAttr("assignTo", da.assignTo.Id())
	da.action.DumpTo(w)
	w.EndAttrs()
}

func (da *DoAssign) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	res, ok := da.action.(driverbottom.Resolvable)
	if ok {
		status := res.Resolve(r)
		if status == driverbottom.NO_VALUE {
			panic("assignTo specified but expr does not produce a value") // should be an error
		}
		da.tools.Storage.EnableSymbol(da.holder)
	}

	return driverbottom.NO_VALUE
}

func (da *DoAssign) ShortDescription() string {
	return "DoAssign[" + da.assignTo.Id() + "<-" + da.action.ShortDescription() + "]"
}

func (da *DoAssign) Nested() driverbottom.Describable {
	return da.action
}

func (da *DoAssign) DetermineInitialState(pres corebottom.ValuePresenter) {
	da.andThen = pres
	fnd, ok := da.action.(corebottom.Findable)
	if !ok {
		// should we call something on pres? such as "notFinder"?
		return
	}
	fnd.DetermineInitialState(da)
}

func (da *DoAssign) DetermineDesiredState(pres corebottom.ValuePresenter) {
	da.andThen = pres
	mb, ok := da.action.(corebottom.ModelBuilder)
	if ok {
		mb.DetermineDesiredState(da)
		return
	}
	mc, ok := da.action.(corebottom.MemoryBuilder)
	if ok {
		mc.Create(da)
		return
	}
	// should we call something on pres? such as "onlyFinder"?
}

func (da *DoAssign) ShouldDestroy() bool {
	amis, ok := da.action.(corebottom.RealityShifter)
	if ok {
		return amis.ShouldDestroy()
	} else {
		return false
	}
}

func (da *DoAssign) UpdateReality() {
	amis, ok := da.action.(corebottom.RealityShifter)
	if ok {
		amis.UpdateReality()
	}
}

func (da *DoAssign) TearDown() {
	amis, ok := da.action.(corebottom.RealityShifter)
	if ok {
		amis.TearDown()
	}
}

func (da *DoAssign) NotFound() {
}

func (da *DoAssign) Present(value any) {
	if da.holder == nil { // can't do anything if we didn't resolve it
		return
	}
	da.tools.Storage.Bind(da.holder, value)
	da.andThen.Present(value)
}

func (da *DoAssign) WantDestruction(loc *errorsink.Location) {
	// TODO: I think we need to pass this on to the target
	panic("need to implement DoAssign.WantDestruction")
}

var _ corebottom.RealityShifter = &DoAssign{}
