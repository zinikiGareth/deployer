package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type bucketCreator struct {
	tools *corebottom.Tools

	loc    *errorsink.Location
	name   string
	props  map[driverbottom.Identifier]driverbottom.Expr
	env    *TestAwsEnv
	cloud  *BucketCloud
	policy corebottom.PolicyDocument
}

func (b *bucketCreator) Loc() *errorsink.Location {
	return b.loc
}

func (b *bucketCreator) ShortDescription() string {
	return "test.S3.Bucket[" + b.name + "]"
}

func (b *bucketCreator) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("test.S3.Bucket[")
	iw.AttrsWhere(b)
	iw.TextAttr("named", b.name)
	for i, e := range b.props {
		iw.NestedAttr(i.Id(), e)
	}
	iw.EndAttrs()
}

func (b *bucketCreator) DetermineInitialState(pres driverbottom.ValuePresenter) {
}

func (b *bucketCreator) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	tmp := b.tools.Recall.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}

	tmp = b.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	b.env = testAwsEnv
	testLogger.Log("ensuring bucket exists action %s\n", b.String())
	pres.Present(b)
}

func (eb *bucketCreator) UpdateReality() {
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

	eb.cloud = b
}

func (eb *bucketCreator) TearDown() {
	tmp := eb.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	testLogger.Log("we need to delete a bucket called %s in region %s\n", eb.name, eb.env.Region)
}

func (eb *bucketCreator) ObtainDest() corebottom.FileDest {
	return eb.cloud
}

func (eb *bucketCreator) String() string {
	return fmt.Sprintf("EnsureBucket[%s:%s]", eb.env.Region, eb.name)
}

// TODO: this should be on the model
func (eb *bucketCreator) Attach(doc corebottom.PolicyDocument) {
	// TODO: this should copy the model
	eb.policy = doc
	tmp := eb.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	testLogger.Log("we need to attach policy with %d items to bucket %s\n", len(doc.Items()), eb.name)
}

var _ corebottom.Ensurable = &bucketCreator{}
var _ corebottom.PolicyAttacher = &bucketCreator{} // should be model
