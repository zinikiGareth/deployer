package main

import (
	"ziniki.org/deployer/coremod/pkg/coremod"
	"ziniki.org/deployer/deployer/pkg/deployer"
)

func ProvideTestRunner(runner deployer.TestRunner) error {
	return coremod.ProvideTestRunner(runner)
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	return coremod.RegisterWithDeployer(deployer)
}
