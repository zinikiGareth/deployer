package basic

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type ShowAction struct {
	tools *pluggable.Tools
	loc   *errors.Location
	exprs []pluggable.Expr
}

func (sa *ShowAction) Loc() *errors.Location {
	return sa.loc
}

func (sa *ShowAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("ShowAction")
	w.AttrsWhere(sa)
	w.ListAttr("exprs")
	for _, v := range sa.exprs {
		w.IndPrintf("%s\n", v.String())
	}
	w.EndList()
	w.EndAttrs()
}

func (sa *ShowAction) ShortDescription() string {
	return fmt.Sprintf("Show[%d]", len(sa.exprs))
}

func (sa *ShowAction) Completed() {
}

func (sa *ShowAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	// ea.resolved = r.Resolve(ea.what)
}

func (sa *ShowAction) Prepare(pres pluggable.ValuePresenter) {
}

func (sa *ShowAction) Execute() {
	// This probably needs a lot more work and a lot more infrastructure
	// I don't think I even know *how* I expect it to work at the moment ...

	// For starters, I instinctively feel I should be writing to stdout, but golden tester doesn't capture that
	// So I deffo need a proxy writer.  But is this the right abstraction?
	tmp := sa.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	logger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		// TODO: make it point to something with Log => fmt.Printf()
		panic("could not get logger")
	}

	for i, e := range sa.exprs {
		if i > 0 {
			logger.Log(" ")
		}
		str := sa.tools.Storage.EvalAsString(e)
		logger.Log("%s", str)
	}
	logger.Log("\n")
}
