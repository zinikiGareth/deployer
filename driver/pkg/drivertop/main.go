package drivertop

import (
	"ziniki.org/deployer/driver/internal/impl"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func RunDeployer(args []string) int {
	return impl.RunDeployer(args)
}

func RunDeployerWithConfig(config func(driverbottom.Driver) error, args []string) int {
	return impl.RunDeployerWithConfig(config, args)
}

func Usage() {
	impl.Usage()
}
