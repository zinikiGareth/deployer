package target

import (
	"fmt"
	"log"
	"slices"

	"ziniki.org/deployer/coremod/internal/vars"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreTarget struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	name  driverbottom.SymbolName

	actions []corebottom.Action
}

func (cc *CoreTarget) Name() driverbottom.SymbolName {
	return cc.name
}

func (cc *CoreTarget) String() string {
	return string(cc.name)
}

func (a *CoreTarget) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	ret := vars.MakeDoAssign(a.tools, holder, assignTo, action)
	return ret
}

func (cc *CoreTarget) Attach(entry any) {
	// log.Printf("%p: attaching %p\n", cc, entry)
	action, ok := entry.(corebottom.Action)
	if !ok {
		log.Fatalf("not an Action: %T", entry)
	}
	cc.actions = append(cc.actions, action)
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
			fnd.DetermineInitialState(t)
		}
	}
}

func (t *CoreTarget) DetermineDesiredState() {
	for k, a := range t.actions {
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		// log.Printf("%p: a is of type %T %p\n", t, a, a)
		mb, ok := a.(corebottom.ModelBuilder)
		if ok {
			mb.DetermineDesiredState(t)
		}
		_, ok = a.(corebottom.MemoryBuilder)
		if ok {
			panic("should have a var attached")
		}
	}
}

func (t *CoreTarget) UpdateReality() {
	for k, a := range t.actions {
		t.tools.Storage.SetStepName(fmt.Sprintf("%s-%d", t.name, k))
		amis, ok := a.(corebottom.RealityShifter)
		if ok {
			amis.UpdateReality()
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

func (d *CoreTarget) NotFound() {
	// If I have understood the flow correctly, if you arrive here,
	// it must be the case that binding is optional and no variable has been provided.
}

func (d *CoreTarget) Present(value any) {
	// If I have understood the flow correctly, if you arrive here,
	// it must be the case that binding is optional and no variable has been provided.
}

var _ corebottom.Target = &CoreTarget{}
var _ driverbottom.TopLevelForm = &CoreTarget{}
var _ driverbottom.AttachResult = &CoreTarget{}
