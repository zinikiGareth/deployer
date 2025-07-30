package exprs

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type alwaysNil struct {
	loc *errorsink.Location
}

func (a *alwaysNil) Loc() *errorsink.Location {
	return a.loc
}

func (a *alwaysNil) ShortDescription() string {
	return "AlwaysNil[]"
}

func (a *alwaysNil) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("AlwaysNil")
	to.AttrsWhere(a)
	to.EndAttrs()
}

func (a *alwaysNil) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND
	return ret
}

func (a *alwaysNil) Eval(s driverbottom.RuntimeStorage) any {
	return nil
}

func (a *alwaysNil) String() string {
	return ""
}

func (a *alwaysNil) ObtainMethod(name string) driverbottom.Method {
	return &alwaysNilMethod{loc: a.loc}
}

func NewAlwaysNil(loc *errorsink.Location) driverbottom.Expr {
	return &alwaysNil{loc: loc}
}

type alwaysNilMethod struct {
	loc *errorsink.Location
}

func (m *alwaysNilMethod) Invoke(storage driverbottom.RuntimeStorage, obj driverbottom.Expr, args []driverbottom.Expr) any {
	return NewAlwaysNil(m.loc)
}

var _ driverbottom.HasMethods = &alwaysNil{}
