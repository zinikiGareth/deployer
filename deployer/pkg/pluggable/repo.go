package pluggable

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errorsink"
)

type RepositoryTraverser interface {
	Visit(who SymbolName, what Describable)
}

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(who SymbolName, is Describable)
	TopLevel(is TargetThing)
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)

	ResolveAll(tools *Tools)
	FindTarget(name SymbolName) TargetThing
}

type TargetThing interface {
	fmt.Stringer
	Resolve(r Resolver)
	Prepare()
	Execute()
	TearDown()
}

type Resolver interface {
	Resolve(name Identifier) Describable
}

type Locatable interface {
	Loc() *errorsink.Location
}

type Describable interface {
	Locatable

	// ShortDescription enables clients to describe what they are pointing to in a unique way
	ShortDescription() string

	// DumpTo renders the whole of the text of the definition in a reproducible and unique, but not necessarily parseable form
	DumpTo(to IndentWriter)
}
