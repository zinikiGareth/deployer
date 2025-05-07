package testS3

type BucketCloud struct {
	name     string
	contents map[string]*BucketEntry
}

type BucketEntry struct {
	Key  string
	Data []byte
}

func (b *BucketCloud) HasFile(name string) bool {
	return b.contents[name] != nil
}

func (b *BucketCloud) Pour(name string) {
	entry := &BucketEntry{Key: name}
	b.contents[name] = entry
}

func NewCloudBucket(name string) *BucketCloud {
	return &BucketCloud{name: name, contents: make(map[string]*BucketEntry)}
}
