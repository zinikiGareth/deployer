package policy

import "ziniki.org/deployer/coremod/pkg/external"

type policyItem struct {
	effect     string
	actions    []string
	resources  []string
	principals []external.PolicyPrincipal
}

func (pi *policyItem) Effect() string {
	return pi.effect
}

func (pi *policyItem) Action(action string) {
	pi.actions = append(pi.actions, action)
}

func (pi *policyItem) Resource(resource string) {
	pi.resources = append(pi.resources, resource)
}

func (pi *policyItem) Principal(principal external.PolicyPrincipal) {
	pi.principals = append(pi.principals, principal)
}

func (pi *policyItem) Actions() []string {
	return pi.actions
}

func (pi *policyItem) Resources() []string {
	return pi.resources
}

func (pi *policyItem) Principals() []external.PolicyPrincipal {
	return pi.principals
}
