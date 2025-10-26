package corepkg

import (
	"reflect"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corestrats"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreFinder struct {
	Name   string
	OfType reflect.Type
}

func (b *CoreFinder) Find(tools *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) corebottom.FindCoin {
	strat := reflect.New(b.OfType).Interface().(corestrats.FindStrategy)
	return &CoreCreator{tools: tools, loc: loc, coin: id, name: named, props: props, findme: strat}
}

func (b *CoreFinder) ShortDescription() string {
	return b.Name
}

var _ corebottom.Finder = &CoreFinder{}
