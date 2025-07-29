package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type configCreator struct {
	tools      *corebottom.Tools
	testLogger testhelpers.TestStepLogger

	loc   *errorsink.Location
	coin  driverbottom.ResolvableHolder
	name  string
	props map[driverbottom.Identifier]driverbottom.Expr
}

func (cc *configCreator) Loc() *errorsink.Location {
	return cc.loc
}

func (cc *configCreator) ShortDescription() string {
	return "test.S3.Config[" + cc.name + "]"
}

func (cc *configCreator) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("test.S3.Config[")
	iw.AttrsWhere(cc)
	iw.TextAttr("named", cc.name)
	for i, e := range cc.props {
		iw.NestedAttr(i.Id(), e)
	}
	iw.EndAttrs()
}

func (cc *configCreator) CoinId() corebottom.CoinId {
	return cc.coin
}

func (cc *configCreator) Create(pres corebottom.ValuePresenter) {
	tmp := cc.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}
	cc.testLogger = testLogger

	if cc.coin == nil {
		panic("need a coin id")
	}
	cc.testLogger.Log("creating model for bucket %s\n", cc.String())
	model := &configModel{loc: cc.loc, storage: cc.tools.Storage, id: cc.coin, name: cc.name, testLogger: cc.testLogger}
	pres.Present(model)
}

func (cc *configCreator) String() string {
	return fmt.Sprintf("CreateConfig[%s]", cc.name)
}

var _ corebottom.MemoryCoinCreator = &configCreator{}
