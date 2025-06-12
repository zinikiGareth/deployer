package testmod

import (
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/testmod/internal/blob"
	"ziniki.org/deployer/testmod/internal/testS3"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	if testRunner != nil {
		eh := testRunner.ErrorHandlerFor("log")
		eh.WriteMsg("Installing things from testmod\n")
	}
	tools := deployer.ObtainTools()
	tools.Register.ProvideDriver("testS3.TestAwsEnv", &testS3.TestAwsEnv{})

	tools.Register.Register("target", "test.assertBucketHas", testS3.NewAssertBucketHandler(tools))
	tools.Register.Register("target", "blob", blob.NewBlobCommandHandler(tools))

	tools.Register.Register("blank", "test.S3.Bucket", &testS3.BucketBlank{})
	return nil
}
