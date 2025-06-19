package coretop

import (
	"ziniki.org/deployer/coremod/internal/deployer"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func NewDeployer(driver driverbottom.Driver, tools *corebottom.Tools) corebottom.Deployer {
	return deployer.NewDeployer(driver, tools)
}
