package registry

import "ziniki.org/deployer/driver/pkg/driverbottom"

type Registry struct {
	points      map[string]map[string]any
	drivers     map[string]any
	initDrivers map[string]any
	tools       *driverbottom.CoreTools
}

func (r *Registry) ExtensionPoint(name string) {
	if r.points[name] != nil {
		panic("duplicate extension point " + name)
	}
	r.points[name] = make(map[string]any)
}

func (r *Registry) Register(point string, called string, impl any) {
	// if !reflect.TypeOf(impl).Implements(what) {
	// 	log.Fatalf("Register %s: %v is not a %v", called, impl, what)
	// }
	m := r.points[point]
	if m == nil {
		panic("there is no extension point " + point)
	}
	m[called] = impl
}

func (r *Registry) Find(point string, called string) any {
	m := r.points[point]
	if m == nil {
		panic("there is no extension point " + point)
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
	im, ok := c.(driverbottom.InitMe)
	if ok {
		ret = im.InitMe(r.tools.Storage)
	} else {
		ret = c
	}
	r.initDrivers[forType] = ret
	return ret
}

func (reg *Registry) BindTools(tools *driverbottom.CoreTools) {
	reg.tools = tools
}

func NewRegistry() *Registry {
	return &Registry{points: make(map[string]map[string]any), drivers: make(map[string]any), initDrivers: make(map[string]any)}
}
