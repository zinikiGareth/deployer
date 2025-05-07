package testS3

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type assertBucketAction struct {
	tools  *pluggable.Tools
	loc    *errors.Location
	bucket pluggable.Expr
	files  []string
}

func (ca *assertBucketAction) Loc() *errors.Location {
	return ca.loc
}

func (ca *assertBucketAction) Where() *errors.Location {
	return ca.loc
}

func (ca *assertBucketAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("AssertBucketAction")
	w.AttrsWhere(ca)
	w.IndPrintf("bucket: %s\n", ca.bucket.String())
	for _, f := range ca.files {
		w.IndPrintf("  assert file: %s\n", f)
	}
	w.EndAttrs()
}

func (ca *assertBucketAction) ShortDescription() string {
	return fmt.Sprintf("AssertBucket[%s]", ca.bucket.String())
}

func (ca *assertBucketAction) Completed() {
}

func (ca *assertBucketAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	// ea.resolved = r.Resolve(ea.what)
}

func (ca *assertBucketAction) Prepare() {
}

func (ca *assertBucketAction) Execute() {
	bucketVar := ca.tools.Storage.Eval(ca.bucket)
	bucket, ok := bucketVar.(*ensureBucket)
	if !ok {
		panic("not the bucket i was looking for")
	}
	// TODO: our test environment needs to create a bucket "in memory"
	// This then needs to be able to be the destination for copy
	// And then we need to be able to retrieve it from "the cloud provider"
	// And test that it has the files we want it to have ...

	for _, f := range ca.files {
		if !bucket.cloud.HasFile(f) {
			panic("do not have the file")
		}
	}
}
