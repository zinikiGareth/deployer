package exprs

import (
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
	panic("unimplemented")
}

func (l *ListExpr) Eval(s driverbottom.RuntimeStorage) any {
	ret := []any{}
	for _, e := range l.exprs {
		ret = append(ret, s.Eval(e))
	}
	return ret
}

func (l *ListExpr) String() string {
	panic("unimplemented")
}

var _ driverbottom.Expr = &ListExpr{}
