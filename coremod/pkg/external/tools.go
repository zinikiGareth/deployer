package external

import (
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type Tools struct {
	*pluggable.CoreTools
	Options *Options
}

type Options struct {
	TearDown bool
}

func NewTools(core *pluggable.CoreTools, options *Options) *Tools {
	return &Tools{CoreTools: core, Options: options}
}
