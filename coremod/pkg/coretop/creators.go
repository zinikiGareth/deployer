package coretop

import (
	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
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
