package policy

import (
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type PolicyCondAction struct {
	tools   *pluggable.Tools
	loc     *errorsink.Location
	actions []pluggable.ModelBuilder

	test  pluggable.Expr
	left  pluggable.Expr
	right pluggable.Expr
}

func (pca *PolicyCondAction) Loc() *errorsink.Location {
	return pca.loc
}

func (pca *PolicyCondAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("PolicyCondAction")
	w.AttrsWhere(pca)
	w.NestedAttr("test", pca.test)
	w.NestedAttr("left", pca.left)
	w.NestedAttr("right", pca.right)
	w.ListAttr("actions")
	for _, a := range pca.actions {
		a.DumpTo(w)
	}
	w.EndList()
	w.EndAttrs()
}

func (pca *PolicyCondAction) ShortDescription() string {
	return "PolicyCond[" + pca.test.ShortDescription() + "]"
}

func (pca *PolicyCondAction) Completed() {
}

func (pca *PolicyCondAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	return pluggable.MAY_BE_BOUND
}

func (pca *PolicyCondAction) BuildModel(pres pluggable.ValuePresenter) {
}

func (pca *PolicyCondAction) ApplyTo(pi *policyItem) {
}

var _ UpdatePolicyAllowAction = &PolicyCondAction{}
