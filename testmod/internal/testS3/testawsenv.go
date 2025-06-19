package testS3

import "ziniki.org/deployer/driver/pkg/driverbottom"

type TestAwsEnv struct {
	Region  string
	Buckets map[string]*BucketCloud
}

func (me *TestAwsEnv) InitMe(storage driverbottom.RuntimeStorage) any {
	me.Region = "us-east-1"
	me.Buckets = make(map[string]*BucketCloud)
	return me
}

func (me *TestAwsEnv) FindBucket(name string) *BucketCloud {
	return me.Buckets[name]
}

func (me *TestAwsEnv) CreateBucket(name string) *BucketCloud {
	if me.Buckets[name] != nil {
		panic("don't create a bucket that already exists")
	}
	ret := NewCloudBucket(name)
	me.Buckets[name] = ret
	return ret
}
