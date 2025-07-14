package corebottom

import "ziniki.org/deployer/driver/pkg/errorsink"

type ValuePresenter interface {
	NotFound()
	Present(value any)
	WantDestruction(loc *errorsink.Location)
}
