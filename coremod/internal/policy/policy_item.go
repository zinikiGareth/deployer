package policy

type policyItem struct {
	effect    string
	actions   []string
	resources []string
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

func (pi *policyItem) Actions() []string {
	return pi.actions
}

func (pi *policyItem) Resources() []string {
	return pi.resources
}
