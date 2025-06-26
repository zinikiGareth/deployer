package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type FindCoin interface {
	driverbottom.Describable
	DetermineInitialState(pres driverbottom.ValuePresenter)
}

type Ensurable interface {
	FindCoin
	DetermineDesiredState(pres driverbottom.ValuePresenter)
	UpdateReality()
	TearDown()
}
