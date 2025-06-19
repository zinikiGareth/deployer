package policy

import (
	"fmt"
	"log"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/interpreters"
)

// The action is created by the handler.  It is added to a target.  It then takes on the rest of the work:
// resolution, preparation, execution

type AttachPolicyAction struct {
	tools  *external.Tools
	loc    *errorsink.Location
	to     driverbottom.Expr
	policy driverbottom.Expr

	attachTo     external.PolicyAttacher
	actualPolicy external.PolicyDocument
}

func (ea *AttachPolicyAction) Loc() *errorsink.Location {
	return ea.loc
}

func (ea *AttachPolicyAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro("AttachPolicyAction")
	w.AttrsWhere(ea)
	w.NestedAttr("to", ea.to)
	w.NestedAttr("policy", ea.policy)
	w.EndAttrs()
}

func (ea *AttachPolicyAction) ShortDescription() string {
	return fmt.Sprintf("AttachPolicy[%s: %s]", ea.to.ShortDescription(), ea.policy.ShortDescription())
}

func (ea *AttachPolicyAction) AddProperty(name driverbottom.Identifier, value driverbottom.Expr) {
}

func (ea *AttachPolicyAction) AddAdverb(adv driverbottom.Adverb, tokens []driverbottom.Token) driverbottom.Interpreter {
	/*
		if adv.Name() == "teardown" {
			if ea.teardown != nil {
				panic("duplicate teardown")
			}
			if len(tokens) != 1 {
				panic("invalid tokens")
			}
			ea.teardown = &MyTearDown{mode: tokens[0].(driverbottom.Identifier).Id()}

		}
	*/
	return interpreters.NewDisallowInnerScope(ea.tools.CoreTools)
}

func (ea *AttachPolicyAction) Completed() {
}

func (ea *AttachPolicyAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ea.to.Resolve(r)
	ea.policy.Resolve(r)
	return driverbottom.NO_VALUE
}

func (ea *AttachPolicyAction) BuildModel(pres driverbottom.ValuePresenter) {
	attachTo := ea.to.Eval(ea.tools.Storage)
	policy := ea.policy.Eval(ea.tools.Storage)

	attacher, ok := attachTo.(external.PolicyAttacher)
	if !ok {
		log.Fatalf("cannot attach things to %T", attachTo)
	}
	isPolicy, ok := policy.(external.PolicyDocument)
	if !ok {
		log.Fatalf("%T was not a policy", policy)
	}
	ea.attachTo = attacher
	ea.actualPolicy = isPolicy
}

func (ea *AttachPolicyAction) UpdateReality() {
	ea.attachTo.Attach(ea.actualPolicy)
}

func (ea *AttachPolicyAction) TearDown() {
	// ea.ens.TearDown()
}
