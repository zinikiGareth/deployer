package ensurable

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Ensurable interface {
	driverbottom.Describable
	BuildModel(pres driverbottom.ValuePresenter)
	UpdateReality()
	TearDown()
}
