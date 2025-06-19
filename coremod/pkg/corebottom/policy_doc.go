package corebottom

type PolicyDocument interface {
	Item(s string) PolicyEffect
	Items() []PolicyEffect
}

type PolicyEffect interface {
	Effect() string
	Actions() []string
	Resources() []string
	Principals() []PolicyPrincipal
	More() map[string][]any

	Action(s string)
	Resource(s string)
	Principal(s PolicyPrincipal)
	AMore(key string, value any)
}

type PolicyPrincipal interface {
	Key() string
	Value() string
}

type PolicyAttacher interface {
	Attach(policy PolicyDocument)
}
