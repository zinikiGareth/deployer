package ensurable

import "ziniki.org/deployer/driver/pkg/pluggable"

type Ensurable interface {
	pluggable.Describable
	BuildModel(pres pluggable.ValuePresenter)
	UpdateReality()
	TearDown()
}
