package vars

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type varBinding struct {
	Location *errors.Location
	Bind     pluggable.Identifier
	BindTo   pluggable.Definition
}

func (v *varBinding) Loc() *errors.Location {
	return v.Location
}

// DumpTo implements target.action.
func (v *varBinding) DumpTo(to pluggable.IndentWriter) {
	to.Intro("bind_var")
	to.AttrsWhere(v)
	to.TextAttr("var", v.Bind.Id())
	to.TextAttr("value", v.BindTo.ShortDescription())
	to.EndAttrs()
}

// Prepare implements target.action.
func (v *varBinding) Prepare(runtime pluggable.RuntimeStorage) pluggable.ExecuteAction {
	tmp := runtime.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if ok {
		testLogger.Log("bind \"%s\" to %s\n", v.Bind.Id(), v.BindTo.ShortDescription())
	}

	return nil
}

// Resolve implements target.action.
func (v *varBinding) Resolve(r pluggable.Resolver) {
}

// ShortDescription implements target.action.
func (v *varBinding) ShortDescription() string {
	panic("unimplemented")
}

// What implements target.action.
func (v *varBinding) What() pluggable.SymbolType {
	panic("unimplemented")
}

// Where implements target.action.
func (v *varBinding) Where() *errors.Location {
	panic("unimplemented")
}

func BindVar(identifier pluggable.Identifier, entry pluggable.Definition) *varBinding {
	ret := &varBinding{Location: identifier.Loc(), Bind: identifier, BindTo: entry}
	return ret
}
