package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type ListExpr struct {
	exprs []driverbottom.Expr
}

func (l *ListExpr) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (l *ListExpr) ShortDescription() string {
	panic("unimplemented")
}

func (l *ListExpr) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
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

func NewListExpr(exprs []driverbottom.Expr) *ListExpr {
	return &ListExpr{exprs: exprs}
}

var _ driverbottom.Expr = &ListExpr{}
