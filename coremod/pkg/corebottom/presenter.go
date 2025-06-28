package corebottom

type ValuePresenter interface {
	NotFound()
	Present(value any)
}
