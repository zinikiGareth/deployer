package target

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type coreTarget struct {
	loc  *errors.Location
	name pluggable.SymbolName

	actions []pluggable.Action
}

func (cc *coreTarget) Add(entry pluggable.Action) {
	cc.actions = append(cc.actions, entry)
}

func (t *coreTarget) Loc() *errors.Location {
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
		a.Resolve(r, t)
	}
}

func (t *coreTarget) Prepare() {
	for _, a := range t.actions {
		a.Prepare(t)
	}
}

func (t *coreTarget) Execute() {
	for _, a := range t.actions {
		a.Execute()
	}
}

func (d *coreTarget) MayBind(val pluggable.Describable) {
	// There is nothing to bind ...
}

func (d *coreTarget) MustBind(val pluggable.Describable) {
	panic("assignTo is not specified") // should be an error
}

func (d *coreTarget) Present(value any) {
	// If I have understood the flow correctly, if you arrive here without having reported an error in MustBind,
	// then binding is optional and no assignTo has been specified.  So doing nothing is fine.
}
