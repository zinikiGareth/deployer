package registry

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Registry struct {
	verbs map[string]pluggable.Action
}

func (r *Registry) RegisterVerb(verb string, action pluggable.Action) {
	r.verbs[verb] = action
}

func (r *Registry) FindVerb(verb string) pluggable.Action {
	return r.verbs[verb]
}

func NewRegistry() *Registry {
	return &Registry{verbs: make(map[string]pluggable.Action)}
}
