package driverbottom

type TopLevelForm interface {
	Describable
	Resolvable
	HasScope
	Name() SymbolName
}

type Scope interface {
	IntroduceSymbol(who SymbolName, is Describable) error
	Traverse(lsnr RepositoryTraverser)
	FindDefinition(name SymbolName) any
}

type HasScope interface {
	Scope() Scope
}
