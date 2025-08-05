package corebottom

import "ziniki.org/deployer/driver/pkg/driverbottom"

type PolicyActionList interface {
	driverbottom.Expr
	Add(r PolicyRuleAction)
	Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement
	ApplyTo(doc PolicyDocument)
}

type PolicyDocument interface {
	driverbottom.Describable
	Item(s string) PolicyEffect
	Items() []PolicyEffect
}

type PolicyRuleAction interface {
	driverbottom.Describable
	Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement
	ApplyTo(doc PolicyDocument)
}

type PolicyEffect interface {
	Effect() string
	Actions() []string
	Resources() []string
	Principals() []PolicyPrincipal
	More() map[string][]any

	Action(s string)
	Resource(s string)
	Principal(s PolicyPrincipal)
	AMore(key string, value any)
}

type UpdatePolicyAllowAction interface {
	driverbottom.Describable
	Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement
	ApplyTo(doc PolicyEffect)
}

type PolicyPrincipal interface {
	Key() string
	Value() string
}

type PolicyAttacher interface {
	Attach(policy PolicyDocument)
}
