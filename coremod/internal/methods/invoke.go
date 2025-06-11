package methods

import (
	"strings"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type InvokeExpr struct {
	pluggable.Locatable
	on   pluggable.Expr
	call pluggable.Identifier
	args []pluggable.Expr
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
	panic("unimplemented")
}

func (i *InvokeExpr) DumpTo(to pluggable.IndentWriter) {
	panic("unimplemented")
}

func (i *InvokeExpr) Resolve(r pluggable.Resolver) {
	i.on.Resolve(r)
	for _, e := range i.args {
		e.Resolve(r)
	}
}

func (i *InvokeExpr) Eval(s pluggable.RuntimeStorage) any {
	obj := i.on.Eval(s)
	hm, ok := obj.(pluggable.HasMethods)
	if !ok {
		panic("we need to think about this since it's a <<runtime>> error")
	}
	meth := hm.ObtainMethod(i.call.Id())
	return meth.Invoke(s, i.on, i.args)
}

type InvokeFunc struct {
	tools *pluggable.Tools
}

func (i *InvokeFunc) Eval(me pluggable.Token, before []pluggable.Expr, after []pluggable.Expr) pluggable.Expr {
	if len(before) != 1 {
		panic("should be an error")
	}
	if len(after) < 1 {
		panic("should be an error")
	}
	meth, ok := after[0].(pluggable.Var)
	if !ok {
		panic("should be an error")
	}
	return &InvokeExpr{Locatable: me, on: before[0], call: meth.Named(), args: after[1:]}
}

func MakeInvokeFunc(tools *pluggable.Tools) *InvokeFunc {
	return &InvokeFunc{tools: tools}
}
