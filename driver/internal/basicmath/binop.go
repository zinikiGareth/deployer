package basicmath

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type binop struct {
	driverbottom.Locatable
	opname string
	lhs    driverbottom.Expr
	rhs    driverbottom.Expr
}

func (m *binop) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (m *binop) ShortDescription() string {
	return fmt.Sprintf("%s [%s,%s]", m.opname, m.lhs.ShortDescription(), m.rhs.ShortDescription())
}

func (m *binop) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (m *binop) String() string {
	panic("unimplemented")
}

func (m *binop) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	return driverbottom.MAY_BE_BOUND
}
