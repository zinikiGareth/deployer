package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type MapExpr struct {
	exprs []driverbottom.Expr
}

func (l *MapExpr) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (l *MapExpr) ShortDescription() string {
	panic("unimplemented")
}

func (l *MapExpr) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (l *MapExpr) Resolve(r driverbottom.Resolver) {
	for _, e := range l.exprs {
		e.Resolve(r)
	}
}

func (l *MapExpr) Eval(s driverbottom.RuntimeStorage) any {
	ret := []any{}
	for _, e := range l.exprs {
		ret = append(ret, s.Eval(e))
	}
	return ret
}

func (l *MapExpr) IsEmpty() bool {
	return len(l.exprs) == 0
}

func (l *MapExpr) Length() int {
	return len(l.exprs)
}

func (l *MapExpr) String() string {
	return fmt.Sprintf("[<%d>]", len(l.exprs))
}

func (l *MapExpr) Members() []driverbottom.Expr {
	return l.exprs
}

func NewMapExpr(exprs []driverbottom.Expr) driverbottom.List {
	return &MapExpr{exprs: exprs}
}

var _ driverbottom.List = &MapExpr{}
