package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type ModelBuilder interface {
	driverbottom.Describable
	driverbottom.Resolvable

	DetermineInitialState(pres driverbottom.ValuePresenter)
	DetermineDesiredState(pres driverbottom.ValuePresenter)
}
