package driverbottom

type TopLevelForm interface {
	Describable
	Resolvable
	Name() SymbolName
}
