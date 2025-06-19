package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/pluggable"
)

type Apply struct {
	pluggable.Locatable
	Func pluggable.Function
	Args []pluggable.Expr
}

func (a *Apply) Resolve(r pluggable.Resolver) {
}

func (a Apply) Eval(s pluggable.RuntimeStorage) any {
	panic("not implemented")
}

func (t *Apply) ShortDescription() string {
	panic("not implemented")
}

func (a Apply) DumpTo(iw pluggable.IndentWriter) {
	panic("not implemented")
}

func (a Apply) String() string {
	return fmt.Sprintf("Apply[%d]", len(a.Args))
}
