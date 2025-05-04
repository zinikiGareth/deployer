package registry

import (
	"log"
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Registry struct {
	impls   map[reflect.Type]map[string]any
	nouns   map[string]pluggable.Noun
	funcs   map[string]pluggable.Function
	drivers map[string]any
}

func (r *Registry) Register(what reflect.Type, called string, impl any) {
	if !reflect.TypeOf(impl).Implements(what) {
		log.Fatalf("%v is not a %v", impl, what)
	}
	m := r.impls[what]
	if m == nil {
		m = make(map[string]any)
		r.impls[what] = m
	}
	m[called] = impl
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

func (r *Registry) Find(ty reflect.Type, called string) any {
	m := r.impls[ty]
	if m == nil {
		log.Fatalf("no verbs have been bound of type %v", ty)
	}
	ret := m[called]
	if ret == nil {
		log.Fatalf("there is no verb %s of type %v", called, ty)
	}
	return ret
}

func (r *Registry) FindFunc(verb string) pluggable.Function {
	return r.funcs[verb]
}

func (r *Registry) FindNoun(noun string) pluggable.Noun {
	return r.nouns[noun]
}

func NewRegistry() *Registry {
	return &Registry{impls: make(map[reflect.Type]map[string]any), funcs: make(map[string]pluggable.Function), nouns: make(map[string]pluggable.Noun), drivers: make(map[string]any)}
}
