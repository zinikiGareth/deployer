package testDB

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type tableModel struct {
	loc        *errorsink.Location
	storage    driverbottom.RuntimeStorage
	id         corebottom.CoinId
	testLogger testhelpers.TestStepLogger
	name       string
}

func (bm *tableModel) Loc() *errorsink.Location {
	return bm.loc
}

func (bm *tableModel) ShortDescription() string {
	return fmt.Sprintf("table[%s]", bm.name)
}

func (bm *tableModel) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("table")
	iw.AttrsWhere(bm)
	iw.TextAttr("name", bm.name)
	iw.EndAttrs()
}

var _ driverbottom.Describable = &tableModel{}
