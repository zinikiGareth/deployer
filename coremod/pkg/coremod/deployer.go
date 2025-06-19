package coremod

import (
	deployer2 "ziniki.org/deployer/coremod/internal/deployer"
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/deployer"
)

func NewDeployer(driver deployer.Driver, tools *external.Tools) external.Deployer {
	return deployer2.NewDeployer(driver, tools)
}
