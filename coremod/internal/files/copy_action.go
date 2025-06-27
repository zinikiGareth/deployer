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
	coin  corebottom.CoinId
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
	ca.coin = corebottom.CoinId(ca.tools.Storage.NewObjId(ca.loc))
	return driverbottom.NO_VALUE
}

// TODO: I think the intent here should be that there IS a model, and that the model reflects the destination contents.
// THUS, initial state should go and find what is there (name, date, length)
// AND desired state should be what is in the source location (name, date, length)

func (ca *copyAction) DetermineInitialState(pres driverbottom.ValuePresenter) {
	model := ca.basicModel()
	// TODO: add in files already in Dest
	ca.tools.Storage.Bind(ca.coin, model)
	pres.Present(model)
}

func (ca *copyAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	model := ca.basicModel()
	ca.tools.Storage.Bind(ca.coin, model)
	// TODO: add in files already in Src
	pres.Present(model)
}

func (ca *copyAction) basicModel() *CopyModel {
	if ca.tools.Reporter.HasErrors() {
		return nil
	}

	from := ca.exprs[0]
	copyFrom := ca.tools.Storage.Eval(from)
	copyFS, ok := copyFrom.(corebottom.FileSource)
	if !ok {
		log.Printf("copyFrom was %T\n", copyFrom)
		ca.tools.Reporter.At(from.Loc().Line)
		ca.tools.Reporter.Report(from.Loc().Offset, "was not a file source")
		return nil
	}
	destVar := ca.tools.Storage.Eval(ca.exprs[1])
	var dest corebottom.DestHolder
	if destVar != nil {
		dest, ok = destVar.(corebottom.DestHolder)
		if !ok {
			panic(fmt.Sprintf("dest was %T not a DestHolder", destVar))
		}
	}
	return &CopyModel{Src: copyFS, Dest: dest}
}

func (ca *copyAction) UpdateReality() {
	model := ca.tools.Storage.GetCoin(ca.coin, corebottom.DETERMINE_DESIRED_MODE).(*CopyModel)
	d := model.Dest.ObtainDest()
	model.Src.PourAll(d)
}

func (ca *copyAction) TearDown() {
	// should we delete or leave alone?
	// need user to tell us
}

var _ corebottom.ModelBuilder = &copyAction{}
