package pluggable

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(where Location, what SymbolType, who SymbolName)
	AddSymbolListener(lsnr SymbolListener)
}
