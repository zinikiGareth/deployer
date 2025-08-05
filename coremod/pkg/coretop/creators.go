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

func NewPolicyAllowAction(tools *corebottom.Tools, loc *errorsink.Location, actions []driverbottom.Expr, resources []driverbottom.Expr, updates []corebottom.UpdatePolicyAllowAction) corebottom.PolicyRuleAction {
	return policy.NewPolicyAllowAction(tools, loc, actions, resources, updates)
}

func NewPolicyPrincipalAction(tools *corebottom.Tools, loc *errorsink.Location, ofType driverbottom.Expr, id driverbottom.Expr) corebottom.UpdatePolicyAllowAction {
	return policy.NewPolicyPrincipalAction(tools, loc, ofType, id)
}

func NewPrincipal(k, v string) corebottom.PolicyPrincipal {
	return policy.NewPrincipal(k, v)
}

func NewPolicyStatementInterpreter(tools *driverbottom.CoreTools, scope driverbottom.Scope, parent driverbottom.PropertyParent, prop driverbottom.Identifier, tokens []driverbottom.Token) driverbottom.Interpreter {
	pd := NewPolicyActionList(prop.Loc())
	parent.AddProperty(prop, pd)
	return drivertop.NewVerbCommandInterpreter(tools, attachToPolicy{list: pd}, "policy-statements", false)
}

func NewCoinPresenter(storage driverbottom.RuntimeStorage, coinId corebottom.CoinId, pres corebottom.ValuePresenter) corebottom.ValuePresenter {
	return basic.NewCoinPresenter(storage, coinId, pres)
}

func NewDummyPresenter() corebottom.ValuePresenter {
	return basic.NewDummyPresenter()
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
