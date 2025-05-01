package pluggable

type Register interface {
	RegisterNoun(noun string, item Noun)
	RegisterVerb(verb string, action Action)
	ProvideDriver(s string, env any)
}

type Recall interface {
	FindAction(verb string) Action
	FindFunc(verb string) Function
	FindNoun(noun string) Noun
	ObtainDriver(driver string) any
}
