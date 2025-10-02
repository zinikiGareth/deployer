package exprs

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type anyExpr struct {
	loc   *errorsink.Location
	value any
}

func (a *anyExpr) Loc() *errorsink.Location {
	return a.loc
}

func (a *anyExpr) ShortDescription() string {
	return "AnyExpr[]"
}

func (a *anyExpr) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("AnyExpr")
	to.AttrsWhere(a)
	to.EndAttrs()
}

func (a *anyExpr) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND
	return ret
}

func (a *anyExpr) Eval(s driverbottom.RuntimeStorage) any {
	return a.value
}

func (a *anyExpr) String() string {
	return "AnyExpr[]"
}

func NewAnyExpr(loc *errorsink.Location, value any) driverbottom.Expr {
	return &anyExpr{loc: loc, value: value}
}

var _ driverbottom.Expr = &anyExpr{}
