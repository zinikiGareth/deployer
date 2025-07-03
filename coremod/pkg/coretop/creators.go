package coretop

import (
	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/methods"
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func MakeInvokeExpr(on driverbottom.Expr, call driverbottom.Identifier, args ...driverbottom.Expr) driverbottom.Expr {
	return methods.MakeInvokeExpr(on, call, args)
}

func MakeGetCoinMethod(loc *errorsink.Location, coin corebottom.CoinId) driverbottom.Expr {
	return basic.MakeGetCoinMethod(loc, coin)
}

func NewPolicyDocument(loc *errorsink.Location) corebottom.PolicyDocument {
	return policy.NewPolicyDocument(loc)
}
