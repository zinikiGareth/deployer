package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Findable interface {
	driverbottom.Describable
	DetermineInitialState(pres driverbottom.ValuePresenter)
	DetermineDesiredState(pres driverbottom.ValuePresenter)
}
