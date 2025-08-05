package exprs

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type AlwaysNil struct {
	loc *errorsink.Location
}

func (a *AlwaysNil) Loc() *errorsink.Location {
	return a.loc
}

func (a *AlwaysNil) ShortDescription() string {
	return "AlwaysNil[]"
}

func (a *AlwaysNil) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("AlwaysNil")
	to.AttrsWhere(a)
	to.EndAttrs()
}

func (a *AlwaysNil) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND
	return ret
}

func (a *AlwaysNil) Eval(s driverbottom.RuntimeStorage) any {
	return nil
}

func (a *AlwaysNil) String() string {
	return ""
}

func (a *AlwaysNil) ObtainMethod(name string) driverbottom.Method {
	return &alwaysNilMethod{loc: a.loc}
}

func NewAlwaysNil(loc *errorsink.Location) driverbottom.Expr {
	return &AlwaysNil{loc: loc}
}

type alwaysNilMethod struct {
	loc *errorsink.Location
}

func (m *alwaysNilMethod) Invoke(storage driverbottom.RuntimeStorage, obj driverbottom.Expr, args []driverbottom.Expr) any {
	return NewAlwaysNil(m.loc)
}

var _ driverbottom.HasMethods = &AlwaysNil{}
