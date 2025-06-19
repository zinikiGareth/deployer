package corebottom

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type Blank interface {
	ShortDescription() string
	Find(tools *Tools, loc *errorsink.Location, named string) any
	Mint(tools *Tools, loc *errorsink.Location, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown TearDown) any
}
