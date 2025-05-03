package registry

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Registry struct {
	actions map[string]pluggable.Action
	nouns   map[string]pluggable.Noun
	funcs   map[string]pluggable.Function
	drivers map[string]any
}

func (r *Registry) RegisterAction(verb string, action pluggable.Action) {
	r.actions[verb] = action
}

func (r *Registry) RegisterFunc(verb string, fn pluggable.Function) {
	r.funcs[verb] = fn
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

func (r *Registry) FindAction(verb string) pluggable.Action {
	return r.actions[verb]
}

func (r *Registry) FindFunc(verb string) pluggable.Function {
	return r.funcs[verb]
}

func (r *Registry) FindNoun(noun string) pluggable.Noun {
	return r.nouns[noun]
}

func NewRegistry() *Registry {
	return &Registry{actions: make(map[string]pluggable.Action), funcs: make(map[string]pluggable.Function), nouns: make(map[string]pluggable.Noun), drivers: make(map[string]any)}
}
