package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type configModel struct {
	loc        *errorsink.Location
	storage    driverbottom.RuntimeStorage
	id         corebottom.CoinId
	testLogger testhelpers.TestStepLogger
	name       string
}

func (bm *configModel) Loc() *errorsink.Location {
	return bm.loc
}

func (bm *configModel) ShortDescription() string {
	return fmt.Sprintf("config[%s]", bm.name)
}

func (bm *configModel) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("config")
	iw.AttrsWhere(bm)
	iw.TextAttr("name", bm.name)
	// iw.NestedAttr("policy", bm.policy)
	iw.EndAttrs()
}

var _ driverbottom.Describable = &configModel{}
