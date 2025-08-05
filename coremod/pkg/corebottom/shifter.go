package corebottom

type RealityShifter interface {
	ModelBuilder

	ShouldDestroy() bool
	UpdateReality()
	TearDown()
}

type BasicShifter interface {
	DetermineInitialState(pres ValuePresenter)
	DetermineDesiredState(pres ValuePresenter)
	UpdateReality()
	TearDown()
}
