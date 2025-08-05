package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type HasCoin interface {
	driverbottom.Describable
	CoinProvider
}

type FindCoin interface {
	HasCoin
	DetermineInitialState(pres ValuePresenter)
}

type Ensurable interface {
	FindCoin
	DetermineDesiredState(pres ValuePresenter)
	UpdateReality()
	TearDown()
}

type MemoryCoinCreator interface {
	HasCoin
	Create(pres ValuePresenter)
}
