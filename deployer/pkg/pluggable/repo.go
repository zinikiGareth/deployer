package pluggable

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(who SymbolName, is Definition)
	AddSymbolListener(lsnr SymbolListener)
}

type Locatable interface {
	Loc() Location
}

type ContainingContext interface {
	Add(entry Locatable)
}
