package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type assertBucketAction struct {
	tools  *corebottom.Tools
	loc    *errorsink.Location
	bucket driverbottom.Expr
	files  []string
}

func (ca *assertBucketAction) Loc() *errorsink.Location {
	return ca.loc
}

func (ca *assertBucketAction) DumpTo(w driverbottom.IndentWriter) {
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

func (ca *assertBucketAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ca.bucket.Resolve(r)
	return driverbottom.NO_VALUE
}

func (ca *assertBucketAction) DetermineInitialState(pres corebottom.ValuePresenter) {
}

func (ca *assertBucketAction) DetermineDesiredState(pres corebottom.ValuePresenter) {
}

func (ca *assertBucketAction) ShouldDestroy() bool {
	return false
}

func (ca *assertBucketAction) UpdateReality() {
	bucketVar := ca.tools.Storage.Eval(ca.bucket)
	bucket, ok := bucketVar.(*bucketModel)
	if !ok {
		ca.tools.Reporter.ReportAtf(ca.bucket.Loc(), "was not a bucket: %T", bucketVar)
	}
	// TODO: our test environment needs to create a bucket "in memory"
	// This then needs to be able to be the destination for copy
	// And then we need to be able to retrieve it from "the cloud provider"
	// And test that it has the files we want it to have ...

	for _, f := range ca.files {
		if !bucket.cloud.HasFile(f) {
			ca.tools.Reporter.ReportAtf(ca.bucket.Loc(), "do not have the file %s", f)
		}
	}
}

func (ca *assertBucketAction) TearDown() {
	// Does this just want to do nothing?
	// Or assert that the bucket has gone away?

}

var _ corebottom.RealityShifter = &assertBucketAction{}
