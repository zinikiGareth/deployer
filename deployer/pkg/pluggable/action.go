package pluggable

type Action interface {
	Describable

	// Resolve asks the definition to examine all of its structure and ask for resolution of any unresolved names
	Resolve(r Resolver, b Binder)

	Prepare(pres ValuePresenter)
	Execute()
}

type ValuePresenter interface {
	Present(value any)
}
