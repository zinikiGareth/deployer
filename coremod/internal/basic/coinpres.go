package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoinPresenter struct {
	storage   driverbottom.RuntimeStorage
	coinId    corebottom.CoinId
	presenter corebottom.ValuePresenter
}

func (c *CoinPresenter) Present(value any) {
	if c.presenter != nil {
		c.presenter.Present(value)
	}
	if c.coinId != nil {
		c.storage.IgnoreDuplicate(value)
		c.storage.Bind(c.coinId, value)
	}
}

func (c *CoinPresenter) NotFound() {
	if c.presenter != nil {
		c.presenter.NotFound()
	}
}

func (c *CoinPresenter) WantDestruction(loc *errorsink.Location) {
	panic("need to handle website.@destroy")
}

func NewCoinPresenter(storage driverbottom.RuntimeStorage, coinId corebottom.CoinId, pres corebottom.ValuePresenter) corebottom.ValuePresenter {
	return &CoinPresenter{storage: storage, coinId: coinId, presenter: pres}
}
