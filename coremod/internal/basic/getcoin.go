package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type GetCoinMethod struct {
	loc  *errorsink.Location
	coin corebottom.CoinId
}

func (gcm *GetCoinMethod) Loc() *errorsink.Location {
	return gcm.loc
}

func (gcm *GetCoinMethod) ShortDescription() string {
	return fmt.Sprintf("GetCoin[%s]", gcm.coin.VarName())
}

func (gcm *GetCoinMethod) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("GetCoin")
	to.AttrsWhere(gcm)
	to.TextAttr("coin", gcm.coin.VarName().Id())
	to.EndAttrs()
}

func (gcm *GetCoinMethod) String() string {
	return fmt.Sprintf("GetCoin[%s]", gcm.coin.VarName())
}

func (gcm *GetCoinMethod) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	return driverbottom.MAY_BE_BOUND
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
	// s.ExportSymbolsTo(utils.NewIndentWriter(os.Stderr))
	// log.Printf("could not recover coin %s in mode %d\n", gcm.coin.VarName(), s.CurrentMode())
	return drivertop.NewAlwaysNil(gcm.loc)
}

func MakeGetCoinMethod(loc *errorsink.Location, coin corebottom.CoinId) driverbottom.Expr {
	return &GetCoinMethod{loc: loc, coin: coin}
}
