package pluggable

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
)

type CoreTools struct {
	Reporter   errorsink.ErrorRepI
	Register   Register
	Recall     Recall
	Resolver   Resolver
	Repository Repository
	Storage    RuntimeStorage
	Parser     ExprParser
	others     map[string]any
}

func (ct *CoreTools) StoreOther(name string, tools any) {
	ct.others[name] = tools
}

func (ct *CoreTools) RetrieveOther(name string) any {
	return ct.others[name]
}

func NewTools(reporter errorsink.ErrorRepI, register Register, recall Recall, repo Repository, storage RuntimeStorage) *CoreTools {
	return &CoreTools{Reporter: reporter, Register: register, Recall: recall, Repository: repo, Storage: storage}
}
