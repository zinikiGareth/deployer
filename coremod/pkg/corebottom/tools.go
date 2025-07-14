package corebottom

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type Tools struct {
	*driverbottom.CoreTools
	Options *Options
}

type Options struct {
	TearDown bool
	Destroy  bool
}

func NewTools(core *driverbottom.CoreTools, options *Options) *Tools {
	if core.RetrieveOther("coremod") != nil {
		panic("NewTools called with coremod already bound")
	}
	ret := &Tools{CoreTools: core, Options: options}
	core.StoreOther("coremod", ret)
	return ret
}
