package testS3

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type bucketCreator struct {
	name     string
	assignTo pluggable.Identifier
}

// This is called during the "Prepare" phase
func (b *bucketCreator) Ensure(runtime pluggable.RuntimeStorage) pluggable.ExecuteAction {
	tmp := runtime.ObtainDriver("testS3.TestAwsEnv")
	testAwsEnv, ok := tmp.(*TestAwsEnv)
	if !ok {
		panic("could not cast env to TestAwsEnv")
	}

	tmp = runtime.ObtainDriver("testhelpers.TestStepLogger")
	testLogger, ok := tmp.(testhelpers.TestStepLogger)
	if !ok {
		panic("could not cast logger to TestStepLogger")
	}

	eb := &ensureBucket{env: testAwsEnv, bucket: b}
	testLogger.Log("ensuring bucket exists action %s\n", eb.String())
	if b.assignTo != nil {
		runtime.Bind(pluggable.SymbolName(b.assignTo.Id()), eb)
	}

	return eb
}
