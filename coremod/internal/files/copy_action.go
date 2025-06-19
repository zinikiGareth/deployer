package files

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type copyAction struct {
	tools *external.Tools
	loc   *errorsink.Location
	exprs []pluggable.Expr

	Src  files.FileSource
	Dest files.DestHolder
}

func (ca *copyAction) Loc() *errorsink.Location {
	return ca.loc
}

func (ca *copyAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("CopyAction")
	w.AttrsWhere(ca)
	for _, v := range ca.exprs {
		w.IndPrintf("%s\n", v.ShortDescription())
	}
	w.EndAttrs()
}

func (ca *copyAction) ShortDescription() string {
	return fmt.Sprintf("Dir[%d]", len(ca.exprs))
}

func (ca *copyAction) Completed() {
}

func (ca *copyAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	for _, e := range ca.exprs {
		e.Resolve(r)
	}
	return pluggable.NO_VALUE
}

func (ca *copyAction) BuildModel(pres pluggable.ValuePresenter) {
	if ca.tools.Reporter.HasErrors() {
		return
	}

	from := ca.exprs[0]
	copyFrom := ca.tools.Storage.Eval(from)
	copyFS, ok := copyFrom.(files.FileSource)
	if !ok {
		ca.tools.Reporter.At(from.Loc().Line)
		ca.tools.Reporter.Report(from.Loc().Offset, "was not a file source")
		return
	}
	ca.Src = copyFS
	destVar := ca.tools.Storage.Eval(ca.exprs[1])
	dest, ok := destVar.(files.DestHolder)
	if !ok {
		panic(fmt.Sprintf("dest was %T not a DestHolder", destVar))
	}
	ca.Dest = dest
}

func (ca *copyAction) UpdateReality() {
	d := ca.Dest.ObtainDest()
	ca.Src.PourAll(d)
}

func (ca *copyAction) TearDown() {
	// should we delete or leave alone?
	// need user to tell us
}
