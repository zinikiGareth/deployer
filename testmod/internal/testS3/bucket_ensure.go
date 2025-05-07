package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type ensureBucket struct {
	env    *TestAwsEnv
	bucket *bucketCreator
	cloud  *BucketCloud
}

func (eb *ensureBucket) Execute(runtime pluggable.RuntimeStorage) {
	tmp := runtime.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	b := eb.env.FindBucket(eb.bucket.name)
	if b != nil {
		testLogger.Log("the bucket %s in region %s already exists\n", eb.bucket.name, eb.env.Region)
	} else {
		testLogger.Log("we need to create a bucket called %s in region %s\n", eb.bucket.name, eb.env.Region)
		// TODO: we should also handle all the properties we have stored somewhere ...
		b = eb.env.CreateBucket(eb.bucket.name)
	}

	eb.cloud = b
}

func (eb *ensureBucket) ObtainDest() files.FileDest {
	return eb.cloud
}

func (eb *ensureBucket) String() string {
	return fmt.Sprintf("EnsureBucket[%s:%s]", eb.env.Region, eb.bucket.name)
}
