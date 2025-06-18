package main

import (
	"ziniki.org/deployer/coremod/pkg/coremod"
	"ziniki.org/deployer/deployer/pkg/deployer"
)

func ProvideTestRunner(runner deployer.TestRunner) error {
	return coremod.ProvideTestRunner(runner)
}

func RegisterWithDriver(deployer deployer.Driver) error {
	return coremod.RegisterWithDriver(deployer)
}
