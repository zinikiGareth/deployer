package creator

import (
	"ziniki.org/deployer/deployer/internal/impl"
	"ziniki.org/deployer/deployer/pkg/deployer"
)

func NewDeployer() deployer.Deployer {
	return impl.NewDeployer()
}
