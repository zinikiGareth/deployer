package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type FindAction struct {
	tools    *corebottom.Tools
	loc      *errorsink.Location
	what     driverbottom.Identifier
	resolved corebottom.Blank
	named    driverbottom.String
	props    map[driverbottom.Identifier]driverbottom.Expr
	ens      corebottom.Findable
}

func (ea *FindAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *FindAction) What() driverbottom.SymbolType {
	return driverbottom.SymbolType(ea.what.Id())
}

func (ea *FindAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("FindAction")
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
		w.Indent()
		for k, v := range ea.props {
			w.IndPrintf("%s <- %s\n", k, v.String())
		}
		w.UnIndent()
	}
	w.EndAttrs()
}

func (ea *FindAction) ShortDescription() string {
	return fmt.Sprintf("Find[%s: %s]", ea.what.Id(), ea.named.Text())
}

func (ea *FindAction) AddProperty(name driverbottom.Identifier, value driverbottom.Expr) {
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

func (ea *FindAction) AddAdverb(adv driverbottom.Adverb, tokens []driverbottom.Token) driverbottom.Interpreter {
	ea.tools.Reporter.Reportf(adv.Loc().Offset, "find cannot handle %s\n", adv.Name())
	return drivertop.NewDisallowInnerScope(ea.tools.CoreTools)
}

func (ea *FindAction) Completed() {
	if ea.tools.Reporter.HasErrors() {
		return
	}
	if ea.named == nil {
		ea.tools.Reporter.At(ea.loc.Line)
		ea.tools.Reporter.Report(ea.loc.Offset, "Find requires a name to be defined")
	}
}

func (ea *FindAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	res, ok := r.Resolve(ea.what).(corebottom.Blank)
	if !ok {
		return driverbottom.ERROR_OCCURRED
	}
	ea.resolved = res
	obj := ea.resolved.Find(ea.tools, ea.Loc(), ea.named.Text())
	ens, ok := obj.(corebottom.Findable)
	if !ok {
		ea.tools.Storage.Errorf(ea.loc, "the type "+ea.what.Id()+" is not findable")
		return driverbottom.ERROR_OCCURRED
	}
	ea.ens = ens
	return driverbottom.MAY_BE_BOUND
}

func (ea *FindAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	ea.ens.DetermineDesiredState(pres)
}

func (ea *FindAction) UpdateReality() {
}

func (ea *FindAction) TearDown() {
}
