package policy

import (
	"fmt"
	"log"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type AttachPolicyAction struct {
	tools  *corebottom.Tools
	loc    *errorsink.Location
	to     driverbottom.Expr
	policy driverbottom.Expr

	attachTo     corebottom.PolicyAttacher
	actualPolicy corebottom.PolicyDocument
}

func (apa *AttachPolicyAction) Loc() *errorsink.Location {
	return apa.loc
}

func (apa *AttachPolicyAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("AttachPolicyAction")
	w.AttrsWhere(apa)
	w.NestedAttr("to", apa.to)
	w.NestedAttr("policy", apa.policy)
	w.EndAttrs()
}

func (apa *AttachPolicyAction) ShortDescription() string {
	return fmt.Sprintf("AttachPolicy[%s: %s]", apa.to.ShortDescription(), apa.policy.ShortDescription())
}

func (apa *AttachPolicyAction) AddProperty(name driverbottom.Identifier, value driverbottom.Expr) {
}

func (apa *AttachPolicyAction) AddAdverb(adv driverbottom.Adverb, tokens []driverbottom.Token) driverbottom.Interpreter {
	return drivertop.NewDisallowInnerScope(apa.tools.CoreTools)
}

func (apa *AttachPolicyAction) Completed() {
}

func (apa *AttachPolicyAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	apa.to.Resolve(r)
	apa.policy.Resolve(r)
	return driverbottom.NO_VALUE
}

func (apa *AttachPolicyAction) DetermineInitialState(pres corebottom.ValuePresenter) {
}

func (apa *AttachPolicyAction) DetermineDesiredState(pres corebottom.ValuePresenter) {
	attachTo := apa.to.Eval(apa.tools.Storage)
	policy := apa.policy.Eval(apa.tools.Storage)

	attacher, ok := attachTo.(corebottom.PolicyAttacher)
	if !ok {
		log.Fatalf("cannot attach things to %T", attachTo)
	}
	isPolicy, ok := policy.(corebottom.PolicyDocument)
	if !ok {
		log.Fatalf("%T was not a policy", policy)
	}
	apa.attachTo = attacher
	apa.actualPolicy = isPolicy
}

func (apa *AttachPolicyAction) ShouldDestroy() bool {
	return false
}

func (apa *AttachPolicyAction) UpdateReality() {
	apa.attachTo.Attach(apa.actualPolicy)
}

func (apa *AttachPolicyAction) TearDown() {
	// ea.ens.TearDown()
}

var _ corebottom.RealityShifter = &AttachPolicyAction{}
