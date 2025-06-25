package corebottom

const (
	RESOLVE_MODE int = iota
	DETERMINE_INITIAL_MODE
	DETERMINE_DESIRED_MODE
	UPDATE_REALITY_MODE
	TEARDOWN_MODE
)

type Deployer interface {
	Deploy(targets ...string) error
	ObtainTools() *Tools // for the benefit of plugins
}
