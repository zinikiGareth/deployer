package pluggable

type Token interface {
	Loc() Location
}

type Identifier interface {
	Token
	Id() string
}

type Action interface {
	Handle(repo Repository, tokens []Token)
}
