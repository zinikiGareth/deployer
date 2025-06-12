package pluggable

type BindingRequirement int

const (
	MUST_BE_BOUND BindingRequirement = iota
	MAY_BE_BOUND
	NO_VALUE
	ERROR_OCCURRED
)

type Action interface {
	Describable

	// Resolve asks the definition to examine all of its structure and ask for resolution of any unresolved names
	Resolve(r Resolver) BindingRequirement

	BuildModel(pres ValuePresenter)
}

type AndMakeItSo interface {
	Action

	UpdateReality()
	TearDown()
}

type ValuePresenter interface {
	Present(value any)
}
