package testS3

import "ziniki.org/deployer/deployer/pkg/pluggable"

type TestAwsEnv struct {
	Region string
}

func (me *TestAwsEnv) InitMe(storage pluggable.RuntimeStorage) {
	me.Region = "us-east-1"
}
