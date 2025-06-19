package external

type Deployer interface {
	// driverbottom.Driver
	Deploy(targets ...string) error
	ObtainTools() *Tools // for the benefit of plugins
}
