package testS3

import (
	"fmt"
	"reflect"

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

}

func (b *bucket) Ensure(runtime pluggable.RuntimeStorage) pluggable.ExecuteAction {
	tmp := runtime.ObtainDriver(reflect.TypeOf(TestAwsEnv{}))
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}

	fmt.Printf("we want to ensure a bucket in region %s\n", testAwsEnv.Region)

	return &ensureBucket{env: testAwsEnv, bucket: b}
}

type BucketNoun struct{}

func (b *BucketNoun) CreateWithName(named string) any {
	return &bucket{name: named}
}

func (b *BucketNoun) ShortDescription() string {
	return "test.S3.Bucket[]"
}
