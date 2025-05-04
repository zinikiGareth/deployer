package testmod

import (
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/testmod/internal/testS3"
)

// var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	// testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	register := deployer.ObtainRegister()
	register.ProvideDriver("testS3.TestAwsEnv", &testS3.TestAwsEnv{})
	// register.Register(reflect.TypeFor[pluggable.Noun](), "test.S3.Bucket", &testS3.BucketNoun{})
	register.RegisterNoun("test.S3.Bucket", &testS3.BucketNoun{})
	return nil
}
