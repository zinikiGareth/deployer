package findable

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Findable interface {
	pluggable.Describable
	Prepare(pres pluggable.ValuePresenter)
	Execute()
}
