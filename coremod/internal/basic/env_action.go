package basic

import (
	"fmt"
	"os"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnvAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	varname pluggable.Expr
}

func (ea *EnvAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *EnvAction) DumpTo(w pluggable.IndentWriter) {
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

func (sa *EnvAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	sa.varname.Resolve(r)
	// b.MustBind(&EnvVar{varname: sa.varname})
	// ea.resolved = r.Resolve(ea.what)
	return pluggable.MUST_BE_BOUND
}

func (ea *EnvAction) Prepare(pres pluggable.ValuePresenter) {
	// TODO: I think ALL this should really be something like e.Eval(runtime).ToString()
	str := ea.tools.Storage.EvalAsString(ea.varname)
	val := os.Getenv(str)
	if val == "" {
		ea.tools.Reporter.At(ea.varname.Loc().Line)
		ea.tools.Reporter.Reportf(ea.varname.Loc().Offset, "the env var %s is not set", str)
		return
	}
	pres.Present(val)
}

func (ea *EnvAction) Execute() {

}

func (ea *EnvAction) TearDown() {

}

type EnvVar struct {
	varname pluggable.Expr
}

func (e *EnvVar) DumpTo(to pluggable.IndentWriter) {
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
