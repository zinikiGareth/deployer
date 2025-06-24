package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Ensurable interface {
	driverbottom.Describable
	DetermineDesiredState(pres driverbottom.ValuePresenter)
	UpdateReality()
	TearDown()
}
