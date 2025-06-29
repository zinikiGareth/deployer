package corebottom

type ModelBuilder interface {
	Findable

	DetermineDesiredState(pres ValuePresenter)
}

type MemoryBuilder interface {
	Action

	Create(pres ValuePresenter)
}
