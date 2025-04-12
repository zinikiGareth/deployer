package pluggable

type Register interface {
	RegisterNoun(noun string, item Noun)
	RegisterVerb(verb string, action Action)
}

type Recall interface {
	FindNoun(noun string) Noun
}
