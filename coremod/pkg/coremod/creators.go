package coremod

import (
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func NewPolicyDocument(loc *errorsink.Location) external.PolicyDocument {
	return policy.NewPolicyDocument(loc)
}
