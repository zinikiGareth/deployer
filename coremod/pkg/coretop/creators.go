package coretop

import (
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func NewPolicyDocument(loc *errorsink.Location) corebottom.PolicyDocument {
	return policy.NewPolicyDocument(loc)
}
