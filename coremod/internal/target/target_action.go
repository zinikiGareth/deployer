package target

import (
	"slices"

	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type CoreTarget struct {
	loc  *errorsink.Location
	name pluggable.SymbolName

	actions []pluggable.ModelBuilder
}

func (cc *CoreTarget) Name() pluggable.SymbolName {
	return cc.name
}

func (cc *CoreTarget) String() string {
	return string(cc.name)
}

func (cc *CoreTarget) Attach(entry any) {
	action, ok := entry.(pluggable.ModelBuilder)
	if !ok {
		panic("not an action")
	}
	cc.actions = append(cc.actions, action)
}

func (t *CoreTarget) Loc() *errorsink.Location {
	return t.loc
}

func (t *CoreTarget) ShortDescription() string {
	return "Target[" + string(t.name) + "]"
}

func (t *CoreTarget) DumpTo(w pluggable.IndentWriter) {
	w.Intro("target %s", t.name)
	w.AttrsWhere(t)
	w.ListAttr("actions")
	for _, a := range t.actions {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (t *CoreTarget) Resolve(r pluggable.Resolver) {
	for _, a := range t.actions {
		binding := a.Resolve(r)
		if binding == pluggable.MUST_BE_BOUND {
			panic("assignTo is not specified") // should be an error
		} else if binding == pluggable.ERROR_OCCURRED {
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
		amis, ok := a.(pluggable.RealityShifter)
		if ok {
			amis.UpdateReality()
		}
	}
}

func (t *CoreTarget) TearDown() {
	for _, a := range slices.Backward(t.actions) {
		amis, ok := a.(pluggable.RealityShifter)
		if ok {
			amis.TearDown()
		}
	}
}

func (d *CoreTarget) Present(value any) {
	// If I have understood the flow correctly, if you arrive here,
	// it must be the case that binding is optional and no variable has been provided.
}
