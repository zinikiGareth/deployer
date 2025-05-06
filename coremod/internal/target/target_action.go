package target

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type action interface {
	pluggable.Definition
	pluggable.Executable
}

type coreTarget struct {
	loc  *errors.Location
	name pluggable.SymbolName

	actions []action
}

func (cc *coreTarget) Add(entry pluggable.Definition) {
	a, ok := entry.(action)
	if !ok {
		panic(fmt.Sprintf("entry %v is not an action", entry))
	}
	cc.actions = append(cc.actions, a)
}

func (t *coreTarget) Loc() *errors.Location {
	return t.loc
}

func (t *coreTarget) Where() *errors.Location {
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
		a.Resolve(r)
	}
}

func (t *coreTarget) Prepare(storage pluggable.RuntimeStorage) pluggable.ExecuteAction {
	for _, a := range t.actions {
		act := a.Prepare(storage)
		storage.BindAction(a, act)
	}
	return nil
}

func (t *coreTarget) Execute(storage pluggable.RuntimeStorage) {
	for _, a := range t.actions {
		av := storage.RetrieveAction(a)
		if av != nil {
			av.Execute(storage)
		}
	}
}
