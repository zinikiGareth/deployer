package main

import (
	"ziniki.org/deployer/coremod/pkg/coretop"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func ProvideTestRunner(runner driverbottom.TestRunner) error {
	return coretop.ProvideTestRunner(runner)
}

func RegisterWithDriver(deployer driverbottom.Driver) error {
	return coretop.RegisterWithDriver(deployer)
}
