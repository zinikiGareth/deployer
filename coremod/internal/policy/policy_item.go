package policy

import "ziniki.org/deployer/coremod/pkg/corebottom"

type policyItem struct {
	effect     string
	actions    []string
	resources  []string
	principals []corebottom.PolicyPrincipal
	more       map[string][]any
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

func (pi *policyItem) Principal(principal corebottom.PolicyPrincipal) {
	pi.principals = append(pi.principals, principal)
}

func (pi *policyItem) AMore(key string, value any) {
	if pi.more[key] == nil {
		pi.more[key] = []any{}
	}
	pi.more[key] = append(pi.more[key], value)
}

func (pi *policyItem) Actions() []string {
	return pi.actions
}

func (pi *policyItem) Resources() []string {
	return pi.resources
}

func (pi *policyItem) Principals() []corebottom.PolicyPrincipal {
	return pi.principals
}

func (pi *policyItem) More() map[string][]any {
	return pi.more
}
