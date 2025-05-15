package testS3

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type assertBucketAction struct {
	tools  *pluggable.Tools
	loc    *errorsink.Location
	bucket pluggable.Expr
	files  []string
}

func (ca *assertBucketAction) Loc() *errorsink.Location {
	return ca.loc
}

func (ca *assertBucketAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("AssertBucketAction")
	w.AttrsWhere(ca)
	w.IndPrintf("bucket: %s\n", ca.bucket.ShortDescription())
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
	ca.bucket.Resolve(r)
}

func (ca *assertBucketAction) Prepare(pres pluggable.ValuePresenter) {
}

func (ca *assertBucketAction) Execute() {
	bucketVar := ca.tools.Storage.Eval(ca.bucket)
	bucket, ok := bucketVar.(*bucketCreator)
	if !ok {
		ca.tools.Reporter.At(ca.bucket.Loc().Line)
		ca.tools.Reporter.Reportf(ca.bucket.Loc().Offset, "was not a bucket: %T", bucketVar)
	}
	// TODO: our test environment needs to create a bucket "in memory"
	// This then needs to be able to be the destination for copy
	// And then we need to be able to retrieve it from "the cloud provider"
	// And test that it has the files we want it to have ...

	for _, f := range ca.files {
		if !bucket.cloud.HasFile(f) {
			ca.tools.Reporter.At(ca.bucket.Loc().Line)
			ca.tools.Reporter.Reportf(ca.bucket.Loc().Offset, "do not have the file %s", f)
		}
	}
}

func (ca *assertBucketAction) TearDown() {
	// Does this just want to do nothing?
	// Or assert that the bucket has gone away?

}
