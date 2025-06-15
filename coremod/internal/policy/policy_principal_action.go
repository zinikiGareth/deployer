package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type policyPrincipalAction struct {
	tools *pluggable.Tools
	loc   *errorsink.Location

	ofType pluggable.Expr
	id     pluggable.Expr
}

func (pca *policyPrincipalAction) Loc() *errorsink.Location {
	return pca.loc
}

func (pca *policyPrincipalAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("policyPrincipalAction")
	w.AttrsWhere(pca)
	w.NestedAttr("type", pca.ofType)
	w.NestedAttr("id", pca.id)
	w.EndAttrs()
}

func (pca *policyPrincipalAction) ShortDescription() string {
	return "policyPrincipal[" + pca.ofType.ShortDescription() + ":" + pca.id.ShortDescription() + "]"
}

func (pca *policyPrincipalAction) Completed() {
}

func (pca *policyPrincipalAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	pca.ofType.Resolve(r)
	pca.id.Resolve(r)
	return pluggable.MAY_BE_BOUND
}

func (pca *policyPrincipalAction) BuildModel(pres pluggable.ValuePresenter) {
}

func (pca *policyPrincipalAction) ApplyTo(pi *policyItem) {
	pi.principals = append(pi.principals, NewPrincipal(pca.tools.Storage.EvalAsString(pca.ofType), pca.tools.Storage.EvalAsString(pca.id)))
}

var _ UpdatePolicyAllowAction = &policyPrincipalAction{}
