package policy_test

import (
	"testing"

	"ziniki.org/deployer/coremod/internal/policy"
)

func TestTypeAssignment(t *testing.T) {
	var _ policy.PolicyRuleAction = &policy.PolicyAllowAction{}
}
