package external

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Tools struct {
	*driverbottom.CoreTools
	Options *Options
}

type Options struct {
	TearDown bool
}

func NewTools(core *driverbottom.CoreTools, options *Options) *Tools {
	return &Tools{CoreTools: core, Options: options}
}
