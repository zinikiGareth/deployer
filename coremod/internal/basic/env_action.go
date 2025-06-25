package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type EnvAction struct {
	tools   *corebottom.Tools
	loc     *errorsink.Location
	varname driverbottom.Expr
}

func (ea *EnvAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *EnvAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("EnvAction")
	w.AttrsWhere(ea)
	w.TextAttr("varname", ea.varname.String())
	w.EndAttrs()
}

func (ea *EnvAction) ShortDescription() string {
	return fmt.Sprintf("Env[%s]", ea.varname.String())
}

func (ea *EnvAction) Completed() {
}

func (sa *EnvAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	sa.varname.Resolve(r)
	// b.MustBind(&EnvVar{varname: sa.varname})
	// ea.resolved = r.Resolve(ea.what)
	return driverbottom.MUST_BE_BOUND
}

func (ea *EnvAction) DetermineInitialState(pres driverbottom.ValuePresenter) {
	fromVar, ok := ea.tools.Storage.EvalAsStringer(ea.varname)
	if !ok {
		panic("not a stringer")
	}
	pres.Present(NewEnvModel(ea.tools.Reporter, *ea.varname.Loc(), fromVar))
}

func (ea *EnvAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	pres.Unchanged()
}

func (ea *EnvAction) UpdateReality() {

}

func (ea *EnvAction) TearDown() {

}

type EnvVar struct {
	varname driverbottom.Expr
}

func (e *EnvVar) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("EnvVar")
	to.AttrsWhere(e.varname)
	to.TextAttr("var", e.varname.String())
	to.EndAttrs()
}

func (e *EnvVar) Loc() *errorsink.Location {
	return e.varname.Loc()
}

func (e *EnvVar) ShortDescription() string {
	return "EnvVar[" + e.varname.String() + "]"
}

var _ corebottom.ModelBuilder = &EnvAction{}
