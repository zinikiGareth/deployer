package policy

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type PolicyCondAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location

	test  driverbottom.Expr
	left  driverbottom.Expr
	right driverbottom.Expr
}

func (pca *PolicyCondAction) Loc() *errorsink.Location {
	return pca.loc
}

func (pca *PolicyCondAction) DumpTo(w driverbottom.IndentWriter) {
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

func (pca *PolicyCondAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	pca.test.Resolve(r)
	pca.left.Resolve(r)
	pca.right.Resolve(r)
	return driverbottom.MAY_BE_BOUND
}

func (pca *PolicyCondAction) BuildModel(pres driverbottom.ValuePresenter) {
}

func (pca *PolicyCondAction) ApplyTo(pi corebottom.PolicyEffect) {
	test, ok1 := pca.tools.Storage.EvalAsStringer(pca.test)
	left, ok2 := pca.tools.Storage.EvalAsStringer(pca.left)
	if !ok1 || !ok2 {
		panic("you lose")
	}
	right := pca.tools.Storage.Eval(pca.right)
	expr := map[string]any{}
	expr[left.String()] = right
	cond := map[string]any{}
	cond[test.String()] = expr
	pi.AMore("Condition", cond)
}

var _ UpdatePolicyAllowAction = &PolicyCondAction{}
