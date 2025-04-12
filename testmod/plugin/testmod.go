package main

import (
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/testmod/internal/testS3"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	// eh := testRunner.ErrorHandlerFor("log")
	// eh.WriteMsg("Need to install things from testmod\n")
	register := deployer.ObtainRegister()
	register.RegisterNoun("test.S3.Bucket", &testS3.BucketNoun{})
	return nil
}
