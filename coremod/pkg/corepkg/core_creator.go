package corepkg

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreCreator struct {
	tools *corebottom.Tools

	loc      *errorsink.Location
	name     string
	coin     corebottom.CoinId
	teardown corebottom.TearDown
	props    map[driverbottom.Identifier]driverbottom.Expr
}

func (c *CoreCreator) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (c *CoreCreator) ShortDescription() string {
	panic("unimplemented")
}

func (c *CoreCreator) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (c *CoreCreator) CoinId() corebottom.CoinId {
	panic("unimplemented")
}

func (c *CoreCreator) DetermineInitialState(pres corebottom.ValuePresenter) {
	panic("unimplemented")
}

func (c *CoreCreator) DetermineDesiredState(pres corebottom.ValuePresenter) {
	panic("unimplemented")
}

func (c *CoreCreator) UpdateReality() {
	panic("unimplemented")
}

func (c *CoreCreator) TearDown() {
	panic("unimplemented")
}

var _ corebottom.Ensurable = &CoreCreator{}
var _ corebottom.FindCoin = &CoreCreator{}
