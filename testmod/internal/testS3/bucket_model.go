package testS3

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type bucketModel struct {
	loc        *errorsink.Location
	storage    driverbottom.RuntimeStorage
	id         corebottom.CoinId
	testLogger testhelpers.TestStepLogger
	name       string
	cloud      *BucketCloud

	policy corebottom.PolicyDocument
}

func (bm *bucketModel) Loc() *errorsink.Location {
	return bm.loc
}

func (bm *bucketModel) ShortDescription() string {
	return fmt.Sprintf("bucket[%s]", bm.name)
}

func (bm *bucketModel) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("bucket")
	iw.AttrsWhere(bm)
	iw.TextAttr("name", bm.name)
	iw.NestedAttr("policy", bm.policy)
	iw.EndAttrs()
}

func (bm *bucketModel) Attach(doc corebottom.PolicyDocument) {
	newbm := &bucketModel{loc: bm.loc, storage: bm.storage, id: bm.id, testLogger: bm.testLogger, name: bm.name, cloud: bm.cloud, policy: doc}
	bm.storage.Bind(bm.id, newbm)
	bm.testLogger.Log("we need to attach policy with %d items to bucket %s\n", len(doc.Items()), bm.name)
}

func (eb *bucketModel) ObtainDest() corebottom.FileDest {
	return eb.cloud
}

var _ driverbottom.Describable = &bucketModel{}
var _ corebottom.PolicyAttacher = &bucketModel{}
var _ corebottom.DestHolder = &bucketModel{}
