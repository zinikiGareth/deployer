package pluggable

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(where Location, what SymbolType, who SymbolName)
	AddSymbolListener(lsnr SymbolListener)
}

type Locatable interface {
	Loc() Location
}

type ContainingContext interface {
	Add(entry Locatable)
}
