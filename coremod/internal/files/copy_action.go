package files

import (
	"fmt"
	"log"

	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type copyAction struct {
	tools *pluggable.Tools
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

func (ca *copyAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	for _, e := range ca.exprs {
		e.Resolve(r)
	}
}

func (ca *copyAction) Prepare(pres pluggable.ValuePresenter) {
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
	log.Printf("%v %v\n", copyFS, dest)
}

func (ca *copyAction) Execute() {
	d := ca.Dest.ObtainDest()
	ca.Src.PourOut("intro", d)
	/*
		files, err := os.ReadDir(src.File)
		if err != nil {
			panic(err)
		}

		d := dest.ObtainDest()
		for _, f := range files {
			d.PourInto(f.Name())
		}
	*/
}

func (ca *copyAction) TearDown() {
	// should we delete or leave alone?
	// need user to tell us
}
