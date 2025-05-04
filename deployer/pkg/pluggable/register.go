package pluggable

import "reflect"

type Register interface {
	Register(what reflect.Type, called string, item any)
	ProvideDriver(s string, env any)
}

type Recall interface {
	Find(what reflect.Type, called string) any
	ObtainDriver(driver string) any
}
