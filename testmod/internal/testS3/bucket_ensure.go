package testS3

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type ensureBucket struct {
	env    *TestAwsEnv
	bucket *bucketCreator
}

func (eb *ensureBucket) Execute(runtime pluggable.RuntimeStorage) {
	tmp := runtime.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	testLogger.Log("we want to ensure a bucket called %s in region %s\n", eb.bucket.name, eb.env.Region)
}

func (eb *ensureBucket) String() string {
	return fmt.Sprintf("EnsureBucket[%s:%s]", eb.env.Region, eb.bucket.name)
}
