package main

import (
	"ziniki.org/deployer/coremod/pkg/coremod"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func ProvideTestRunner(runner driverbottom.TestRunner) error {
	return coremod.ProvideTestRunner(runner)
}

func RegisterWithDriver(deployer driverbottom.Driver) error {
	return coremod.RegisterWithDriver(deployer)
}
