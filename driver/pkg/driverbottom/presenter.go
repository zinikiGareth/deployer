package driverbottom

type ValuePresenter interface {
	NotFound()
	Present(value any)
	Unchanged()
}
