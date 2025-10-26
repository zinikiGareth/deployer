package corepkg

import (
	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corestrats"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type Composite struct {
	tools *corebottom.Tools
	loc   *errorsink.Location

	named    driverbottom.String
	creators []corebottom.BasicShifter

	strat corestrats.CompositeStrategy
}

func (a *Composite) Loc() *errorsink.Location {
	return a.loc
}

func (a *Composite) ShortDescription() string {
	panic("unimplemented")
}

func (a *Composite) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (a *Composite) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	panic("unimplemented")
}

func (a *Composite) DetermineInitialState(pres corebottom.ValuePresenter) {
	panic("unimplemented")
}

func (a *Composite) DetermineDesiredState(pres corebottom.ValuePresenter) {
	dummy := basic.NewDummyPresenter()
	for _, c := range a.creators {
		var mypres corebottom.ValuePresenter
		cp, ok := c.(corebottom.CoinProvider)
		if ok {
			mypres = basic.NewCoinPresenter(a.tools.Storage, cp.CoinId(), dummy)
		} else {
			mypres = dummy
		}
		c.DetermineDesiredState(mypres)
	}
	// TODO: we probably should have some kind of composite model we present
}

func (a *Composite) UpdateReality() {
	panic("unimplemented")
}

func (a *Composite) TearDown() {
	panic("unimplemented")
}

func (a *Composite) ShouldDestroy() bool {
	panic("unimplemented")
}

var _ corebottom.RealityShifter = &Composite{}
