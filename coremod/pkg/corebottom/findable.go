package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Findable interface {
	driverbottom.Describable
	driverbottom.Resolvable

	DetermineInitialState(pres driverbottom.ValuePresenter)
}
