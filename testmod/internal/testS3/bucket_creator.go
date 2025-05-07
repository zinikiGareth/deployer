package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type bucketCreator struct {
	tools *pluggable.Tools

	name string

	env   *TestAwsEnv
	cloud *BucketCloud
}

// This is called during the "Prepare" phase
func (b *bucketCreator) Prepare() {
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
	// if b.assignTo != nil {
	// 	b.tools.Storage.Bind(pluggable.SymbolName(b.assignTo.Id()), eb)
	// }
}

func (eb *bucketCreator) Execute() {
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

func (eb *bucketCreator) ObtainDest() files.FileDest {
	return eb.cloud
}

func (eb *bucketCreator) String() string {
	return fmt.Sprintf("EnsureBucket[%s:%s]", eb.env.Region, eb.name)
}
