package corebottom

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoinId driverbottom.Holder

type Blank interface {
	ShortDescription() string
	Find(tools *Tools, loc *errorsink.Location, id CoinId, named string) any
	Mint(tools *Tools, loc *errorsink.Location, id CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown TearDown) any
}

type MemoryCoin interface {
	ShortDescription() string
	Mint(tools *Tools, loc *errorsink.Location, id CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) any
}
