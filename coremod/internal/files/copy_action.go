package files

import (
	"fmt"
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type copyAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	exprs []driverbottom.Expr

	Src  corebottom.FileSource
	Dest corebottom.DestHolder
}

func (ca *copyAction) Loc() *errorsink.Location {
	return ca.loc
}

func (ca *copyAction) DumpTo(w driverbottom.IndentWriter) {
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

func (ca *copyAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, e := range ca.exprs {
		e.Resolve(r)
	}
	return driverbottom.NO_VALUE
}

func (ca *copyAction) DetermineInitialState(pres driverbottom.ValuePresenter) {
}

func (ca *copyAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	if ca.tools.Reporter.HasErrors() {
		return
	}

	from := ca.exprs[0]
	copyFrom := ca.tools.Storage.Eval(from)
	copyFS, ok := copyFrom.(corebottom.FileSource)
	if !ok {
		log.Printf("copyFrom was %T\n", copyFrom)
		ca.tools.Reporter.At(from.Loc().Line)
		ca.tools.Reporter.Report(from.Loc().Offset, "was not a file source")
		return
	}
	ca.Src = copyFS
	destVar := ca.tools.Storage.Eval(ca.exprs[1])
	dest, ok := destVar.(corebottom.DestHolder)
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

var _ corebottom.ModelBuilder = &copyAction{}
