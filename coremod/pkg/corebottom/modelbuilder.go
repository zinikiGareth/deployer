package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type ModelBuilder interface {
	driverbottom.Describable
	Findable

	DetermineDesiredState(pres driverbottom.ValuePresenter)
}
