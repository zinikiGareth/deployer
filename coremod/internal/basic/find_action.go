package basic

import (
	"fmt"
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type FindAction struct {
	tools    *corebottom.Tools
	scope    driverbottom.Scope
	loc      *errorsink.Location
	what     driverbottom.Identifier
	resolved corebottom.Blank
	named    driverbottom.String
	props    map[driverbottom.Identifier]driverbottom.Expr
	coinId   corebottom.CoinId
	coin     corebottom.FindCoin
}

func (fa *FindAction) Loc() *errorsink.Location {
	return fa.loc
}

func (fa *FindAction) What() driverbottom.SymbolType {
	return driverbottom.SymbolType(fa.what.Id())
}

func (fa *FindAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("FindAction")
	w.AttrsWhere(fa)
	w.TextAttr("what", fa.what.Id())
	if fa.resolved == nil {
		w.TextAttr("not-resolved", fa.what.Id())
	} else {
		w.TextAttr("resolved", fa.resolved.ShortDescription())
	}
	if fa.named != nil {
		w.TextAttr("named", fa.named.Text())
	}
	if len(fa.props) > 0 {
		w.Indent()
		for k, v := range fa.props {
			w.IndPrintf("%s <- %s\n", k, v.String())
		}
		w.UnIndent()
	}
	w.EndAttrs()
}

func (fa *FindAction) ShortDescription() string {
	return fmt.Sprintf("Find[%s: %s]", fa.what.Id(), fa.named.Text())
}

func (fa *FindAction) AddProperty(name driverbottom.Identifier, value driverbottom.Expr) {
	if name.Id() == "name" {
		if fa.named != nil {
			fa.tools.Reporter.Report(name.Loc().Offset, "duplicate definition of name")
			return
		}
		str, ok := value.(driverbottom.String)
		if !ok {
			fa.tools.Reporter.Report(value.Loc().Offset, "name must be a string")
			return
		}
		fa.named = str
	} else {
		if fa.props[name] != nil {
			fa.tools.Reporter.Reportf(name.Loc().Offset, "duplicate definition of %s", name.Id())
			return
		}
		fa.props[name] = value
	}
}

func (fa *FindAction) AddAdverb(adv driverbottom.Adverb, tokens []driverbottom.Token) driverbottom.Interpreter {
	fa.tools.Reporter.Reportf(adv.Loc().Offset, "find cannot handle %s\n", adv.Name())
	return drivertop.NewDisallowInnerScope(fa.tools.CoreTools)
}

func (fa *FindAction) Completed() {
	if fa.tools.Reporter.HasErrors() {
		return
	}
	if fa.named == nil {
		fa.tools.Reporter.ReportAtf(fa.loc, "Find requires a name to be defined")
	}
}

func (fa *FindAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	tmp := r.Resolve(fa.scope, fa.what)
	res, ok := tmp.(corebottom.Blank)
	if !ok {
		log.Printf("could not make %T a Blank", tmp)
		return driverbottom.ERROR_OCCURRED
	}
	fa.resolved = res
	fa.coinId = corebottom.CoinId(fa.tools.Storage.NewObjId(fa.named.Loc()))
	fa.coin = fa.resolved.Find(fa.tools, fa.Loc(), fa.coinId, fa.named.Text(), fa.props)
	return driverbottom.MUST_BE_BOUND
}

func (fa *FindAction) CoinId() corebottom.CoinId {
	return fa.coinId
}

func (fa *FindAction) DetermineInitialState(pres corebottom.ValuePresenter) {
	fa.coin.DetermineInitialState(&findPres{tools: fa.tools, named: fa.named, parent: pres})
}

type findPres struct {
	tools  *corebottom.Tools
	named  driverbottom.String
	parent corebottom.ValuePresenter
}

func (f *findPres) NotFound() {
	log.Printf("could not find %s\n", f.named.Text())
}

func (f *findPres) Present(value any) {
	f.parent.Present(value)
}

func (f *findPres) WantDestruction(loc *errorsink.Location) {
	panic("surely that's an error in find?")
}

var _ corebottom.ValuePresenter = &findPres{}
var _ corebottom.Findable = &FindAction{}
var _ corebottom.CoinProvider = &FindAction{}
