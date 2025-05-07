package pluggable

import (
	"ziniki.org/deployer/deployer/pkg/errors"
)

type RepositoryTraverser interface {
	Visit(who SymbolName, what Definition)
}

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(who SymbolName, is Definition)
	TopLevel(is Definition)
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)

	ResolveAll(tools *Tools)
	FindTarget(name SymbolName) TargetThing
}

type Resolver interface {
	Resolve(name Identifier) Noun
}

type Locatable interface {
	Loc() *errors.Location
}

type ContainingContext interface {
	Add(entry Definition)
}
