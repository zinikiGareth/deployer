package target

import (
	"log"
	"slices"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreTarget struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	name  driverbottom.SymbolName

	actions []driverbottom.ModelBuilder
}

func (cc *CoreTarget) Name() driverbottom.SymbolName {
	return cc.name
}

func (cc *CoreTarget) String() string {
	return string(cc.name)
}

func (a *CoreTarget) MakeAssign(holder driverbottom.Describable, assignTo driverbottom.Identifier, action driverbottom.ModelBuilder) any {
	ret := MakeDoAssign(a.tools, holder, assignTo, action)
	return ret
}

func (cc *CoreTarget) Attach(entry any) {
	action, ok := entry.(driverbottom.ModelBuilder)
	if !ok {
		log.Fatalf("not a ModelBuilder: %T", entry)
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

func (t *CoreTarget) Resolve(r driverbottom.Resolver) {
	for _, a := range t.actions {
		binding := a.Resolve(r)
		if binding == driverbottom.MUST_BE_BOUND {
			panic("assignTo is not specified") // should be an error
		} else if binding == driverbottom.ERROR_OCCURRED {
			panic("an error occurred")
		}
	}
}

func (t *CoreTarget) BuildModel() {
	for _, a := range t.actions {
		a.BuildModel(t)
	}
}

func (t *CoreTarget) UpdateReality() {
	for _, a := range t.actions {
		amis, ok := a.(driverbottom.RealityShifter)
		if ok {
			amis.UpdateReality()
		}
	}
}

func (t *CoreTarget) TearDown() {
	for _, a := range slices.Backward(t.actions) {
		amis, ok := a.(driverbottom.RealityShifter)
		if ok {
			amis.TearDown()
		}
	}
}

func (d *CoreTarget) Present(value any) {
	// If I have understood the flow correctly, if you arrive here,
	// it must be the case that binding is optional and no variable has been provided.
}

var _ driverbottom.TopLevelForm = &CoreTarget{}
var _ driverbottom.AttachResult = &CoreTarget{}
