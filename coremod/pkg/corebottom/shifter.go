package corebottom

type RealityShifter interface {
	ModelBuilder

	UpdateReality()
	TearDown()
}
