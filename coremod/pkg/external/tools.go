package external

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Tools struct {
	pluggable.CoreTools
	Options *Options
}

type Options struct {
	TearDown bool
}

func NewTools(reporter errorsink.ErrorRepI, register pluggable.Register, recall pluggable.Recall, repo pluggable.Repository, storage pluggable.RuntimeStorage, options *Options) *Tools {
	return &Tools{CoreTools: pluggable.CoreTools{Reporter: reporter, Register: register, Recall: recall, Repository: repo, Storage: storage}, Options: options}
}
