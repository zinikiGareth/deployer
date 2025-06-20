package corebottom

const (
	BUILD_MODEL_MODE int = iota
	UPDATE_REALITY_MODE
)

type Deployer interface {
	// driverbottom.Driver
	Deploy(targets ...string) error
	ObtainTools() *Tools // for the benefit of plugins
}
