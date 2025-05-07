package testS3

import "ziniki.org/deployer/deployer/pkg/pluggable"

// A BucketNoun is a noun that is registered with the deployer which is what gets us here.
// The definition of a noun (currently) includes the concept that each instance has a proper name
// and thus at some point (during parsing) "CreateWithName" can/will be called.
type BucketNoun struct{}

// Not that I particularly like the word "create" here; it feels more like stamping out coins
// or instantiating or something.  This is absolutely not creating a bucket, it is creating a creator.
func (b *BucketNoun) CreateWithName(named string, assignTo pluggable.Identifier) any {
	return &bucketCreator{name: named, assignTo: assignTo}
}

func (b *BucketNoun) ShortDescription() string {
	return "test.S3.Bucket[]"
}
