package pluggable

import "ziniki.org/deployer/deployer/pkg/errors"

type Tools struct {
	Reporter   errors.ErrorRepI
	Register   Register
	Recall     Recall
	Resolver   Resolver
	Repository Repository
	Storage    RuntimeStorage
	Parser     ExprParser
}

func NewTools(reporter errors.ErrorRepI, register Register, recall Recall, repo Repository, storage RuntimeStorage) *Tools {
	return &Tools{Reporter: reporter, Register: register, Recall: recall, Repository: repo, Storage: storage}
}
