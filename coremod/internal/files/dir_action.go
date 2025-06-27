package files

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type dirAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	exprs []driverbottom.Expr
	// res   *PathHolder
}

func (da *dirAction) Loc() *errorsink.Location {
	return da.loc
}

func (da *dirAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("DirAction")
	w.AttrsWhere(da)
	for _, v := range da.exprs {
		w.IndPrintf("%s\n", v.ShortDescription())
	}
	w.EndAttrs()
}

func (da *dirAction) ShortDescription() string {
	return fmt.Sprintf("Dir[%d]", len(da.exprs))
}

func (da *dirAction) Completed() {
}

func (da *dirAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, e := range da.exprs {
		e.Resolve(r)
	}
	// da.res = &PathHolder{loc: da.loc}
	return driverbottom.MUST_BE_BOUND
}

func (da *dirAction) DetermineInitialState(pres driverbottom.ValuePresenter) {
}

func (da *dirAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	if da.tools.Reporter.HasErrors() {
		return
	}

	var paths = []any{}
	for _, e := range da.exprs {
		v := da.tools.Storage.Eval(e)
		paths = append(paths, v)
	}
	dir := NewDirModel(da.loc, paths)
	pres.Present(dir)
}

func (ea *dirAction) UpdateReality() {

}

func (ea *dirAction) TearDown() {

}

var _ corebottom.ModelBuilder = &dirAction{}
