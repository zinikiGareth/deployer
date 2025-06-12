package external

type PolicyDocument interface {
	Name() string
	Item(s string) PolicyEffect
	Items() []PolicyEffect
}

type PolicyEffect interface {
	Effect() string
	Actions() []string
	Resources() []string

	Action(s string)
	Resource(s string)
}
