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

func (wat *WithAssignTo) Attach(action any) {
	assign := wat.container.MakeAssign(wat.holder, wat.assignTo, action)
	// log.Printf("actually attaching %p with %p to %p\n", assign, action, wat.container)
	wat.container.Attach(assign)
}

func (wat *WithAssignTo) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	panic("I *am* the variable")
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

func (v *VarHolder) VarName() driverbottom.Identifier {
	return v.storeFor
}

func (v *VarHolder) Resolve(r driverbottom.Resolver) {
}

func (v *VarHolder) Eval(s driverbottom.RuntimeStorage) any {
	return s.Get(v)
}

func (v *VarHolder) String() string {
	return v.ShortDescription()
}

func NewVarHolder(name driverbottom.Identifier) driverbottom.Holder {
	return &VarHolder{storeFor: name}
}

var _ driverbottom.Holder = &VarHolder{}
var _ driverbottom.Expr = &VarHolder{}
