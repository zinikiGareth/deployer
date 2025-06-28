package policy

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type policyAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location

	exprs []driverbottom.Expr
}

func (pca *policyAction) Loc() *errorsink.Location {
	return pca.loc
}

func (pca *policyAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("policyAction")
	w.AttrsWhere(pca)
	w.ListAttr("exprs")
	for _, e := range pca.exprs {
		e.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (pca *policyAction) ShortDescription() string {
	return fmt.Sprintf("policyAction[%d]", len(pca.exprs))
}

func (pca *policyAction) Completed() {
}

func (pca *policyAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	for _, a := range pca.exprs {
		a.Resolve(r)
	}
	return driverbottom.MAY_BE_BOUND
}

func (pca *policyAction) DetermineInitialState(pres corebottom.ValuePresenter) {
}

func (pca *policyAction) DetermineDesiredState(pres corebottom.ValuePresenter) {
}

func (pca *policyAction) ApplyTo(pi corebottom.PolicyEffect) {
	for _, a := range pca.exprs {
		a1, ok := pca.tools.Storage.EvalAsStringer(a)
		if !ok {
			panic("not a stringer")
		}
		pi.Action(a1.String())
	}
}

var _ UpdatePolicyAllowAction = &policyAction{}
var _ corebottom.ModelBuilder = &policyAction{}
