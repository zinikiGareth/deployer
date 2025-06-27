package testS3

import (
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type bucketModel struct {
	storage    driverbottom.RuntimeStorage
	id         corebottom.CoinId
	testLogger testhelpers.TestStepLogger
	name       string
	cloud      *BucketCloud

	policy corebottom.PolicyDocument
}

func (bm *bucketModel) Attach(doc corebottom.PolicyDocument) {
	log.Printf("bm.id == %s\n", bm.id.VarName().Id())
	newbm := &bucketModel{storage: bm.storage, id: bm.id, testLogger: bm.testLogger, name: bm.name, cloud: bm.cloud, policy: doc}
	bm.storage.Bind(bm.id, newbm)
	bm.testLogger.Log("we need to attach policy with %d items to bucket %s\n", len(doc.Items()), bm.name)
}

func (eb *bucketModel) ObtainDest() corebottom.FileDest {
	return eb.cloud
}

var _ corebottom.PolicyAttacher = &bucketModel{}
var _ corebottom.DestHolder = &bucketModel{}
