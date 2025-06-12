package coremod

import (
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/errorsink"
)

func NewPolicyDocument(loc *errorsink.Location, name string) external.PolicyDocument {
	return policy.NewPolicyDocument(loc, name)
}
