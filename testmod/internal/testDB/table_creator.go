package testDB

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
	"ziniki.org/deployer/testmod/internal/testenv"
)

type tableCreator struct {
	tools      *corebottom.Tools
	env        *testenv.TestAwsEnv
	testLogger testhelpers.TestStepLogger

	loc   *errorsink.Location
	coin  driverbottom.Holder
	name  string
	props map[driverbottom.Identifier]driverbottom.Expr
	model *tableModel
}

func (b *tableCreator) Loc() *errorsink.Location {
	return b.loc
}

func (b *tableCreator) ShortDescription() string {
	return "test.DB.Table[" + b.name + "]"
}

func (b *tableCreator) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("test.S3.Bucket[")
	iw.AttrsWhere(b)
	iw.TextAttr("named", b.name)
	for i, e := range b.props {
		iw.NestedAttr(i.Id(), e)
	}
	iw.EndAttrs()
}

func (b *tableCreator) CoinId() corebottom.CoinId {
	return b.coin
}

func (b *tableCreator) DetermineInitialState(pres corebottom.ValuePresenter) {
	tmp := b.tools.Recall.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*testenv.TestAwsEnv)
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

	testLogger.Log("looking for table %s\n", b.String())
	// TODO: the test infrastructure should have the ability to have these in place
	pres.NotFound()
}

func (b *tableCreator) DetermineDesiredState(pres corebottom.ValuePresenter) {
	if b.coin == nil {
		panic("need a coin id")
	}
	b.testLogger.Log("creating model for table %s\n", b.String())
	b.model = &tableModel{loc: b.loc, storage: b.tools.Storage, id: b.coin, name: b.name, testLogger: b.testLogger}
	pres.Present(b.model)
}

func (eb *tableCreator) UpdateReality() {
	/*
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
	*/
}

func (eb *tableCreator) TearDown() {
	tmp := eb.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	testLogger.Log("we need to delete a bucket called %s in region %s\n", eb.name, eb.env.Region)
}

func (eb *tableCreator) String() string {
	return fmt.Sprintf("EnsureBucket[%s:%s]", eb.env.Region, eb.name)
}

var _ corebottom.FindCoin = &tableCreator{}
var _ corebottom.Ensurable = &tableCreator{}
