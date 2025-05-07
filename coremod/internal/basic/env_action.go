package basic

import (
	"fmt"
	"log"
	"os"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type EnvAction struct {
	tools    *pluggable.Tools
	loc      *errors.Location
	expr     pluggable.Expr
	assignTo pluggable.Identifier
}

func (ea *EnvAction) Loc() *errors.Location {
	return ea.loc
}

func (ea *EnvAction) Where() *errors.Location {
	return ea.loc
}

func (ea *EnvAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("EnvAction")
	w.AttrsWhere(ea)
	w.IndPrintf("%s\n", ea.expr.String())
	w.IndPrintf("%s\n", ea.assignTo.String())
	w.EndAttrs()
}

func (ea *EnvAction) ShortDescription() string {
	return fmt.Sprintf("Env[%s=>%s]", ea.expr.String(), ea.assignTo.Id())
}

func (ea *EnvAction) Completed() {
}

func (sa *EnvAction) Resolve(r pluggable.Resolver) {
	// ea.resolved = r.Resolve(ea.what)
}

func (ea *EnvAction) Prepare(runtime pluggable.RuntimeStorage) pluggable.ExecuteAction {
	// TODO: I think ALL this should really be something like e.Eval(runtime).ToString()
	str, ok := ea.expr.(pluggable.String)
	if ok {
		val := os.Getenv(str.Text())
		runtime.Bind(pluggable.SymbolName(ea.assignTo.Id()), val)
		return nil
	} else {
		log.Fatalf("cannot show %v", ea.expr)
		return nil
	}
}
