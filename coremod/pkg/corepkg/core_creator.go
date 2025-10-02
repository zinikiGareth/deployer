package corepkg

import (
	"reflect"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
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
	TearDown(creator *CoreCreator, initial any, teardown corebottom.TearDown)
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
	initial := c.tools.Storage.GetCoin(c.coin, corebottom.DETERMINE_INITIAL_MODE)
	c.strategy.TearDown(c, initial, c.teardown)
}

func (c *CoreCreator) Recall() driverbottom.Recall {
	return c.tools.Recall
}

func (c *CoreCreator) Adopt(item any) {
	c.tools.Storage.Adopt(c.coin, item)
}

func (c *CoreCreator) Created(item any) {
	c.tools.Storage.Bind(c.coin, item)
}

func (c *CoreCreator) DeferredMethod(name string) driverbottom.Method {
	return &deferredMethod{core: c, name: name}
}

type deferredMethod struct {
	core *CoreCreator
	name string
}

// Invoke implements driverbottom.Method.
func (d *deferredMethod) Invoke(storage driverbottom.RuntimeStorage, obj driverbottom.Expr, args []driverbottom.Expr) any {
	return utils.DeferString(func() string {
		curr := storage.GetCoinFrom(d.core.coin, []int{1, 3}).(driverbottom.HasMethods)
		return curr.ObtainMethod(d.name).Invoke(storage, drivertop.NewAnyExpr(obj.Loc(), curr), args).(string)
	})
}

func SimpleMethod(f func(storage driverbottom.RuntimeStorage, obj driverbottom.Expr) any) driverbottom.Method {
	return &simpleMethod{f: f}
}

type simpleMethod struct {
	f func(storage driverbottom.RuntimeStorage, obj driverbottom.Expr) any
}

func (s *simpleMethod) Invoke(storage driverbottom.RuntimeStorage, obj driverbottom.Expr, args []driverbottom.Expr) any {
	return s.f(storage, obj)
}

var _ corebottom.Ensurable = &CoreCreator{}
var _ corebottom.FindCoin = &CoreCreator{}
