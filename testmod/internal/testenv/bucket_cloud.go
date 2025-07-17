package testenv

import (
	"io"

	"ziniki.org/deployer/coremod/pkg/corebottom"
)

type BucketCloud struct {
	name     string
	contents map[string]*BucketEntry
}

type BucketEntry struct {
	Key    string
	Data   []byte
	Nested *BucketCloud
}

func (b *BucketCloud) HasFile(name string) bool {
	return b.contents[name] != nil
}

func (b *BucketCloud) Relative(name string) (corebottom.FileDest, error) {
	nested := NewCloudBucket(name)
	entry := &BucketEntry{Key: name, Nested: nested}
	b.contents[name] = entry
	return nested, nil
}

func (b *BucketCloud) PourInto(name string, contents io.Reader) {
	entry := &BucketEntry{Key: name}
	b.contents[name] = entry
}

func NewCloudBucket(name string) *BucketCloud {
	return &BucketCloud{name: name, contents: make(map[string]*BucketEntry)}
}

var _ corebottom.FileDest = &BucketCloud{}
