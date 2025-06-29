package testS3

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type ConfigurationCoin struct{}

func (b *ConfigurationCoin) Mint(ct *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) any {
	return &configCreator{tools: ct, loc: loc, coin: id, name: named, props: props}
}

func (b *ConfigurationCoin) ShortDescription() string {
	return "test.S3.Configuration[]"
}

var _ corebottom.MemoryCoin = &ConfigurationCoin{}
