package testS3

import "ziniki.org/deployer/deployer/pkg/pluggable"

type bucket struct{
	name string
}

func (b *bucket) Ensure(runtime pluggable.RuntimeStorage) {

}

type BucketNoun struct{}

func (b *BucketNoun) CreateWithName(named string) any {
	return &bucket{name: named}
}

func (b *BucketNoun) ShortDescription() string {
	return "test.S3.Bucket[]"
}
