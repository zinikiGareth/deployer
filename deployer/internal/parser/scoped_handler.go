package parser

type ScopedHandlers struct {
}

func NewScopeHandlers() Scoper {
	return &ScopedHandlers{}
}
