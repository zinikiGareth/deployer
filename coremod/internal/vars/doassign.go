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
}

func (d *DoAssign) Loc() *errorsink.Location {
	return d.assignTo.Loc()
}

func (d *DoAssign) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("AssignTo")
	w.AttrsWhere(d.assignTo)
	w.TextAttr("assignTo", d.assignTo.Id())
	d.action.DumpTo(w)
	w.EndAttrs()
}

func (d *DoAssign) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	res, ok := d.action.(driverbottom.Resolvable)
	if ok {
		status := res.Resolve(r)
		if status == driverbottom.NO_VALUE {
			panic("assignTo specified but expr does not produce a value") // should be an error
		}
		d.tools.Storage.EnableSymbol(d.assignTo)
	}

	return driverbottom.NO_VALUE
}

func (d *DoAssign) ShortDescription() string {
	return "DoAssign[" + d.assignTo.Id() + "<-" + d.action.ShortDescription() + "]"
}

func (d *DoAssign) DetermineInitialState(pres corebottom.ValuePresenter) {
	fnd, ok := d.action.(corebottom.Findable)
	if !ok {
		// should we call something on pres? such as "notFinder"?
		return
	}
	fnd.DetermineInitialState(d)
}

func (d *DoAssign) DetermineDesiredState(pres corebottom.ValuePresenter) {
	mb, ok := d.action.(corebottom.ModelBuilder)
	if !ok {
		// should we call something on pres? such as "onlyFinder"?
		return
	}
	mb.DetermineDesiredState(d)
}

func (d *DoAssign) UpdateReality() {
	amis, ok := d.action.(corebottom.RealityShifter)
	if ok {
		amis.UpdateReality()
	}
}

func (d *DoAssign) TearDown() {
	amis, ok := d.action.(corebottom.RealityShifter)
	if ok {
		amis.TearDown()
	}
}

func (d *DoAssign) NotFound() {
}

func (d *DoAssign) Present(value any) {
	if d.holder == nil { // can't do anything if we didn't resolve it
		return
	}
	d.tools.Storage.Bind(d.holder, value)
}

var _ corebottom.ModelBuilder = &DoAssign{}
