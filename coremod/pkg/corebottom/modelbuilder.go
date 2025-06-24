package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type ModelBuilder interface {
	driverbottom.Describable
	driverbottom.Resolvable

	DetermineDesiredState(pres driverbottom.ValuePresenter)
}
