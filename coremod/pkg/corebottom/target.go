package corebottom

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type Target interface {
	fmt.Stringer
	driverbottom.Describable
	driverbottom.Resolvable
	BuildModel()
	UpdateReality()
	TearDown()
}
