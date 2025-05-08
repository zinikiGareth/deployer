package ensurable

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Ensurable interface {
	pluggable.Describable
	Prepare(pres pluggable.ValuePresenter)
	Execute()
}
