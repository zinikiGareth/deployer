package pluggable

import "ziniki.org/deployer/deployer/pkg/errorsink"

type Tools struct {
	Reporter   errorsink.ErrorRepI
	Register   Register
	Recall     Recall
	Resolver   Resolver
	Repository Repository
	Storage    RuntimeStorage
	Parser     ExprParser
}

func NewTools(reporter errorsink.ErrorRepI, register Register, recall Recall, repo Repository, storage RuntimeStorage) *Tools {
	return &Tools{Reporter: reporter, Register: register, Recall: recall, Repository: repo, Storage: storage}
}
