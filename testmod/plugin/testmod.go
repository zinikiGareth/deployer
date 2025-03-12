package main

import (
	"ziniki.org/deployer/deployer/pkg/deployer"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	eh := testRunner.ErrorHandlerFor("log")
	eh.WriteMsg("Need to install things from testmod\n")
	return nil
}
