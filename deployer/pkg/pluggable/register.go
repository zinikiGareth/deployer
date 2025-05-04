package pluggable

import "reflect"

type Register interface {
	Register(what reflect.Type, called string, item any)
	RegisterNoun(noun string, item Noun)
	RegisterFunc(verb string, function Function)
	ProvideDriver(s string, env any)
}

type Recall interface {
	Find(what reflect.Type, called string) any
	FindFunc(verb string) Function
	FindNoun(noun string) Noun
	ObtainDriver(driver string) any
}
