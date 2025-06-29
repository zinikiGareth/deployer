package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Action interface {
	driverbottom.Describable
	driverbottom.Resolvable
}
