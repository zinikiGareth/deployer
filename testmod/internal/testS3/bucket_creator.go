package testS3

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type bucketCreator struct {
	name     string
	assignTo pluggable.Identifier
	tools    *pluggable.Tools
}

// This is called during the "Prepare" phase
func (b *bucketCreator) Ensure() {
	tmp := b.tools.Recall.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}

	tmp = b.tools.Recall.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	eb := &ensureBucket{env: testAwsEnv, bucket: b}
	testLogger.Log("ensuring bucket exists action %s\n", eb.String())
	if b.assignTo != nil {
		b.tools.Storage.Bind(pluggable.SymbolName(b.assignTo.Id()), eb)
	}
}
