package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type policyPrincipalAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location

	ofType driverbottom.Expr
	id     driverbottom.Expr
}

func (pca *policyPrincipalAction) Loc() *errorsink.Location {
	return pca.loc
}

func (pca *policyPrincipalAction) DumpTo(w driverbottom.IndentWriter) {
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

func (pca *policyPrincipalAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	pca.ofType.Resolve(r)
	pca.id.Resolve(r)
	return driverbottom.MAY_BE_BOUND
}

func (pca *policyPrincipalAction) BuildModel(pres driverbottom.ValuePresenter) {
}

func (pca *policyPrincipalAction) ApplyTo(pi corebottom.PolicyEffect) {
	o, ok1 := pca.tools.Storage.EvalAsStringer(pca.ofType)
	i, ok2 := pca.tools.Storage.EvalAsStringer(pca.id)
	if !ok1 || !ok2 {
		panic("not ok")
	}
	pi.Principal(NewPrincipal(o.String(), i.String()))
}

var _ UpdatePolicyAllowAction = &policyPrincipalAction{}
