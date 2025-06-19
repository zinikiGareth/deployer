package coremod

import (
	"ziniki.org/deployer/coremod/internal/deployer"
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func NewDeployer(driver driverbottom.Driver, tools *external.Tools) external.Deployer {
	return deployer.NewDeployer(driver, tools)
}
