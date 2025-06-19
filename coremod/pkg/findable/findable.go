package findable

import "ziniki.org/deployer/driver/pkg/pluggable"

type Findable interface {
	pluggable.Describable
	BuildModel(pres pluggable.ValuePresenter)
	UpdateReality()
}
