package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type Apply struct {
	driverbottom.Locatable
	Func driverbottom.Function
	Args []driverbottom.Expr
}

func (a *Apply) Resolve(r driverbottom.Resolver) {
}

func (a Apply) Eval(s driverbottom.RuntimeStorage) any {
	panic("not implemented")
}

func (t *Apply) ShortDescription() string {
	panic("not implemented")
}

func (a Apply) DumpTo(iw driverbottom.IndentWriter) {
	panic("not implemented")
}

func (a Apply) String() string {
	return fmt.Sprintf("Apply[%d]", len(a.Args))
}
