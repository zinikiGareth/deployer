package coremod

import (
	"ziniki.org/deployer/coremod/internal/deployer"
	"ziniki.org/deployer/coremod/pkg/external"
)

func NewDeployer(tools *external.Tools) external.Deployer {
	return deployer.NewDeployer(tools)
}
