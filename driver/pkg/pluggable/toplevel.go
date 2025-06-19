package pluggable

import "fmt"

type TopLevelForm interface {
	Describable
	Resolvable
	Name() SymbolName
}

type TargetThing interface {
	fmt.Stringer
	Describable
	Resolvable
	BuildModel()
	UpdateReality()
	TearDown()
}
