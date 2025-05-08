package basic

import (
	"fmt"
	"log"
	"os"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnvAction struct {
	tools   *pluggable.Tools
	loc     *errors.Location
	varname pluggable.Expr
}

func (ea *EnvAction) Loc() *errors.Location {
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

func (sa *EnvAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	// ea.resolved = r.Resolve(ea.what)
}

func (ea *EnvAction) Prepare(pres pluggable.ValuePresenter) {
	// TODO: I think ALL this should really be something like e.Eval(runtime).ToString()
	str, ok := ea.varname.(pluggable.String)
	if ok {
		val := os.Getenv(str.Text())
		pres.Present(val)
		return
	} else {
		log.Fatalf("cannot show %v", ea.varname)
		return
	}
}

func (ea *EnvAction) Execute() {

}
