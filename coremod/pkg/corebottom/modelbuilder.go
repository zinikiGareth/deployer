package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type ModelBuilder interface {
	driverbottom.Describable
	driverbottom.Resolvable

	BuildModel(pres driverbottom.ValuePresenter)
}
