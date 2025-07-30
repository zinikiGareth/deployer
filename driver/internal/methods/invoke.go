package methods

import (
	"log"
	"os"
	"strings"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type InvokeExpr struct {
	driverbottom.Locatable
	on   driverbottom.Expr
	call driverbottom.Identifier
	args []driverbottom.Expr
}

func (i *InvokeExpr) Loc() *errorsink.Location {
	return i.Locatable.Loc()
}

func (i *InvokeExpr) ShortDescription() string {
	var sb strings.Builder
	sb.WriteString("Invoke{")
	sb.WriteString(i.on.ShortDescription())
	sb.WriteString("->")
	sb.WriteString(i.call.String())
	sb.WriteRune('(')
	for i, a := range i.args {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(a.ShortDescription())
	}
	sb.WriteString(")}")
	return sb.String()
}

func (i *InvokeExpr) String() string {
	return i.ShortDescription()
}

func (i *InvokeExpr) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("Invoke")
	iw.AttrsWhere(i)
	iw.NestedAttr("on", i.on)
	iw.NestedAttr("meth", i.call)
	iw.ListAttr("args")
	for _, a := range i.args {
		a.DumpTo(iw)
	}
	iw.EndList()
	iw.EndAttrs()
}

func (i *InvokeExpr) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND
	i.on.Resolve(r)
	for _, e := range i.args {
		e.Resolve(r)
	}
	return ret
}

func (i *InvokeExpr) Eval(s driverbottom.RuntimeStorage) any {
	// log.Printf("on = %T %v\n", i.on, i.on)
	obj := i.on.Eval(s)
	hm, ok := obj.(driverbottom.HasMethods)
	if !ok {
		s.DumpTo(os.Stderr)
		log.Printf("Value for %v was of type %T which was not a HasMethods\n", i.on, obj)
		panic("could not evaluate")
	}
	meth := hm.ObtainMethod(i.call.Id())
	if meth == nil {
		log.Fatalf("No method %s on %T %v", i.call.Id(), i.on, i.on)
	}
	return meth.Invoke(s, i.on, i.args)
}

func MakeInvokeExpr(on driverbottom.Expr, call driverbottom.Identifier, args []driverbottom.Expr) driverbottom.Expr {
	return &InvokeExpr{Locatable: on, on: on, call: call, args: args}
}

type InvokeFunc struct {
	tools *driverbottom.CoreTools
}

func (i *InvokeFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (i *InvokeFunc) Precedence() int {
	return 10
}

func (i *InvokeFunc) Associativity() bool {
	return true
}

func (i *InvokeFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) driverbottom.Expr {
	if len(before) != 1 {
		panic("should be an error")
	}
	if len(after) < 1 {
		panic("should be an error")
	}
	meth, ok := after[0].(driverbottom.Var)
	if !ok {
		panic("should be an error")
	}
	return &InvokeExpr{Locatable: me, on: before[0], call: meth.Named(), args: after[1:]}
}

func MakeInvokeFunc(tools *driverbottom.CoreTools) *InvokeFunc {
	return &InvokeFunc{tools: tools}
}

var _ driverbottom.Function = &InvokeFunc{}
var _ driverbottom.Expr = &InvokeExpr{}
