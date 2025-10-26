package target

import (
	"fmt"
	"log"
	"slices"

	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/vars"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreTarget struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	name  driverbottom.SymbolName
	scope driverbottom.Scope

	actions     []corebottom.Action
	haveDestroy bool
}

func (t *CoreTarget) Name() driverbottom.SymbolName {
	return t.name
}

func (t *CoreTarget) Scope() driverbottom.Scope {
	return t.scope
}

func (t *CoreTarget) String() string {
	return string(t.name)
}

func (t *CoreTarget) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	ret := vars.MakeDoAssign(t.tools, holder, assignTo, action)
	return ret
}

func (t *CoreTarget) Attach(entry any) error {
	// log.Printf("%p: attaching %p\n", cc, entry)
	action, ok := entry.(corebottom.Action)
	if !ok {
		return fmt.Errorf("not an Action: %T", entry)
	}
	t.actions = append(t.actions, action)
	return nil
}

func (t *CoreTarget) Loc() *errorsink.Location {
	return t.loc
}

func (t *CoreTarget) ShortDescription() string {
	return "Target[" + string(t.name) + "]"
}

func (t *CoreTarget) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("target %s", t.name)
	w.AttrsWhere(t)
	w.ListAttr("actions")
	for _, a := range t.actions {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (t *CoreTarget) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for k, a := range t.actions {
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		binding := a.Resolve(r)
		switch binding {
		case driverbottom.MUST_BE_BOUND:
			t.tools.Reporter.ReportAtf(a.Loc(), "a value expression must be bound to a variable")
			return driverbottom.ERROR_OCCURRED
		case driverbottom.ERROR_OCCURRED:
			return driverbottom.ERROR_OCCURRED
		}
	}
	return driverbottom.NO_VALUE
}

func (t *CoreTarget) DetermineInitialState() {
	for k, a := range t.actions {
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		fnd, ok := a.(corebottom.Findable)
		if ok {
			fnd.DetermineInitialState(t.presenter(a))
		}
	}
}

func (t *CoreTarget) DetermineDesiredState() {
	if t.tools.Options.Destroy && !t.haveDestroy {
		t.tools.Reporter.ReportAtf(t.loc, "no @destroy elements specified in active target (but --destroy flag specified)")
		return
	}
	for k, a := range t.actions {
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		// log.Printf("%p: a is of type %T %p\n", t, a, a)
		mb, ok := a.(corebottom.ModelBuilder)
		if ok {
			mb.DetermineDesiredState(t.presenter(a))
		}
		_, ok = a.(corebottom.MemoryBuilder)
		if ok {
			panic("should have a var attached")
		}
	}
}

func (t *CoreTarget) UpdateReality() {
	for k, a := range t.actions {
		if t.tools.Reporter.HasErrors() {
			break
		}
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		amis, ok := a.(corebottom.RealityShifter)
		if ok {
			if t.tools.Options.Destroy && amis.ShouldDestroy() {
				amis.TearDown()
			} else {
				amis.UpdateReality()
			}
		} else {
			fndr, ok := a.(corebottom.Findable)
			if ok {
				isCoin, ok := a.(corebottom.CoinProvider)
				if ok {
					// if we didn't find the value first time around, try again
					if !t.tools.Storage.HasCoin(isCoin.CoinId(), 1) {
						fndr.DetermineInitialState(t.presenter(a))
					}
				}
			} else {
				log.Printf("not a RealityShifter or Findable but %T", a)
			}
		}
	}
}

func (t *CoreTarget) TearDown() {
	for k, a := range slices.Backward(t.actions) {
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		amis, ok := a.(corebottom.RealityShifter)
		if ok {
			amis.TearDown()
		}
	}
}

func (t *CoreTarget) presenter(a driverbottom.Describable) corebottom.ValuePresenter {
	var pres corebottom.ValuePresenter = t
	da, ok := a.(*vars.DoAssign)
	if ok {
		a = da.Nested()
	}
	coinProvider, ok := a.(corebottom.CoinProvider)
	if ok {
		pres = basic.NewCoinPresenter(t.tools.Storage, coinProvider.CoinId(), t)
	}
	return pres
}

func (t *CoreTarget) NotFound() {
	// If I have understood the flow correctly, if you arrive here,
	// it must be the case that binding is optional and no variable has been provided.
}

func (t *CoreTarget) Present(value any) {
	// If I have understood the flow correctly, if you arrive here,
	// it must be the case that binding is optional and no variable has been provided.
}

func (t *CoreTarget) WantDestruction(loc *errorsink.Location) {
	if !t.tools.Options.Destroy {
		t.tools.Reporter.ReportAtf(loc, "@destroy specified in active target without the --destroy flag")
	} else {
		t.haveDestroy = true
	}
}

var _ corebottom.Target = &CoreTarget{}
var _ driverbottom.TopLevelForm = &CoreTarget{}
var _ driverbottom.AttachResult = &CoreTarget{}
