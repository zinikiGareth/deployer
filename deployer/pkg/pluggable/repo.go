package pluggable

import (
	"ziniki.org/deployer/deployer/pkg/errors"
)

type RepositoryTraverser interface {
	Visit(who SymbolName, what Action)
}

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(who SymbolName, is Action)
	TopLevel(is Action)
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)

	ResolveAll(tools *Tools)
	FindTarget(name SymbolName) TargetThing
}

type TargetThing interface {
	Prepare()
	Execute()
}

type Resolver interface {
	Resolve(name Identifier) Blank
}

type Locatable interface {
	Loc() *errors.Location
}

type ContainingContext interface {
	Add(entry Action)
}
