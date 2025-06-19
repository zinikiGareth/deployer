package policy

import (
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type PolicyCondAction struct {
	tools *external.Tools
	loc   *errorsink.Location

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
	w.EndAttrs()
}

func (pca *PolicyCondAction) ShortDescription() string {
	return "PolicyCond[" + pca.test.ShortDescription() + "]"
}

func (pca *PolicyCondAction) Completed() {
}

func (pca *PolicyCondAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	pca.test.Resolve(r)
	pca.left.Resolve(r)
	pca.right.Resolve(r)
	return pluggable.MAY_BE_BOUND
}

func (pca *PolicyCondAction) BuildModel(pres pluggable.ValuePresenter) {
}

func (pca *PolicyCondAction) ApplyTo(pi external.PolicyEffect) {
	test := pca.tools.Storage.EvalAsString(pca.test)
	left := pca.tools.Storage.EvalAsString(pca.left)
	right := pca.tools.Storage.Eval(pca.right)
	expr := map[string]any{}
	expr[left] = right
	cond := map[string]any{}
	cond[test] = expr
	pi.AMore("Condition", cond)
}

var _ UpdatePolicyAllowAction = &PolicyCondAction{}
