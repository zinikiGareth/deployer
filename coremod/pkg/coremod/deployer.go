package coremod

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/deployer"
)

type Deployer interface {
	deployer.Driver
	Deploy(targets ...string) error
	ObtainTools() *external.Tools // for the benefit of plugins
}
