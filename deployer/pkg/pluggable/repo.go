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
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)

	ResolveAll(sink errors.ErrorSink)
}

type Resolver interface {
	Resolve(name Identifier) Definition
}

type Locatable interface {
	Loc() Location
}

type ContainingContext interface {
	Add(entry Definition)
}
