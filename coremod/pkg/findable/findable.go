package findable

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Findable interface {
	driverbottom.Describable
	BuildModel(pres driverbottom.ValuePresenter)
	UpdateReality()
}
