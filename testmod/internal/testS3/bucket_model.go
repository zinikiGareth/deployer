package testS3

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type bucketModel struct {
	testLogger testhelpers.TestStepLogger
	name       string
	// region     fmt.Stringer
	cloud *BucketCloud

	policy corebottom.PolicyDocument
}

func (bm *bucketModel) Attach(doc corebottom.PolicyDocument) {
	// TODO: this should copy the model
	bm.policy = doc
	bm.testLogger.Log("we need to attach policy with %d items to bucket %s\n", len(doc.Items()), bm.name)
}

func (eb *bucketModel) ObtainDest() corebottom.FileDest {
	return eb.cloud
}

var _ corebottom.PolicyAttacher = &bucketModel{}
var _ corebottom.DestHolder = &bucketModel{}
