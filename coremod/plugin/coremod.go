package main

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

func ProvideTestRunner(runner deployer.TestRunner) error {
	fmt.Printf("Have been given a testrunner %v, so configure for tests\n", runner)
	return nil
}

func RegisterWithDeployer(deployer *deployer.Deployer) error {
	fmt.Printf("Need to install things from coremod\n")
	return nil
}
