package basic

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type GetCoinMethod struct {
	loc  *errorsink.Location
	coin corebottom.CoinId
}

func (gcm *GetCoinMethod) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (gcm *GetCoinMethod) ShortDescription() string {
	panic("unimplemented")
}

func (gcm *GetCoinMethod) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (gcm *GetCoinMethod) String() string {
	panic("unimplemented")
}

func (gcm *GetCoinMethod) Resolve(r driverbottom.Resolver) {
	panic("unimplemented")
}

func (gcm *GetCoinMethod) Eval(s driverbottom.RuntimeStorage) any {
	ret := s.GetCoin(gcm.coin, s.CurrentMode())
	if ret != nil {
		return ret
	}
	if s.CurrentMode() > 2 {
		ret := s.GetCoin(gcm.coin, 2)
		if ret != nil {
			return ret
		}
	}
	if s.CurrentMode() > 0 {
		ret := s.GetCoin(gcm.coin, 0)
		if ret != nil {
			return ret
		}
	}
	log.Fatalf("could not recover coin %s in mode %d", gcm.coin.VarName(), s.CurrentMode())
	return nil
}

func MakeGetCoinMethod(loc *errorsink.Location, coin corebottom.CoinId) driverbottom.Expr {
	return &GetCoinMethod{loc: loc, coin: coin}
}
