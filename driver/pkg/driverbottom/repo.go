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
	AtLevel(level int) Scope
	TopLevel(tlf TopLevelForm)
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)

	ResolveAll(tools *CoreTools)
	FindTop(name SymbolName) TopLevelForm
	TopScope() Scope
	CurrentScope() Scope

	GetDefinition(id SymbolName) Describable
}

type Resolver interface {
	Resolve(scope Scope, name Identifier) any
}

type Resolvable interface {
	Resolve(r Resolver) BindingRequirement
}
