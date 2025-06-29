package basic

import (
	"fmt"
	"log"
	"slices"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type MemoryCoinAction struct {
	tools    *corebottom.Tools
	loc      *errorsink.Location
	what     driverbottom.Identifier
	resolved corebottom.MemoryCoin
	named    driverbottom.String
	// TODO: this really isn't a map here, because what we want is to index by string name
	// Now, we possibly want the "identifier" to that we have the location of the symbol, but it can't be the key
	// because the same "symbol" at two different locations will be different, and we can't index by it.
	props map[driverbottom.Identifier]driverbottom.Expr
	coin  corebottom.MemoryCoinCreator
}

func (mca *MemoryCoinAction) Loc() *errorsink.Location {
	return mca.loc
}

func (mca *MemoryCoinAction) What() driverbottom.SymbolType {
	return driverbottom.SymbolType(mca.what.Id())
}

func (mca *MemoryCoinAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("MemoryCoinAction")
	w.AttrsWhere(mca)
	w.TextAttr("what", mca.what.Id())
	if mca.resolved == nil {
		w.TextAttr("not-resolved", mca.what.Id())
	} else {
		w.TextAttr("resolved", mca.resolved.ShortDescription())
	}
	if mca.named != nil {
		w.TextAttr("named", mca.named.Text())
	}
	if len(mca.props) > 0 {
		w.IndPrintf("additional properties:\n")
		w.Indent()
		keys := []string{}
		ids := map[string]string{}
		for id, val := range mca.props {
			keys = append(keys, id.Id())
			ids[id.Id()] = val.String()
		}
		slices.Sort(keys)
		for _, k := range keys {
			w.IndPrintf("%s <- %s\n", k, ids[k])
		}
		w.UnIndent()
	}
	w.EndAttrs()
}

func (mca *MemoryCoinAction) ShortDescription() string {
	return fmt.Sprintf("Ensure[%s: %s]", mca.what.Id(), mca.named.Text())
}

func (mca *MemoryCoinAction) AddProperty(name driverbottom.Identifier, value driverbottom.Expr) {
	if name.Id() == "name" {
		if mca.named != nil {
			mca.tools.Reporter.Report(name.Loc().Offset, "duplicate definition of name")
			return
		}
		str, ok := value.(driverbottom.String)
		if !ok {
			mca.tools.Reporter.Report(value.Loc().Offset, "name must be a string")
			return
		}
		mca.named = str
	} else {
		if mca.props[name] != nil {
			mca.tools.Reporter.Reportf(name.Loc().Offset, "duplicate definition of %s", name.Id())
			return
		}
		mca.props[name] = value
	}
}

func (mca *MemoryCoinAction) AddAdverb(adv driverbottom.Adverb, tokens []driverbottom.Token) driverbottom.Interpreter {
	return drivertop.NewDisallowInnerScope(mca.tools.CoreTools)
}

func (mca *MemoryCoinAction) Completed() {
	if mca.tools.Reporter.HasErrors() {
		return
	}
	if mca.named == nil {
		mca.tools.Reporter.ReportAtf(mca.loc, "ensure requires a name to be defined")
	}
}

func (mca *MemoryCoinAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	tmp := r.Resolve(mca.what)
	if tmp == nil {
		return driverbottom.ERROR_OCCURRED
	}
	res, ok := tmp.(corebottom.MemoryCoin)
	if !ok {
		log.Printf("could not make %T a MemoryCoin", tmp)
		return driverbottom.ERROR_OCCURRED
	}
	for _, y := range mca.props {
		y.Resolve(r)
	}
	mca.resolved = res
	obj := mca.resolved.Mint(mca.tools, mca.Loc(), corebottom.CoinId(mca.tools.Storage.NewObjId(mca.named.Loc())), mca.named.Text(), mca.props)
	coin, ok := obj.(corebottom.MemoryCoinCreator)
	if !ok {
		log.Printf("could not make %T a memorycoincreator", obj)
		mca.tools.Storage.Errorf(mca.loc, "the type "+mca.what.Id()+" is not ensurable")
		return driverbottom.ERROR_OCCURRED
	}
	mca.coin = coin
	return driverbottom.MUST_BE_BOUND
}

func (mca *MemoryCoinAction) Create(pres corebottom.ValuePresenter) {
	mca.coin.Create(NewCoinPresenter(mca.tools.Storage, mca.coin.CoinId(), pres))
}

var _ corebottom.MemoryBuilder = &MemoryCoinAction{}
