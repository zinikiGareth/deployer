package vars

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func MakeDoAssign(tools *corebottom.Tools, holder driverbottom.Describable, assignTo driverbottom.Identifier, asAny any) *DoAssign {
	action, ok := asAny.(corebottom.ModelBuilder)
	if !ok {
		log.Fatalf("is not a model builder: %T", asAny)
	}
	ret := DoAssign{tools: tools, assignTo: assignTo, holder: holder, action: action}
	return &ret
}

type DoAssign struct {
	tools    *corebottom.Tools
	assignTo driverbottom.Identifier
	holder   driverbottom.Describable
	action   corebottom.ModelBuilder
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
	status := d.action.Resolve(r)
	if status == driverbottom.NO_VALUE {
		panic("assignTo specified but expr does not produce a value") // should be an error
	}
	return driverbottom.NO_VALUE
}

func (d *DoAssign) ShortDescription() string {
	return "DoAssign[" + d.assignTo.Id() + "<-" + d.action.ShortDescription() + "]"
}

func (d *DoAssign) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	d.action.DetermineDesiredState(d)
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

func (d *DoAssign) Present(value any) {
	if d.holder == nil { // can't do anything if we didn't resolve it
		return
	}
	d.tools.Storage.Bind(d.holder, value)
}

func (d *DoAssign) Unchanged() {
	/*
		if d.holder == nil { // can't do anything if we didn't resolve it
			return
		}
		d.tools.Storage.Bind(d.holder, value)
	*/
}
