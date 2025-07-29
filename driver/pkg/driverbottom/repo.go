package driverbottom

import "ziniki.org/deployer/driver/pkg/errorsink"

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

	ResolveAll(tools *CoreTools) bool
	FindTop(name SymbolName) TopLevelForm
	TopScope() Scope
	CurrentScope() Scope

	GetDefinition(id SymbolName) Describable
}

type Resolver interface {
	Resolve(scope Scope, name Identifier) any
	ErrorAtf(loc *errorsink.Location, fmt string, opts ...any)
}

type Resolvable interface {
	Resolve(r Resolver) BindingRequirement
}

func (b BindingRequirement) Merge(other BindingRequirement) BindingRequirement {
	if b == ERROR_OCCURRED || other == ERROR_OCCURRED {
		return ERROR_OCCURRED
	}
	return other // is it more complicated than this?
}