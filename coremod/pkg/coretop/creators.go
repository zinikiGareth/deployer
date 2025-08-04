package coretop

import (
	"fmt"

	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func MakeGetCoinMethod(loc *errorsink.Location, coin corebottom.CoinId) driverbottom.Expr {
	return basic.MakeGetCoinMethod(loc, coin)
}

func NewPolicyDocument(loc *errorsink.Location) corebottom.PolicyDocument {
	return policy.NewPolicyDocument(loc)
}

func NewPolicyActionList(loc *errorsink.Location) corebottom.PolicyActionList {
	return policy.NewPolicyActionList(loc)
}

func NewPrincipal(k, v string) corebottom.PolicyPrincipal {
	return policy.NewPrincipal(k, v)
}

func NewPolicyStatementInterpreter(tools *driverbottom.CoreTools, scope driverbottom.Scope, parent driverbottom.PropertyParent, prop driverbottom.Identifier, tokens []driverbottom.Token) driverbottom.Interpreter {
	pd := NewPolicyActionList(prop.Loc())
	parent.AddProperty(prop, pd)
	return drivertop.NewVerbCommandInterpreter(tools, attachToPolicy{list: pd}, "policy-statements", false)
}

type attachToPolicy struct {
	list corebottom.PolicyActionList
}

func (a attachToPolicy) Attach(item any) error {
	pra, ok := item.(corebottom.PolicyRuleAction)
	if !ok {
		return fmt.Errorf("cannot attach %T to PolicyActionList, not a PolicyRuleAction", item)
	}

	a.list.Add(pra)
	return nil
}

func (a attachToPolicy) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	panic("this should not be able to happen in this context")
}
