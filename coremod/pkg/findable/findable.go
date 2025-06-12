package findable

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Findable interface {
	pluggable.Describable
	BuildModel(pres pluggable.ValuePresenter)
	UpdateReality()
}
