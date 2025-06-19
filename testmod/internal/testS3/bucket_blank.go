package testS3

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type BucketBlank struct{}

func (b *BucketBlank) Find(ct *corebottom.Tools, loc *errorsink.Location, named string) any {
	return &bucketFinder{tools: ct, loc: loc, name: named}
}

func (b *BucketBlank) Mint(ct *corebottom.Tools, loc *errorsink.Location, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown corebottom.TearDown) any {
	return &bucketCreator{tools: ct, loc: loc, name: named}
}

func (b *BucketBlank) Loc() *errorsink.Location {
	panic("not implemented")
}

func (b *BucketBlank) ShortDescription() string {
	return "test.S3.Bucket[]"
}

func (b *BucketBlank) DumpTo(iw driverbottom.IndentWriter) {
	panic("not implemented")
}

var _ corebottom.Blank = &BucketBlank{}
