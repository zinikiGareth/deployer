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
	coin  driverbottom.Holder
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
	if cc.coin == nil {
		panic("need a coin id")
	}
	cc.testLogger.Log("creating model for bucket %s\n", cc.String())
	model := &configModel{loc: cc.loc, storage: cc.tools.Storage, id: cc.coin, name: cc.name, testLogger: cc.testLogger}
	pres.Present(model)
}

/*
func (b *configCreator) DetermineInitialState(pres corebottom.ValuePresenter) {
	tmp := b.tools.Recall.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}
	b.env = testAwsEnv

	tmp = b.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}
	b.testLogger = testLogger

	testLogger.Log("looking for bucket %s\n", b.String())
	// TODO: the test infrastructure should have the ability to have these in place
	pres.NotFound()
}

func (eb *configCreator) UpdateReality() {
	tmp := eb.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	b := eb.env.FindBucket(eb.name)
	if b != nil {
		testLogger.Log("the bucket %s in region %s already exists\n", eb.name, eb.env.Region)
	} else {
		testLogger.Log("we need to create a bucket called %s in region %s\n", eb.name, eb.env.Region)
		// TODO: we should also handle all the properties we have stored somewhere ...
		b = eb.env.CreateBucket(eb.name)
	}

	eb.model.cloud = b
}

func (eb *configCreator) TearDown() {
	tmp := eb.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	testLogger.Log("we need to delete a bucket called %s in region %s\n", eb.name, eb.env.Region)
}
*/

func (cc *configCreator) String() string {
	return fmt.Sprintf("CreateConfig[%s]", cc.name)
}

var _ corebottom.MemoryCoinCreator = &configCreator{}
