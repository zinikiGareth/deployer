package basic

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type CoinPresenter struct {
	storage   driverbottom.RuntimeStorage
	coinId    corebottom.CoinId
	presenter corebottom.ValuePresenter
}

func (c *CoinPresenter) Present(value any) {
	log.Printf("presenting value %p\n", value)
	c.presenter.Present(value)
	if c.coinId != nil {
		log.Printf("ignoring value %p\n", value)
		c.storage.IgnoreDuplicate(value)
		log.Printf("binding value %p for %s\n", value, c.coinId.VarName().Id())
		c.storage.Bind(c.coinId, value)
	}
}

func (c *CoinPresenter) NotFound() {
	c.presenter.NotFound()
}

func NewCoinPresenter(storage driverbottom.RuntimeStorage, coinId corebottom.CoinId, pres corebottom.ValuePresenter) corebottom.ValuePresenter {
	return &CoinPresenter{storage: storage, coinId: coinId, presenter: pres}
}
