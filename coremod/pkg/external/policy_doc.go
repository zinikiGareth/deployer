package external

type PolicyDocument interface {
	Item(s string) PolicyEffect
	Items() []PolicyEffect
}

type PolicyEffect interface {
	Effect() string
	Actions() []string
	Resources() []string
	Principals() []PolicyPrincipal

	Action(s string)
	Resource(s string)
	Principal(s PolicyPrincipal)
}

type PolicyPrincipal interface {
	Key() string
	Value() string
}
