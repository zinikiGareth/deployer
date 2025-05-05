package basic

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ShowAction struct {
	tools *pluggable.Tools
	loc   *errors.Location
	exprs []pluggable.Expr
}

func (sa *ShowAction) Loc() *errors.Location {
	return sa.loc
}

func (sa *ShowAction) Where() *errors.Location {
	return sa.loc
}

func (sa *ShowAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("ShowCommand")
	w.AttrsWhere(sa)
	for _, v := range sa.exprs {
		w.IndPrintf("%s\n", v.String())
	}
	w.EndAttrs()
}

func (sa *ShowAction) ShortDescription() string {
	return fmt.Sprintf("Show[%d]", len(sa.exprs))
}

func (sa *ShowAction) Completed() {
}

func (sa *ShowAction) Resolve(r pluggable.Resolver) {
	// ea.resolved = r.Resolve(ea.what)
}

func (sa *ShowAction) Prepare(runtime pluggable.RuntimeStorage) (pluggable.ExecuteAction, any) {
	return nil, nil
}
