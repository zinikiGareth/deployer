package main

import (
	"ziniki.org/deployer/coremod/target"
	"ziniki.org/deployer/deployer/pkg/deployer"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	if testRunner != nil {
		eh := testRunner.ErrorHandlerFor("log")
		eh.WriteMsg("Installing things from coremod\n")
	}
	register := deployer.ObtainRegister()
	register.RegisterAction("target", &target.CoreTargetVerb{})
	return nil
}
