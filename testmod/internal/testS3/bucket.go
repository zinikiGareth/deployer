package testS3

import (
	"fmt"
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type bucket struct {
	name string
}

func (b *bucket) Ensure(runtime pluggable.RuntimeStorage) {
	if runtime.IsMode(pluggable.EXECUTE_MODE) {
		tmp := runtime.ObtainDriver(reflect.TypeOf(TestAwsEnv{}))
		testAwsEnv, ok := tmp.(*TestAwsEnv)
		if !ok {
			panic("could not cast env to TestAwsEnv")
		}

		fmt.Printf("we want to ensure a bucket in region %s\n", testAwsEnv.Region)
	}
}

type BucketNoun struct{}

func (b *BucketNoun) CreateWithName(named string) any {
	return &bucket{name: named}
}

func (b *BucketNoun) ShortDescription() string {
	return "test.S3.Bucket[]"
}
