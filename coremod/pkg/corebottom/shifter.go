package corebottom

type RealityShifter interface {
	ModelBuilder

	ShouldDestroy() bool
	UpdateReality()
	TearDown()
}
