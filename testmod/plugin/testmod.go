package main

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/testmod/pkg/testmod"
)

func ProvideTestRunner(runner driverbottom.TestRunner) error {
	return testmod.ProvideTestRunner(runner)
}

func RegisterWithDriver(deployer driverbottom.Driver) error {
	return testmod.RegisterWithDriver(deployer)
}
