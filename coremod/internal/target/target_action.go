package target

import (
	"slices"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type coreTarget struct {
	loc  *errorsink.Location
	name pluggable.SymbolName

	actions []pluggable.Action
}

func (cc *coreTarget) String() string {
	return string(cc.name)
}

func (cc *coreTarget) Add(entry pluggable.Action) {
	cc.actions = append(cc.actions, entry)
}

func (t *coreTarget) Loc() *errorsink.Location {
	return t.loc
}

func (t *coreTarget) ShortDescription() string {
	return "Target[" + string(t.name) + "]"
}

func (t *coreTarget) DumpTo(w pluggable.IndentWriter) {
	w.Intro("target %s", t.name)
	w.AttrsWhere(t)
	w.ListAttr("actions")
	for _, a := range t.actions {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (t *coreTarget) Resolve(r pluggable.Resolver) {
	for _, a := range t.actions {
		binding := a.Resolve(r)
		if binding == pluggable.MUST_BE_BOUND {
			panic("assignTo is not specified") // should be an error
		}
	}
}

func (t *coreTarget) Prepare() {
	for _, a := range t.actions {
		a.Prepare(t)
	}
}

func (t *coreTarget) Execute() {
	for _, a := range t.actions {
		amis, ok := a.(pluggable.AndMakeItSo)
		if ok {
			amis.Execute()
		}
	}
}

func (t *coreTarget) TearDown() {
	for _, a := range slices.Backward(t.actions) {
		amis, ok := a.(pluggable.AndMakeItSo)
		if ok {
			amis.TearDown()
		}
	}
}

func (d *coreTarget) Present(value any) {
	// If I have understood the flow correctly, if you arrive here without having reported an error in MustBind,
	// then binding is optional and no assignTo has been specified.  So doing nothing is fine.
}
