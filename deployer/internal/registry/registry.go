package registry

import (
	"log"
	"reflect"
)

type Registry struct {
	impls   map[reflect.Type]map[string]any
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

func (r *Registry) Find(ty reflect.Type, called string) any {
	m := r.impls[ty]
	if m == nil {
		return nil
	}
	return m[called]
}

func (r *Registry) ProvideDriver(s string, env any) {
	r.drivers[s] = env
}

func (r *Registry) ObtainDriver(s string) any {
	return r.drivers[s]
}

func NewRegistry() *Registry {
	return &Registry{impls: make(map[reflect.Type]map[string]any), drivers: make(map[string]any)}
}
