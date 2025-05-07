package testmod

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/pluggable"
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

	tools.Register.Register(reflect.TypeFor[pluggable.TargetCommand](), "test.assertBucketHas", testS3.NewAssertBucketHandler(tools))

	tools.Register.Register(reflect.TypeFor[pluggable.Blank](), "test.S3.Bucket", &testS3.BucketBlank{})
	return nil
}
