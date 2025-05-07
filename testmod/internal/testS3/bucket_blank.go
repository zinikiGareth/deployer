package testS3

import "ziniki.org/deployer/deployer/pkg/pluggable"

type BucketBlank struct{}

func (b *BucketBlank) Mint(tools *pluggable.Tools, named string) any {
	return &bucketCreator{tools: tools, name: named}
}

func (b *BucketBlank) ShortDescription() string {
	return "test.S3.Bucket[]"
}
