package driverbottom

type BindingRequirement int

const (
	MUST_BE_BOUND BindingRequirement = iota
	MAY_BE_BOUND
	NO_VALUE
	ERROR_OCCURRED
)

type RepositoryTraverser interface {
	Visit(who SymbolName, what Describable)
}

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(who SymbolName, is Describable)
	TopLevel(tlf TopLevelForm)
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)

	ResolveAll(tools *CoreTools)
	FindTop(name SymbolName) TopLevelForm

	GetDefinition(id SymbolName) Describable
}

type Resolver interface {
	Resolve(name Identifier) any
}

type Resolvable interface {
	Resolve(r Resolver) BindingRequirement
}
