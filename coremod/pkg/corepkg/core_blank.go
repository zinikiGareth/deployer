package corepkg

import (
	"reflect"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corestrats"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreBlank struct {
	Name   string
	OfType reflect.Type
}

func (b *CoreBlank) Mint(tools *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown corebottom.TearDown) corebottom.Ensurable {
	strat := reflect.New(b.OfType).Interface().(corestrats.CreationStrategy)
	return &CoreCreator{tools: tools, teardown: teardown, loc: loc, coin: id, name: named, props: props, findme: strat.(FindStrategy), strategy: strat}
}

func (b *CoreBlank) Find(tools *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) corebottom.FindCoin {
	strat := reflect.New(b.OfType).Interface().(FindStrategy)
	return &CoreCreator{tools: tools, loc: loc, coin: id, name: named, props: props, findme: strat}
}

func (b *CoreBlank) ShortDescription() string {
	return b.Name
}

var _ corebottom.Blank = &CoreBlank{}
