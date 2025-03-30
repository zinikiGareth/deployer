package pluggable

type RepositoryTraverser interface {
	Visit(who SymbolName, what Definition)
}

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(who SymbolName, is Definition)
	AddSymbolListener(lsnr SymbolListener)
	Traverse(lsnr RepositoryTraverser)
}

type Locatable interface {
	Loc() Location
}

type ContainingContext interface {
	Add(entry Locatable)
}
