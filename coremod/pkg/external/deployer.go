package external

type Deployer interface {
	// deployer.Driver
	Deploy(targets ...string) error
	ObtainTools() *Tools // for the benefit of plugins
}
