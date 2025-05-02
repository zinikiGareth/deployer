package pluggable

type Register interface {
	RegisterNoun(noun string, item Noun)
	RegisterAction(verb string, action Action)
	RegisterFunc(verb string, function Function)
	ProvideDriver(s string, env any)
}

type Recall interface {
	FindAction(verb string) Action
	FindFunc(verb string) Function
	FindNoun(noun string) Noun
	ObtainDriver(driver string) any
}
