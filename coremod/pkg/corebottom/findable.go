package corebottom

type Findable interface {
	Action
	DetermineInitialState(pres ValuePresenter)
}
