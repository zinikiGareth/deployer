package registry

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Registry struct {
	verbs   map[string]pluggable.Action
	nouns   map[string]pluggable.Noun
	drivers map[string]any
}

func (r *Registry) RegisterVerb(verb string, action pluggable.Action) {
	r.verbs[verb] = action
}

func (r *Registry) RegisterNoun(noun string, item pluggable.Noun) {
	r.nouns[noun] = item
}

func (r *Registry) ProvideDriver(s string, env any) {
	r.drivers[s] = env
}

func (r *Registry) ObtainDriver(s string) any {
	return r.drivers[s]
}

func (r *Registry) FindVerb(verb string) pluggable.Action {
	return r.verbs[verb]
}

func (r *Registry) FindNoun(noun string) pluggable.Noun {
	return r.nouns[noun]
}

func NewRegistry() *Registry {
	return &Registry{verbs: make(map[string]pluggable.Action), nouns: make(map[string]pluggable.Noun), drivers: make(map[string]any)}
}
