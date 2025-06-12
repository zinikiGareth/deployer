package interpreters

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type WithAssignTo struct {
	tools     *pluggable.Tools
	assignTo  pluggable.Identifier
	container pluggable.AttachResult
}

func (wat *WithAssignTo) Attach(d any) {
	container, ok := d.(pluggable.Action)
	if !ok {
		panic("not an action")
	}
	wat.container.Attach(MakeDoAssign(wat.tools, wat.assignTo, container))
}

func MakeDoAssign(tools *pluggable.Tools, assignTo pluggable.Identifier, action pluggable.Action) *DoAssign {
	holder := &VarHolder{storeFor: assignTo}
	ret := DoAssign{tools: tools, assignTo: assignTo, holder: holder, action: action}
	tools.Repository.IntroduceSymbol(pluggable.SymbolName(assignTo.Id()), holder)
	return &ret
}

type VarHolder struct {
	storeFor pluggable.Identifier
}

func (v *VarHolder) Loc() *errorsink.Location {
	return v.storeFor.Loc()
}

func (v *VarHolder) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("VarHolder")
	iw.AttrsWhere(v.storeFor)
	iw.TextAttr("storeFor", v.storeFor.Id())
	iw.EndAttrs()
}

func (v *VarHolder) ShortDescription() string {
	return fmt.Sprintf("VarExpr[%s]", v.storeFor.Id())
}

type DoAssign struct {
	tools    *pluggable.Tools
	assignTo pluggable.Identifier
	holder   pluggable.Describable
	action   pluggable.Action
}

func (d *DoAssign) Loc() *errorsink.Location {
	return d.assignTo.Loc()
}

func (d *DoAssign) DumpTo(w pluggable.IndentWriter) {
	w.Intro("AssignTo")
	w.AttrsWhere(d.assignTo)
	w.TextAttr("assignTo", d.assignTo.Id())
	d.action.DumpTo(w)
	w.EndAttrs()
}

func (d *DoAssign) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	status := d.action.Resolve(r)
	if status == pluggable.NO_VALUE {
		panic("assignTo specified but expr does not produce a value") // should be an error
	}
	return pluggable.NO_VALUE
}

func (d *DoAssign) ShortDescription() string {
	return "DoAssign[" + d.assignTo.Id() + "<-" + d.action.ShortDescription() + "]"
}

func (d *DoAssign) Prepare(pres pluggable.ValuePresenter) {
	d.action.Prepare(d)
}

func (d *DoAssign) Execute() {
	amis, ok := d.action.(pluggable.AndMakeItSo)
	if ok {
		amis.Execute()
	}
}

func (d *DoAssign) TearDown() {
	amis, ok := d.action.(pluggable.AndMakeItSo)
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
