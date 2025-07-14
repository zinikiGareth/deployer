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

type EnsureAction struct {
	tools       *corebottom.Tools
	scope       driverbottom.Scope
	loc         *errorsink.Location
	what        driverbottom.Identifier
	resolved    corebottom.Blank
	named       driverbottom.String
	teardown    corebottom.TearDown
	markDestroy driverbottom.Adverb
	// TODO: this really isn't a map here, because what we want is to index by string name
	// Now, we possibly want the "identifier" to that we have the location of the symbol, but it can't be the key
	// because the same "symbol" at two different locations will be different, and we can't index by it.
	props map[driverbottom.Identifier]driverbottom.Expr
	ens   corebottom.Ensurable
}

func (ea *EnsureAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *EnsureAction) What() driverbottom.SymbolType {
	return driverbottom.SymbolType(ea.what.Id())
}

func (ea *EnsureAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("EnsureAction")
	w.AttrsWhere(ea)
	w.TextAttr("what", ea.what.Id())
	if ea.resolved == nil {
		w.TextAttr("not-resolved", ea.what.Id())
	} else {
		w.TextAttr("resolved", ea.resolved.ShortDescription())
	}
	if ea.named != nil {
		w.TextAttr("named", ea.named.Text())
	}
	if len(ea.props) > 0 {
		w.IndPrintf("additional properties:\n")
		w.Indent()
		keys := []string{}
		ids := map[string]driverbottom.Expr{}
		for id, val := range ea.props {
			keys = append(keys, id.Id())
			ids[id.Id()] = val
		}
		slices.Sort(keys)
		for _, k := range keys {
			w.NestedAttr(k, ids[k])
		}
		w.UnIndent()
	}
	w.EndAttrs()
}

func (ea *EnsureAction) ShortDescription() string {
	return fmt.Sprintf("Ensure[%s: %s]", ea.what.Id(), ea.named.Text())
}

func (ea *EnsureAction) AddProperty(name driverbottom.Identifier, value driverbottom.Expr) {
	if name.Id() == "name" {
		if ea.named != nil {
			ea.tools.Reporter.Report(name.Loc().Offset, "duplicate definition of name")
			return
		}
		str, ok := value.(driverbottom.String)
		if !ok {
			ea.tools.Reporter.Report(value.Loc().Offset, "name must be a string")
			return
		}
		ea.named = str
	} else {
		if ea.props[name] != nil {
			ea.tools.Reporter.Reportf(name.Loc().Offset, "duplicate definition of %s", name.Id())
			return
		}
		ea.props[name] = value
	}
}

func (ea *EnsureAction) AddAdverb(adv driverbottom.Adverb, tokens []driverbottom.Token) driverbottom.Interpreter {
	switch adv.Name() {
	case "teardown":
		if ea.teardown != nil {
			panic("duplicate teardown")
		}
		if len(tokens) != 1 {
			panic("invalid tokens")
		}
		ea.teardown = &MyTearDown{mode: tokens[0].(driverbottom.Identifier).Id()}
	case "destroy":
		ea.markDestroy = adv
	default:
		ea.tools.Reporter.ReportAtf(adv.Loc(), "there is no adverb %s", adv.Name())
	}
	return drivertop.NewDisallowInnerScope(ea.tools.CoreTools)
}

func (ea *EnsureAction) Completed() {
	if ea.tools.Reporter.HasErrors() {
		return
	}
	if ea.named == nil {
		ea.tools.Reporter.ReportAtf(ea.loc, "ensure requires a name to be defined")
	}
	if ea.teardown == nil {
		ea.tools.Reporter.ReportAtf(ea.loc, "ensure requires a teardown strategy to be declared")
	}
}

func (ea *EnsureAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	tmp := r.Resolve(ea.scope, ea.what)
	res, ok := tmp.(corebottom.Blank)
	if !ok {
		log.Printf("could not make %T a Blank", tmp)
		return driverbottom.ERROR_OCCURRED
	}
	for _, y := range ea.props {
		y.Resolve(r)
	}
	ea.resolved = res
	ea.ens = ea.resolved.Mint(ea.tools, ea.Loc(), corebottom.CoinId(ea.tools.Storage.NewObjId(ea.named.Loc())), ea.named.Text(), ea.props, ea.teardown)
	return driverbottom.MAY_BE_BOUND
}

func (ea *EnsureAction) DetermineInitialState(pres corebottom.ValuePresenter) {
	ea.ens.DetermineInitialState(NewCoinPresenter(ea.tools.Storage, ea.ens.CoinId(), pres))
	if ea.markDestroy != nil {
		pres.WantDestruction(ea.markDestroy.Loc())
	}
}

func (ea *EnsureAction) DetermineDesiredState(pres corebottom.ValuePresenter) {
	ea.ens.DetermineDesiredState(NewCoinPresenter(ea.tools.Storage, ea.ens.CoinId(), pres))
}

func (ea *EnsureAction) UpdateReality() {
	ea.ens.UpdateReality()
}

func (ea *EnsureAction) TearDown() {
	ea.ens.TearDown()
}

type MyTearDown struct {
	mode string
}

func (m *MyTearDown) Mode() string {
	return m.mode
}

var _ corebottom.ModelBuilder = &EnsureAction{}
