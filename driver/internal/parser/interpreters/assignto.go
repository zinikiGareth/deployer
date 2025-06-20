package interpreters

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func WillAssignTo(tools *driverbottom.CoreTools, container driverbottom.AttachResult, assignTo driverbottom.Identifier) *WithAssignTo {
	holder := &VarHolder{storeFor: assignTo}
	ret := WithAssignTo{tools: tools, assignTo: assignTo, container: container, holder: holder}
	tools.Repository.IntroduceSymbol(driverbottom.SymbolName(assignTo.Id()), holder)

	return &ret
}

type WithAssignTo struct {
	tools     *driverbottom.CoreTools
	holder    *VarHolder
	assignTo  driverbottom.Identifier
	container driverbottom.AttachResult
}

func (wat *WithAssignTo) Attach(d any) {
	action, ok := d.(driverbottom.ModelBuilder)
	if !ok {
		panic("not an action")
	}
	wat.container.Attach(MakeDoAssign(wat.tools, wat.holder, wat.assignTo, action))
}

func MakeDoAssign(tools *driverbottom.CoreTools, holder *VarHolder, assignTo driverbottom.Identifier, action driverbottom.ModelBuilder) *DoAssign {
	ret := DoAssign{tools: tools, assignTo: assignTo, holder: holder, action: action}
	return &ret
}

type VarHolder struct {
	storeFor driverbottom.Identifier
}

func (v *VarHolder) Loc() *errorsink.Location {
	return v.storeFor.Loc()
}

func (v *VarHolder) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("VarHolder")
	iw.AttrsWhere(v.storeFor)
	iw.TextAttr("storeFor", v.storeFor.Id())
	iw.EndAttrs()
}

func (v *VarHolder) ShortDescription() string {
	return fmt.Sprintf("VarExpr[%s]", v.storeFor.Id())
}

type DoAssign struct {
	tools    *driverbottom.CoreTools
	assignTo driverbottom.Identifier
	holder   driverbottom.Describable
	action   driverbottom.ModelBuilder
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

func (d *DoAssign) BuildModel(pres driverbottom.ValuePresenter) {
	d.action.BuildModel(d)
}

func (d *DoAssign) UpdateReality() {
	amis, ok := d.action.(driverbottom.RealityShifter)
	if ok {
		amis.UpdateReality()
	}
}

func (d *DoAssign) TearDown() {
	amis, ok := d.action.(driverbottom.RealityShifter)
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
