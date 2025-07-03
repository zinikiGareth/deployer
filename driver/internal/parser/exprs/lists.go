package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type ListExpr struct {
	loc   *errorsink.Location
	exprs []driverbottom.Expr
}

func (l *ListExpr) Loc() *errorsink.Location {
	return l.loc
}

func (l *ListExpr) ShortDescription() string {
	return fmt.Sprintf("List[%d]", len(l.exprs))
}

func (l *ListExpr) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("List")
	to.AttrsWhere(l)
	for k, e := range l.exprs {
		to.NestedAttr(fmt.Sprintf("entry %d", k), e)
	}

}

func (l *ListExpr) Resolve(r driverbottom.Resolver) {
	for _, e := range l.exprs {
		e.Resolve(r)
	}
}

func (l *ListExpr) Eval(s driverbottom.RuntimeStorage) any {
	ret := []any{}
	for _, e := range l.exprs {
		ret = append(ret, s.Eval(e))
	}
	return ret
}

func (l *ListExpr) IsEmpty() bool {
	return len(l.exprs) == 0
}

func (l *ListExpr) Length() int {
	return len(l.exprs)
}

func (l *ListExpr) String() string {
	return fmt.Sprintf("[<%d>]", len(l.exprs))
}

func (l *ListExpr) Members() []driverbottom.Expr {
	return l.exprs
}

func NewListExpr(loc *errorsink.Location, exprs []driverbottom.Expr) driverbottom.List {
	return &ListExpr{loc: loc, exprs: exprs}
}

var _ driverbottom.List = &ListExpr{}
