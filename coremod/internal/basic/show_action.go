package basic

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type ShowAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	exprs []driverbottom.Expr
}

func (sa *ShowAction) Loc() *errorsink.Location {
	return sa.loc
}

func (sa *ShowAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("ShowAction")
	w.AttrsWhere(sa)
	w.ListAttr("exprs")
	for _, v := range sa.exprs {
		w.IndPrintf("%s\n", v.ShortDescription())
	}
	w.EndList()
	w.EndAttrs()
}

func (sa *ShowAction) ShortDescription() string {
	return fmt.Sprintf("Show[%d]", len(sa.exprs))
}

func (sa *ShowAction) Completed() {
}

func (sa *ShowAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, e := range sa.exprs {
		e.Resolve(r)
	}
	return driverbottom.NO_VALUE
}

func (sa *ShowAction) DetermineInitialState(pres driverbottom.ValuePresenter) {
}

func (sa *ShowAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
}

func (sa *ShowAction) UpdateReality() {
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
		str, ok := sa.tools.Storage.EvalAsStringer(e)
		if !ok {
			panic("not a stringer")
		}
		logger.Log("%s", str)
	}
	logger.Log("\n")
}

func (sa *ShowAction) TearDown() {
	// This seems one option among several
	sa.UpdateReality()
}

var _ corebottom.ModelBuilder = &ShowAction{}
