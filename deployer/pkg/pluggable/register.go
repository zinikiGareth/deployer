package pluggable

type Register interface {
	RegisterVerb(verb string, action Action)
}
