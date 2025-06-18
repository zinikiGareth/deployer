package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type bucketFinder struct {
	tools *external.Tools

	loc  *errorsink.Location
	name string

	env   *TestAwsEnv
	cloud *BucketCloud
}

func (b *bucketFinder) Loc() *errorsink.Location {
	return b.loc
}

func (b *bucketFinder) ShortDescription() string {
	return "test.S3.Bucket[" + b.name + "]"
}

func (b *bucketFinder) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("test.S3.Bucket[")
	iw.AttrsWhere(b)
	iw.TextAttr("named", b.name)
	iw.EndAttrs()
}

// This is called during the "Prepare" phase
func (fb *bucketFinder) BuildModel(pres pluggable.ValuePresenter) {
	tmp := fb.tools.Recall.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}

	tmp = fb.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	fb.env = testAwsEnv
	testLogger.Log("finding bucket action %s\n", fb.String())

	bc := fb.env.FindBucket(fb.name)
	if bc != nil {
		testLogger.Log("found bucket %s in region %s\n", fb.name, fb.env.Region)
	}
	fb.cloud = bc

	pres.Present(fb)
}

func (eb *bucketFinder) UpdateReality() {
}

func (eb *bucketFinder) ObtainDest() files.FileDest {
	return eb.cloud
}

func (eb *bucketFinder) String() string {
	return fmt.Sprintf("FindBucket[%s:%s]", eb.env.Region, eb.name)
}
