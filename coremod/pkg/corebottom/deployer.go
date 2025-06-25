package corebottom

const (
	DETERMINE_INITIAL_MODE int = iota
	DETERMINE_DESIRED_MODE
	UPDATE_REALITY_MODE
	TEARDOWN_MODE
)

type Deployer interface {
	// driverbottom.Driver
	Deploy(targets ...string) error
	ObtainTools() *Tools // for the benefit of plugins
}
