package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type Tools struct {
	Reporter errors.ErrorRepI
	Recall   Recall
	Resolver Resolver
}

func NewTools(reporter errors.ErrorRepI) *Tools {
	return &Tools{Reporter: reporter}
}
