package corepkg

import (
	"reflect"

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

	strategy CreationStrategy
}

type CreationStrategy interface {
	DetermineInitialState(creator *CoreCreator, pres corebottom.ValuePresenter)
	DetermineDesiredState(creator *CoreCreator, pres corebottom.ValuePresenter)
	UpdateReality(creator *CoreCreator, initial any, desired any)
	TearDown(creator *CoreCreator, initial any)
}

func (creator *CoreCreator) GetEnv(driver string, ofType reflect.Type, meth string, field string) {
	ae := creator.Recall().ObtainDriver(driver)
	if ae == nil {
		panic("did not find driver " + driver)
	}
	if !reflect.TypeOf(ae).AssignableTo(ofType) {
		panic("value of type " + reflect.TypeOf(ae).String() + " not assignable to " + ofType.String())
	}
	if m, ok := ofType.MethodByName(meth); ok {
		cli := m.Func.Call([]reflect.Value{reflect.ValueOf(ae)})[0]
		reflect.ValueOf(creator.strategy).Elem().FieldByName(field).Set(cli)
	} else {
		panic("there is no method " + meth + " in " + ofType.String())
	}
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

func (c *CoreCreator) Name() string {
	return c.name
}

func (c *CoreCreator) CoinId() corebottom.CoinId {
	return c.coin
}

func (c *CoreCreator) DetermineInitialState(pres corebottom.ValuePresenter) {
	c.strategy.DetermineInitialState(c, pres)
}

func (c *CoreCreator) DetermineDesiredState(pres corebottom.ValuePresenter) {
	c.strategy.DetermineDesiredState(c, pres)
}

func (c *CoreCreator) UpdateReality() {
	initial := c.tools.Storage.GetCoin(c.coin, corebottom.DETERMINE_INITIAL_MODE)
	desired := c.tools.Storage.GetCoin(c.coin, corebottom.DETERMINE_DESIRED_MODE)

	c.strategy.UpdateReality(c, initial, desired)
}

func (c *CoreCreator) TearDown() {
	panic("unimplemented")
}

func (c *CoreCreator) Recall() driverbottom.Recall {
	return c.tools.Recall
}

var _ corebottom.Ensurable = &CoreCreator{}
var _ corebottom.FindCoin = &CoreCreator{}
