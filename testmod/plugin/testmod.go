package main

import (
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/testmod/pkg/testmod"
)

func ProvideTestRunner(runner deployer.TestRunner) error {
	return testmod.ProvideTestRunner(runner)
}

func RegisterWithDriver(deployer deployer.Driver) error {
	return testmod.RegisterWithDriver(deployer)
}
