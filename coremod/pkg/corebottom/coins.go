package corebottom

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoinId driverbottom.ResolvableHolder

type Blank interface {
	ShortDescription() string
	Find(tools *Tools, loc *errorsink.Location, id CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) FindCoin
	Mint(tools *Tools, loc *errorsink.Location, id CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown TearDown) Ensurable
}

type MemoryCoin interface {
	ShortDescription() string
	Mint(tools *Tools, loc *errorsink.Location, id CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) MemoryCoinCreator
}

type CoinProvider interface {
	CoinId() CoinId
}
