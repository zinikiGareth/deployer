package corepkg

import (
	"reflect"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreBlank struct {
	Name   string
	OfType reflect.Type
}

type CoreCreatorWrapper interface {
	Creator() *CoreCreator
}

func (b *CoreBlank) Mint(tools *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown corebottom.TearDown) corebottom.Ensurable {
	foo := reflect.New(b.OfType).Interface().(CoreCreatorWrapper)
	foo.Creator().tools = tools
	// foo.teardown = teardown
	// foo.loc = loc
	// foo.coin = id
	// foo.name = named
	// foo.props = props
	return foo.Creator()
}

func (b *CoreBlank) Find(tools *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr) corebottom.FindCoin {
	foo := reflect.New(b.OfType).Interface().(*CoreCreator)
	foo.tools = tools
	foo.loc = loc
	foo.name = named
	foo.props = props
	return foo
}

func (b *CoreBlank) ShortDescription() string {
	return b.Name
}

var _ corebottom.Blank = &CoreBlank{}
