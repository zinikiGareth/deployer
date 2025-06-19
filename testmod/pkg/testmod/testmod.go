package testmod

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/deployer"
	"ziniki.org/deployer/testmod/internal/blob"
	"ziniki.org/deployer/testmod/internal/testS3"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDriver(deployer deployer.Driver) error {
	if testRunner != nil {
		eh := testRunner.ErrorHandlerFor("log")
		eh.WriteMsg("Installing things from testmod\n")
	}
	toolsI := deployer.ObtainCoreTools().RetrieveOther("coremod")
	if toolsI == nil {
		panic("must load coremod first")
	}
	tools := toolsI.(*external.Tools)

	tools.Register.ProvideDriver("testS3.TestAwsEnv", &testS3.TestAwsEnv{})

	tools.Register.Register("target", "test.assertBucketHas", testS3.NewAssertBucketHandler(tools))
	tools.Register.Register("target", "blob", blob.NewBlobCommandHandler(tools))

	tools.Register.Register("blank", "test.S3.Bucket", &testS3.BucketBlank{})
	return nil
}
