package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type Tools struct {
	Reporter   errors.ErrorRepI
	Recall     Recall
	Resolver   Resolver
	Repository Repository
	Parser     ExprParser
}

func NewTools(reporter errors.ErrorRepI, repo Repository) *Tools {
	return &Tools{Reporter: reporter, Repository: repo}
}
