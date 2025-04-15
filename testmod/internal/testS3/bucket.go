package testS3

import (
	"ziniki.org/deployer/golden/pkg/testing"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type bucket struct {
	name string
}

type ensureBucket struct {
	env    *TestAwsEnv
	bucket *bucket
}

func (eb *ensureBucket) Execute(runtime pluggable.RuntimeStorage) {
	tmp := runtime.ObtainDriver("testing.TestStepLogger")
	testLogger, ok := tmp.(testing.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	testLogger.Log("we want to ensure a bucket called %s in region %s\n", eb.bucket.name, eb.env.Region)
}

func (b *bucket) Ensure(runtime pluggable.RuntimeStorage) pluggable.ExecuteAction {
	tmp := runtime.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}

	return &ensureBucket{env: testAwsEnv, bucket: b}
}

type BucketNoun struct{}

func (b *BucketNoun) CreateWithName(named string) any {
	return &bucket{name: named}
}

func (b *BucketNoun) ShortDescription() string {
	return "test.S3.Bucket[]"
}
