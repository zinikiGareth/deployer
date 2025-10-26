package files

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corestrats"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type fileAction struct {
	exprs []driverbottom.Expr
}

func (da *fileAction) DumpArgs(w driverbottom.IndentWriter) {
	for _, v := range da.exprs {
		w.IndPrintf("%s\n", v.ShortDescription())
	}
}

func (da *fileAction) ShortDescription() string {
	return fmt.Sprintf("%d", len(da.exprs))
}

func (da *fileAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, e := range da.exprs {
		e.Resolve(r)
	}
	return driverbottom.MUST_BE_BOUND
}

func (da *fileAction) DetermineInitialState(tools *corebottom.Tools, loc *errorsink.Location, pres corebottom.ValuePresenter) {
	if tools.Reporter.HasErrors() {
		return
	}

	var paths = []any{}
	for _, e := range da.exprs {
		v := tools.Storage.Eval(e)
		paths = append(paths, v)
	}
	dir, err := NewFileModel(loc, paths)
	if err != nil {
		tools.Reporter.ReportAtf(loc, err.Error())
	} else {
		pres.Present(dir)
	}
}

func (da *fileAction) DetermineDesiredState(tools *corebottom.Tools, loc *errorsink.Location, pres corebottom.ValuePresenter) {
	da.DetermineInitialState(tools, loc, pres)
}

var _ corestrats.CoreActionStrategy = &fileAction{}
