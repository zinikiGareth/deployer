package registry

import (
	"log"
	"reflect"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Registry struct {
	impls       map[reflect.Type]map[string]any
	drivers     map[string]any
	initDrivers map[string]any
	tools       *pluggable.Tools
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

func (r *Registry) ObtainDriver(forType string) any {
	ret := r.initDrivers[forType]
	if ret != nil {
		return ret
	}
	c := r.drivers[forType]
	if c == nil {
		panic("there is no driver for " + forType)
	}
	im, ok := c.(pluggable.InitMe)
	if ok {
		ret = im.InitMe(r.tools.Storage)
	} else {
		ret = c
	}
	r.initDrivers[forType] = ret
	return ret
}

func (reg *Registry) BindTools(tools *pluggable.Tools) {
	reg.tools = tools
}

func NewRegistry() *Registry {
	return &Registry{impls: make(map[reflect.Type]map[string]any), drivers: make(map[string]any), initDrivers: make(map[string]any)}
}
